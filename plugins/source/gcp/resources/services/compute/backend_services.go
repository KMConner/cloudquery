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

func BackendServices() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_backend_services",
		Resolver:  fetchBackendServices,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "affinity_cookie_ttl_sec",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AffinityCookieTtlSec"),
			},
			{
				Name:     "backends",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Backends"),
			},
			{
				Name:     "cdn_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CdnPolicy"),
			},
			{
				Name:     "circuit_breakers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CircuitBreakers"),
			},
			{
				Name:     "connection_draining",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConnectionDraining"),
			},
			{
				Name:     "connection_tracking_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConnectionTrackingPolicy"),
			},
			{
				Name:     "consistent_hash",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConsistentHash"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "custom_request_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CustomRequestHeaders"),
			},
			{
				Name:     "custom_response_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CustomResponseHeaders"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "edge_security_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EdgeSecurityPolicy"),
			},
			{
				Name:     "enable_cdn",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableCDN"),
			},
			{
				Name:     "failover_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailoverPolicy"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "health_checks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("HealthChecks"),
			},
			{
				Name:     "iap",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Iap"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "load_balancing_scheme",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoadBalancingScheme"),
			},
			{
				Name:     "locality_lb_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocalityLbPolicies"),
			},
			{
				Name:     "locality_lb_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LocalityLbPolicy"),
			},
			{
				Name:     "log_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogConfig"),
			},
			{
				Name:     "max_stream_duration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaxStreamDuration"),
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
				Name:     "outlier_detection",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutlierDetection"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "port_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PortName"),
			},
			{
				Name:     "protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Protocol"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "security_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityPolicy"),
			},
			{
				Name:     "security_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecuritySettings"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_bindings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ServiceBindings"),
			},
			{
				Name:     "session_affinity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SessionAffinity"),
			},
			{
				Name:     "subsetting",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subsetting"),
			},
			{
				Name:     "timeout_sec",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TimeoutSec"),
			},
		},
	}
}

func fetchBackendServices(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListBackendServicesRequest{
		Project: c.ProjectId,
	}
	it := c.Services.ComputeBackendServicesClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp.Value.BackendServices

	}
	return nil
}
