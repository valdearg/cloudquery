// Code generated by codegen; DO NOT EDIT.

package accesscontextmanager

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/accesscontextmanager/apiv1"
)

func UserAccessBindings() *schema.Table {
	return &schema.Table{
		Name:      "gcp_accesscontextmanager_user_access_bindings",
		Resolver:  fetchUserAccessBindings,
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
			},
			{
				Name:     "group_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupKey"),
			},
			{
				Name:     "access_levels",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AccessLevels"),
			},
		},
	}
}

func fetchUserAccessBindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListGcpUserAccessBindingsRequest{}
	gcpClient, err := accesscontextmanager.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListGcpUserAccessBindings(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
