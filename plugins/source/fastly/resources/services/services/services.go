// Code generated by codegen; DO NOT EDIT.

package services

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "fastly_services",
		Description: `https://developer.fastly.com/reference/api/services/service/`,
		Resolver:    fetchServices,
		Columns: []schema.Column{
			{
				Name:     "active_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ActiveVersion"),
			},
			{
				Name:     "comment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Comment"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "customer_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerID"),
			},
			{
				Name:     "deleted_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "versions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Versions"),
			},
		},

		Relations: []*schema.Table{
			ServiceVersions(),
		},
	}
}