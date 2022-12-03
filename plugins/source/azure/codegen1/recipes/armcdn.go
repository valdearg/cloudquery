// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"

func Armcdn() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armcdn.NewAFDEndpointsClient,
		},
    
		{
			NewFunc: armcdn.NewAFDOriginsClient,
		},
    
		{
			NewFunc: armcdn.NewOperationsClient,
		},
    
		{
			NewFunc: armcdn.NewValidateClient,
		},
    
		{
			NewFunc: armcdn.NewManagementClient,
		},
    
		{
			NewFunc: armcdn.NewRulesClient,
		},
    
		{
			NewFunc: armcdn.NewEdgeNodesClient,
		},
    
		{
			NewFunc: armcdn.NewRoutesClient,
		},
    
		{
			NewFunc: armcdn.NewRuleSetsClient,
		},
    
		{
			NewFunc: armcdn.NewSecretsClient,
		},
    
		{
			NewFunc: armcdn.NewCustomDomainsClient,
		},
    
		{
			NewFunc: armcdn.NewEndpointsClient,
		},
    
		{
			NewFunc: armcdn.NewManagedRuleSetsClient,
		},
    
		{
			NewFunc: armcdn.NewProfilesClient,
		},
    
		{
			NewFunc: armcdn.NewAFDOriginGroupsClient,
		},
    
		{
			NewFunc: armcdn.NewAFDProfilesClient,
		},
    
		{
			NewFunc: armcdn.NewPoliciesClient,
		},
    
		{
			NewFunc: armcdn.NewSecurityPoliciesClient,
		},
    
		{
			NewFunc: armcdn.NewOriginsClient,
		},
    
		{
			NewFunc: armcdn.NewResourceUsageClient,
		},
    
		{
			NewFunc: armcdn.NewAFDCustomDomainsClient,
		},
    
		{
			NewFunc: armcdn.NewLogAnalyticsClient,
		},
    
		{
			NewFunc: armcdn.NewOriginGroupsClient,
		},
    
	}
	return resources
}