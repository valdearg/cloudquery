// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction"

func Armnetworkfunction() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsByResourceGroupClient,
		},
    
		{
			NewFunc: armnetworkfunction.NewClient,
		},
    
		{
			NewFunc: armnetworkfunction.NewCollectorPoliciesClient,
		},
    
		{
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsClient,
		},
    
		{
			NewFunc: armnetworkfunction.NewAzureTrafficCollectorsBySubscriptionClient,
		},
    
	}
	return resources
}