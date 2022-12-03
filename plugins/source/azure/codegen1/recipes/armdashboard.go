// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"

func Armdashboard() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armdashboard.NewPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armdashboard.NewOperationsClient,
		},
    
		{
			NewFunc: armdashboard.NewPrivateLinkResourcesClient,
		},
    
		{
			NewFunc: armdashboard.NewGrafanaClient,
		},
    
	}
	return resources
}