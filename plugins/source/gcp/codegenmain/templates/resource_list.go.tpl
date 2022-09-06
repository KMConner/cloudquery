// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"github.com/pkg/errors"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.{{.Service}}{{.ClientName}}.{{.ListFunctionName}}()
		if err != nil {
			return errors.WithStack(err)
		}
    res <- output.{{.OutputField}}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
