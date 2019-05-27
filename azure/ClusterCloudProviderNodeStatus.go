package azure

type ClusterCloudProviderNodeStatus string

const (
	ClusterCloudProviderNodeStatusNotEnabledOnSubscription = "NotEnabledOnSubscription"
	ClusterCloudProviderNodeStatusNotAvailableInRegion     = "NotAvailableInRegion"
)