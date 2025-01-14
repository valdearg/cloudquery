// Code generated by codegen; DO NOT EDIT.
package frontdoor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func NetworkExperimentProfiles() *schema.Table {
	return &schema.Table{
		Name:      "azure_frontdoor_network_experiment_profiles",
		Resolver:  fetchNetworkExperimentProfiles,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_network),
		Transform: transformers.TransformWithStruct(&armfrontdoor.Profile{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchNetworkExperimentProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armfrontdoor.NewNetworkExperimentProfilesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
