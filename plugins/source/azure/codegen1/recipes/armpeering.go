// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"

func Armpeering() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armpeering.NewServiceCountriesClient,
		},
    
		{
			NewFunc: armpeering.NewPrefixesClient,
		},
    
		{
			NewFunc: armpeering.NewOperationsClient,
		},
    
		{
			NewFunc: armpeering.NewServiceLocationsClient,
		},
    
		{
			NewFunc: armpeering.NewServiceProvidersClient,
		},
    
		{
			NewFunc: armpeering.NewServicesClient,
		},
    
		{
			NewFunc: armpeering.NewLookingGlassClient,
		},
    
		{
			NewFunc: armpeering.NewManagementClient,
		},
    
		{
			NewFunc: armpeering.NewPeerAsnsClient,
		},
    
		{
			NewFunc: armpeering.NewPeeringsClient,
		},
    
		{
			NewFunc: armpeering.NewCdnPeeringPrefixesClient,
		},
    
		{
			NewFunc: armpeering.NewConnectionMonitorTestsClient,
		},
    
		{
			NewFunc: armpeering.NewLegacyPeeringsClient,
		},
    
		{
			NewFunc: armpeering.NewLocationsClient,
		},
    
		{
			NewFunc: armpeering.NewReceivedRoutesClient,
		},
    
		{
			NewFunc: armpeering.NewRegisteredAsnsClient,
		},
    
		{
			NewFunc: armpeering.NewRegisteredPrefixesClient,
		},
    
	}
	return resources
}