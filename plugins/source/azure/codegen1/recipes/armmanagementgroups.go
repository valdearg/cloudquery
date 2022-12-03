// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"

func Armmanagementgroups() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armmanagementgroups.NewEntitiesClient,
		},
    
		{
			NewFunc: armmanagementgroups.NewHierarchySettingsClient,
		},
    
		{
			NewFunc: armmanagementgroups.NewManagementGroupSubscriptionsClient,
		},
    
		{
			NewFunc: armmanagementgroups.NewClient,
		},
    
		{
			NewFunc: armmanagementgroups.NewOperationsClient,
		},
    
		{
			NewFunc: armmanagementgroups.NewAPIClient,
		},
    
	}
	return resources
}