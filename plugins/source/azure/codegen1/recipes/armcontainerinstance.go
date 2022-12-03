// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"

func Armcontainerinstance() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armcontainerinstance.NewContainerGroupsClient,
		},
    
		{
			NewFunc: armcontainerinstance.NewContainersClient,
		},
    
		{
			NewFunc: armcontainerinstance.NewLocationClient,
		},
    
		{
			NewFunc: armcontainerinstance.NewOperationsClient,
		},
    
		{
			NewFunc: armcontainerinstance.NewSubnetServiceAssociationLinkClient,
		},
    
	}
	return resources
}