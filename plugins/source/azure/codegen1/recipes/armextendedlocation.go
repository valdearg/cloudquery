// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"

func Armextendedlocation() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armextendedlocation.NewCustomLocationsClient,
		},
    
		{
			NewFunc: armextendedlocation.NewResourceSyncRulesClient,
		},
    
	}
	return resources
}