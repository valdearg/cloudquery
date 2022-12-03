// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"

func Armmediaservices() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armmediaservices.NewTransformsClient,
		},
    
		{
			NewFunc: armmediaservices.NewOperationResultsClient,
		},
    
		{
			NewFunc: armmediaservices.NewTracksClient,
		},
    
		{
			NewFunc: armmediaservices.NewAssetTrackOperationResultsClient,
		},
    
		{
			NewFunc: armmediaservices.NewOperationsClient,
		},
    
		{
			NewFunc: armmediaservices.NewStreamingLocatorsClient,
		},
    
		{
			NewFunc: armmediaservices.NewAccountFiltersClient,
		},
    
		{
			NewFunc: armmediaservices.NewClient,
		},
    
		{
			NewFunc: armmediaservices.NewLocationsClient,
		},
    
		{
			NewFunc: armmediaservices.NewAssetFiltersClient,
		},
    
		{
			NewFunc: armmediaservices.NewAssetTrackOperationStatusesClient,
		},
    
		{
			NewFunc: armmediaservices.NewContentKeyPoliciesClient,
		},
    
		{
			NewFunc: armmediaservices.NewLiveOutputsClient,
		},
    
		{
			NewFunc: armmediaservices.NewJobsClient,
		},
    
		{
			NewFunc: armmediaservices.NewLiveEventsClient,
		},
    
		{
			NewFunc: armmediaservices.NewStreamingEndpointsClient,
		},
    
		{
			NewFunc: armmediaservices.NewOperationStatusesClient,
		},
    
		{
			NewFunc: armmediaservices.NewPrivateLinkResourcesClient,
		},
    
		{
			NewFunc: armmediaservices.NewAssetsClient,
		},
    
		{
			NewFunc: armmediaservices.NewPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armmediaservices.NewStreamingPoliciesClient,
		},
    
	}
	return resources
}