package client

import (
	"context"
	"database/sql"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"golang.org/x/exp/slices"
)

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}

func (c *Client) getPrimaryKey(table *schema.Table) []string {
	if !c.enabledPks() {
		return nil
	}
	return table.PrimaryKeys()
}

func (c *Client) getColumnDefinitions(columns schema.ColumnList) []string {
	definitions := make([]string, len(columns))
	for i, column := range columns {
		sqlType := msSQLType(column.Type)
		columnName := sanitizeIdentifier(column.Name)
		fieldDef := columnName + " " + sqlType
		if column.Name == "_cq_id" {
			// _cq_id column should always have a "unique not null" constraint
			fieldDef += " UNIQUE NOT NULL"
		}
		definitions[i] = fieldDef
	}
	return definitions
}

type (
	msSQLTableInfo struct {
		name    string
		columns []*msSQLColumn
	}
	msSQLColumn struct {
		name string
		typ  string
	}
)

func (info *msSQLTableInfo) get(column string) *msSQLColumn {
	idx := slices.IndexFunc(info.columns, func(c *msSQLColumn) bool {
		return c.name == column
	})
	if idx < 0 {
		return nil
	}
	return info.columns[idx]
}

func (c *Client) getTableColumns(ctx context.Context, tableName string) (*msSQLTableInfo, error) {
	tc := msSQLTableInfo{
		name: tableName,
	}

	// https://stackoverflow.com/a/58995161
	const sqlSelectColumnTypes = `
select
    [name] = c.[name], [type] = case
    when tp.[name] in ('varchar', 'char', 'varbinary')
    then tp.[name] + '(' + IIF(c.max_length = -1, 'max', cast (c.max_length as varchar (25))) + ')'
    when tp.[name] in ('nvarchar', 'nchar')
    then tp.[name] + '(' + IIF(c.max_length = -1, 'max', cast (c.max_length / 2 as varchar (25)))+ ')'
    when tp.[name] in ('decimal', 'numeric')
    then tp.[name] + '(' + cast (c.[precision] as varchar (25)) + ', ' + cast (c.[scale] as varchar (25)) + ')'
    when tp.[name] in ('datetime2')
    then tp.[name] + '(' + cast (c.[scale] as varchar (25)) + ')'
    else tp.[name]
end
from sys.tables t 
join sys.schemas s ON t.schema_id = s.schema_id
join sys.columns c ON t.object_id = c.object_id
join sys.types tp ON c.user_type_id = tp.user_type_id
where 
    s.[name] = 'dbo' -- use default schema here
    and t.[name] = $tableName`
	rows, err := c.db.QueryContext(ctx, sqlSelectColumnTypes, sql.Named("tableName", tableName))
	if err != nil {
		return nil, err
	}

	if err := processRows(rows, func(row *sql.Rows) error {
		var name string
		var typ string

		if err := row.Scan(&name, &typ); err != nil {
			return err
		}

		tc.columns = append(tc.columns, &msSQLColumn{
			name: strings.ToLower(name),
			typ:  strings.ToLower(typ),
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return &tc, nil
}

func (c *Client) getTablePKColumns(ctx context.Context, tableName string) ([]string, error) {
	const sqlSelectPKColumns = `select
    [name] = tc.name,
    [order] = ic.key_ordinal
from sys.schemas s
    inner join sys.tables t on s.schema_id = t.schema_id
    inner join sys.indexes i on t.object_id = i.object_id
    inner join sys.index_columns ic on i.object_id = ic.object_id and i.index_id = ic.index_id
    inner join sys.columns tc on ic.object_id = tc.object_id and ic.column_id = tc.column_id
where s.[name] = 'dbo' -- use default schema here
  and t.[name] = $tableName
  and i.is_primary_key = 1
order by ic.key_ordinal`
	rows, err := c.db.QueryContext(ctx, sqlSelectPKColumns, sql.Named("tableName", tableName))
	if err != nil {
		return nil, err
	}

	var result []string
	if err := processRows(rows, func(row *sql.Rows) error {
		var name string
		var idx int

		if err := rows.Scan(&name, &idx); err != nil {
			return err
		}

		result = append(result, name)

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) getStalePks(table *schema.Table, primaryKey []string) []string {
	if !c.enabledPks() {
		return nil
	}

	schemaPK := table.PrimaryKeys()
	var stale []string
	for _, key := range primaryKey {
		if !slices.Contains(schemaPK, key) {
			stale = append(stale, key)
		}
	}

	return stale
}

func getDropNotNullQuery(table *schema.Table, stalePK []string) string {
	queries := make([]string, len(stalePK))
	for i, col := range stalePK {
		queries[i] = "alter table " + sanitizeIdentifier(table.Name) + " alter column " + sanitizeIdentifier(col) + " drop not null;"
	}

	return strings.Join(queries, "\n")
}
