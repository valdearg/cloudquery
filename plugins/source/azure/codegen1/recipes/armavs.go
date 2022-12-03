// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"

func Armavs() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armavs.NewGlobalReachConnectionsClient,
		},
    
		{
			NewFunc: armavs.NewCloudLinksClient,
		},
    
		{
			NewFunc: armavs.NewPrivateCloudsClient,
		},
    
		{
			NewFunc: armavs.NewVirtualMachinesClient,
		},
    
		{
			NewFunc: armavs.NewHcxEnterpriseSitesClient,
		},
    
		{
			NewFunc: armavs.NewLocationsClient,
		},
    
		{
			NewFunc: armavs.NewScriptCmdletsClient,
		},
    
		{
			NewFunc: armavs.NewScriptExecutionsClient,
		},
    
		{
			NewFunc: armavs.NewClustersClient,
		},
    
		{
			NewFunc: armavs.NewDatastoresClient,
		},
    
		{
			NewFunc: armavs.NewOperationsClient,
		},
    
		{
			NewFunc: armavs.NewPlacementPoliciesClient,
		},
    
		{
			NewFunc: armavs.NewAddonsClient,
		},
    
		{
			NewFunc: armavs.NewAuthorizationsClient,
		},
    
		{
			NewFunc: armavs.NewWorkloadNetworksClient,
		},
    
		{
			NewFunc: armavs.NewScriptPackagesClient,
		},
    
	}
	return resources
}