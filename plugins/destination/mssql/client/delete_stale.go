package client

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	for _, table := range tables.FlattenTables() {
		_, err = tx.ExecContext(ctx,
			fmt.Sprintf(`delete from %s where %s = $sourceName and %s < $syncTime`,
				sanitizeIdentifier(table.Name),
				sanitizeIdentifier(schema.CqSourceNameColumn.Name),
				sanitizeIdentifier(schema.CqSyncTimeColumn.Name),
			),
			sql.Named("sourceName", sourceName),
			sql.Named("syncTime", syncTime),
		)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
