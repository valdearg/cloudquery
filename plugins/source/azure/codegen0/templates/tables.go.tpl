// Code generated by codegen; DO NOT EDIT.
package plugin

import (
  "github.com/cloudquery/plugin-sdk/schema"  
  {{- range .}}
  "github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/{{.PackageName}}"
  {{- end}}
)

func generatedTables() []*schema.Table {
  return []*schema.Table{
    {{- range .}}
    {{.PackageName}}.{{.Name | ToCamel}}(),
    {{- end}}
  }
}