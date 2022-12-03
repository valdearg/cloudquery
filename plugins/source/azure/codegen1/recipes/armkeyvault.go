// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"

func Armkeyvault() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armkeyvault.NewMHSMPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armkeyvault.NewOperationsClient,
		},
    
		{
			NewFunc: armkeyvault.NewPrivateLinkResourcesClient,
		},
    
		{
			NewFunc: armkeyvault.NewKeysClient,
		},
    
		{
			NewFunc: armkeyvault.NewManagedHsmsClient,
		},
    
		{
			NewFunc: armkeyvault.NewVaultsClient,
		},
    
		{
			NewFunc: armkeyvault.NewPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armkeyvault.NewSecretsClient,
		},
    
		{
			NewFunc: armkeyvault.NewMHSMPrivateLinkResourcesClient,
		},
    
	}
	return resources
}