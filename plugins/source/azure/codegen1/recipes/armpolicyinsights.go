// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"

func Armpolicyinsights() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armpolicyinsights.NewAttestationsClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewPolicyStatesClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewPolicyTrackedResourcesClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewRemediationsClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewOperationsClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewPolicyEventsClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewPolicyMetadataClient,
		},
    
		{
			NewFunc: armpolicyinsights.NewPolicyRestrictionsClient,
		},
    
	}
	return resources
}