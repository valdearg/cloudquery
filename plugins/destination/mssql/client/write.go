package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) Write(ctx context.Context, tables schema.Tables, resources <-chan *destination.ClientResource) error {
	query := upsert
	if c.spec.WriteMode == specs.WriteModeAppend {
		query = insert
	}

	batchSize := c.spec.BatchSize

	batch := make([]*stmt, 0, batchSize)
	for r := range resources {
		table := tables.Get(r.TableName)
		if table == nil {
			panic(fmt.Errorf("table %s not found", r.TableName))
		}

		batch = append(batch, &stmt{query: query(table), params: r.Data})
		if len(batch) == batchSize {
			if err := c.send(ctx, batch); err != nil {
				return err
			}

			// just clip, capacity is left intact (opposed to batch[:0:0])
			batch = batch[:0]
		}
	}

	return c.send(ctx, batch)
}

func insert(table *schema.Table) string {
	var sb strings.Builder

	sb.WriteString("insert into ")
	sb.WriteString(sanitizeIdentifier(table.Name))
	sb.WriteString(" (")

	columns := tableColumnsSanitized(table)
	sb.WriteString(strings.Join(columns, ", "))
	sb.WriteString(") values (")

	params := make([]string, len(columns))
	for i := range columns {
		params[i] = strconv.Itoa(i)
	}
	sb.WriteString(strings.Join(params, ", "))

	sb.WriteString(")")

	return sb.String()
}
func upsert(table *schema.Table) string {
	var sb strings.Builder

	sb.WriteString(insert(table))

	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(sanitizeIdentifier(fmt.Sprintf("%s_cqpk", table.Name)))
	sb.WriteString(" do update set ")

	excluded := make([]string, 0, len(table.Columns))
	for _, c := range table.Columns {
		if c.Name == schema.CqIDColumn.Name || c.Name == schema.CqParentIDColumn.Name {
			continue
		}

		name := sanitizeIdentifier(c.Name)
		excluded = append(excluded, name+"=excluded."+name) // excluded references the new values
	}

	sb.WriteString(strings.Join(excluded, ", "))

	return sb.String()
}
