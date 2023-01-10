package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/slices"
)

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	c.logger.Info().Str("table", table.Name).Msg("Table exists, auto-migrating")

	pkPresent, err := c.getTablePKColumns(ctx, table.Name)
	if err != nil {
		return err
	}

	stalePks := c.getStalePks(table, pkPresent)
	if len(stalePks) > 0 {
		dropConstraintSQL := "alter table " + sanitizeIdentifier(table.Name) +
			" drop " + sanitizeIdentifier(table.Name+"_cqpk")
		sep := strings.Repeat("-", len(dropConstraintSQL)+1)
		query := fmt.Sprintf("%s\n%s;\n%s\n%s", sep, dropConstraintSQL, getDropNotNullQuery(table, stalePks), sep)
		return fmt.Errorf(
			`the following primary keys were removed from the schema %q for table %q.
You can migrate the table manually by running:
%s`, stalePks, table.Name, query)
	}

	return c.ensureColumns(ctx, table, pkPresent)
}

func (c *Client) ensureColumns(ctx context.Context, table *schema.Table, pkPresent []string) error {
	current, err := c.getTableColumns(ctx, table.Name)
	if err != nil {
		return err
	}

	tableName := sanitizeIdentifier(table.Name)
	recreatePK := false

	for _, column := range table.Columns {
		columnName := sanitizeIdentifier(column.Name)
		columnType := msSQLType(column.Type)

		switch currColumn := current.get(column.Name); {
		case currColumn == nil:
			// column doesn't exist
			c.logger.Info().
				Str("table", table.Name).
				Str("column", column.Name).
				Msg("Column doesn't exist, creating")

			recreatePK = recreatePK || column.CreationOptions.PrimaryKey

			if _, err := c.db.ExecContext(ctx, "alter table "+tableName+" add "+columnName+" "+columnType); err != nil {
				return fmt.Errorf("failed to add column %s to table %s: %w", column.Name, table.Name, err)
			}
		case currColumn.typ != columnType:
			// column exists but type is different
			c.logger.Info().
				Str("table", table.Name).
				Str("column", column.Name).
				Str("old_type", currColumn.typ).
				Str("new_type", columnType).
				Msg("Column exists but type is different, re-creating")

			// if the new PK contains this column we will need to recreate the primary key
			recreatePK = recreatePK || column.CreationOptions.PrimaryKey

			// right now we will drop the column and re-create. in the future we will have an option to automigrate
			if _, err := c.db.ExecContext(ctx, "alter table "+tableName+" drop column "+columnName); err != nil {
				return fmt.Errorf("failed to drop column %s on table %s: %w", column.Name, table.Name, err)
			}

			if _, err := c.db.ExecContext(ctx, "alter table "+tableName+" add "+columnName+" "+columnType); err != nil {
				return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, table.Name, err)
			}
		}

		// column exists and type is the same but constraints might differ
		// check that PK contains both of new & old columns or neither of them
		if column.CreationOptions.PrimaryKey != slices.Contains(pkPresent, column.Name) {
			recreatePK = true

			if c.enabledPks() {
				c.logger.Info().
					Str("table", table.Name).
					Str("column", column.Name).
					Bool("pk", column.CreationOptions.PrimaryKey).
					Msg("Column exists with different primary keys")
			}
		}
	}

	if c.enabledPks() && recreatePK {
		return c.recreatePK(ctx, table)
	}

	return nil
}

func (c *Client) recreatePK(ctx context.Context, table *schema.Table) (err error) {
	c.logger.Info().Str("table", table.Name).Msg("Recreating primary keys")
	tableName := sanitizeIdentifier(table.Name)
	constraintName := sanitizeIdentifier(table.Name + "_cqpk")

	if err := c.setNotNullOnPks(ctx, table); err != nil {
		return fmt.Errorf("failed to enforce not null on primary keys: %w", err)
	}

	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction to recreate primary keys: %w", err)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				c.logger.Error().Err(err).Msg("failed to rollback transaction")
			}
		}
	}()

	if _, err := tx.ExecContext(ctx, "alter table "+tableName+" drop "+constraintName); err != nil {
		return fmt.Errorf("failed to drop primary key constraint on table %s: %w", table.Name, err)
	}

	if _, err := tx.ExecContext(ctx,
		"alter table "+tableName+
			" add constraint "+constraintName+
			" primary key ("+strings.Join(table.PrimaryKeys(), ",")+")"); err != nil {
		return fmt.Errorf("failed to add primary key constraint on table %s: %w", table.Name, err)
	}

	return tx.Commit()
}

func (c *Client) setNotNullOnPks(ctx context.Context, table *schema.Table) error {
	tableName := sanitizeIdentifier(table.Name)
	for _, col := range table.PrimaryKeys() {
		query := "alter table " + tableName + " alter column " + sanitizeIdentifier(col) + " not null"
		if _, err := c.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to set not null on column %s on table %s: %w", col, table.Name, err)
		}
	}
	return nil
}
