// Code generated by codegen; DO NOT EDIT.

package secretmanager

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:      "gcp_secretmanager_secrets",
		Resolver:  fetchSecrets,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "replication",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Replication"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "topics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Topics"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "rotation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rotation"),
			},
			{
				Name:     "version_aliases",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VersionAliases"),
			},
		},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListSecretsRequest{
		Parent: "projects/" + c.ProjectId,
	}
	it := c.Services.SecretmanagerClient.ListSecrets(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
