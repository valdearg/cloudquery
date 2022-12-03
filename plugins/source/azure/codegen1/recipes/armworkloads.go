// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func Armworkloads() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armworkloads.NewMonitorsClient,
		},
    
		{
			NewFunc: armworkloads.NewSAPCentralInstancesClient,
		},
    
		{
			NewFunc: armworkloads.NewPhpWorkloadsClient,
		},
    
		{
			NewFunc: armworkloads.NewProviderInstancesClient,
		},
    
		{
			NewFunc: armworkloads.NewSAPApplicationServerInstancesClient,
		},
    
		{
			NewFunc: armworkloads.NewSAPDatabaseInstancesClient,
		},
    
		{
			NewFunc: armworkloads.NewClient,
		},
    
		{
			NewFunc: armworkloads.NewOperationsClient,
		},
    
		{
			NewFunc: armworkloads.NewSAPVirtualInstancesClient,
		},
    
		{
			NewFunc: armworkloads.NewSKUsClient,
		},
    
		{
			NewFunc: armworkloads.NewWordpressInstancesClient,
		},
    
	}
	return resources
}