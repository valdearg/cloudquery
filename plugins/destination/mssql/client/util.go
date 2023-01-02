package client

import (
	"database/sql"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func sanitizeIdentifier(parts ...string) string {
	res := make([]string, len(parts))

	for i, part := range parts {
		res[i] = `[` + strings.ReplaceAll(part, string([]byte{0}), "") + `]`
	}

	return strings.Join(parts, ".")
}

func processRows(rows *sql.Rows, process func(row *sql.Rows) error) error {
	defer rows.Close()

	for next := true; next; next = rows.NextResultSet() {
		for rows.Next() {
			if err := process(rows); err != nil {
				return err
			}
		}
	}

	return rows.Err()
}

func tableColumnsSanitized(table *schema.Table) []string {
	result := make([]string, len(table.Columns))
	for i, column := range table.Columns {
		result[i] = sanitizeIdentifier(column.Name)
	}
	return result
}
