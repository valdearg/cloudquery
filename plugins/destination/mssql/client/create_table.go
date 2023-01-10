package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) createTable(ctx context.Context, table *schema.Table) error {
	c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")

	var sb strings.Builder
	sb.WriteString("create table ")
	sb.WriteString(sanitizeIdentifier(table.Name))

	sb.WriteString(" (")
	sb.WriteString(strings.Join(c.getColumnDefinitions(table.Columns), ", "))

	sb.WriteString(", constraint ")
	sb.WriteString(sanitizeIdentifier(table.Name + "_cqpk"))

	sb.WriteString(" primary key (")
	primaryKey := c.getPrimaryKey(table)
	if len(primaryKey) == 0 {
		// if no primary keys are defined, add a PK constraint for _cq_id
		primaryKey = []string{"_cq_id"}
	}
	sb.WriteString(strings.Join(primaryKey, ","))
	sb.WriteString("))")

	_, err := c.db.ExecContext(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", table.Name, err)
	}

	return nil
}
