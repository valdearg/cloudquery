// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"

func Armeventgrid() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armeventgrid.NewVerifiedPartnersClient,
		},
    
		{
			NewFunc: armeventgrid.NewOperationsClient,
		},
    
		{
			NewFunc: armeventgrid.NewExtensionTopicsClient,
		},
    
		{
			NewFunc: armeventgrid.NewSystemTopicEventSubscriptionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewTopicEventSubscriptionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewEventSubscriptionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewPartnerNamespacesClient,
		},
    
		{
			NewFunc: armeventgrid.NewPartnerRegistrationsClient,
		},
    
		{
			NewFunc: armeventgrid.NewDomainEventSubscriptionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewPartnerConfigurationsClient,
		},
    
		{
			NewFunc: armeventgrid.NewPrivateLinkResourcesClient,
		},
    
		{
			NewFunc: armeventgrid.NewSystemTopicsClient,
		},
    
		{
			NewFunc: armeventgrid.NewChannelsClient,
		},
    
		{
			NewFunc: armeventgrid.NewDomainTopicEventSubscriptionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewTopicTypesClient,
		},
    
		{
			NewFunc: armeventgrid.NewPartnerTopicEventSubscriptionsClient,
		},
    
		{
			NewFunc: armeventgrid.NewPartnerTopicsClient,
		},
    
		{
			NewFunc: armeventgrid.NewTopicsClient,
		},
    
		{
			NewFunc: armeventgrid.NewDomainsClient,
		},
    
		{
			NewFunc: armeventgrid.NewDomainTopicsClient,
		},
    
	}
	return resources
}