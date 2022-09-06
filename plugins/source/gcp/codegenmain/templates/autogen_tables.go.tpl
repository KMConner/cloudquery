// Code generated by codegen; DO NOT EDIT directly.

package plugin

import (
  "github.com/cloudquery/plugin-sdk/schema"
	{{- range .}}
	"github.com/cloudquery/plugins/source/gcp/resources/services/{{.Service}}"
	{{- end}}
)


func PluginAutoGeneratedTables() []*schema.Table {
  return []*schema.Table{
			{{- range .}}
      {{- if not .ChildTable}}
			{{.Service}}.{{.SubService | ToCamel}}(),
      {{- end}}
			{{- end}}
  }
}
