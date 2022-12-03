// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"

func Armdataprotection() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armdataprotection.NewClient,
		},
    
		{
			NewFunc: armdataprotection.NewExportJobsClient,
		},
    
		{
			NewFunc: armdataprotection.NewOperationStatusBackupVaultContextClient,
		},
    
		{
			NewFunc: armdataprotection.NewRestorableTimeRangesClient,
		},
    
		{
			NewFunc: armdataprotection.NewRecoveryPointsClient,
		},
    
		{
			NewFunc: armdataprotection.NewExportJobsOperationResultClient,
		},
    
		{
			NewFunc: armdataprotection.NewOperationsClient,
		},
    
		{
			NewFunc: armdataprotection.NewOperationStatusResourceGroupContextClient,
		},
    
		{
			NewFunc: armdataprotection.NewResourceGuardsClient,
		},
    
		{
			NewFunc: armdataprotection.NewBackupPoliciesClient,
		},
    
		{
			NewFunc: armdataprotection.NewBackupVaultOperationResultsClient,
		},
    
		{
			NewFunc: armdataprotection.NewJobsClient,
		},
    
		{
			NewFunc: armdataprotection.NewOperationResultClient,
		},
    
		{
			NewFunc: armdataprotection.NewOperationStatusClient,
		},
    
		{
			NewFunc: armdataprotection.NewBackupInstancesClient,
		},
    
		{
			NewFunc: armdataprotection.NewBackupVaultsClient,
		},
    
	}
	return resources
}