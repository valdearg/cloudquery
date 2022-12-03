// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication"

func Armcommunication() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armcommunication.NewServicesClient,
		},
    
		{
			NewFunc: armcommunication.NewEmailServicesClient,
		},
    
		{
			NewFunc: armcommunication.NewDomainsClient,
		},
    
		{
			NewFunc: armcommunication.NewOperationsClient,
		},
    
	}
	return resources
}