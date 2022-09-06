// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Addresses() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_addresses",
		Resolver:  fetchAddresses,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Address"),
			},
			{
				Name:     "address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AddressType"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "ip_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpVersion"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "network_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkTier"),
			},
			{
				Name:     "prefix_length",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PrefixLength"),
			},
			{
				Name:     "purpose",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Purpose"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subnetwork",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Subnetwork"),
			},
			{
				Name:     "users",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Users"),
			},
		},
	}
}

func fetchAddresses(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListAddressesRequest{
		Project: c.ProjectId,
	}
	it := c.Services.ComputeAddressesClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.Addresses

	}
	return nil
}
