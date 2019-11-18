package v20191231preview

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster struct {
	// The resource ID (immutable).
	ID string `json:"id,omitempty"`

	// The resource name (immutable).
	Name string `json:"name,omitempty"`

	// The resource type (immutable).
	Type string `json:"type,omitempty"`

	// The resource location (immutable).
	Location string `json:"location,omitempty"`

	// The resource tags.
	Tags Tags `json:"tags,omitempty"`

	// The cluster properties.
	Properties Properties `json:"properties,omitempty"`
}

// Tags represents an OpenShift cluster's tags.
type Tags map[string]string

// Properties represents an OpenShift cluster's properties.
type Properties struct {
	// The cluster provisioning state (immutable).
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`

	// The cluster network profile.
	NetworkProfile NetworkProfile `json:"networkProfile,omitempty"`

	// The cluster master profile.
	MasterProfile MasterProfile `json:"masterProfile,omitempty"`

	// The cluster worker profiles.
	WorkerProfiles []WorkerProfile `json:"workerProfiles,omitempty"`

	// The URL to access the cluster API server (immutable).
	APIServerURL string `json:"apiserverURL,omitempty"`

	// The URL to access the cluster console (immutable).
	ConsoleURL string `json:"consoleURL,omitempty"`
}

// ProvisioningState represents a provisioning state.
type ProvisioningState string

// ProvisioningState constants
const (
	ProvisioningStateUpdating  ProvisioningState = "Updating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateFailed    ProvisioningState = "Failed"
)

// NetworkProfile represents a network profile.
type NetworkProfile struct {
	// The CIDR used for the cluster VMs and IPs (immutable).
	VNetCIDR string `json:"vnetCidr,omitempty"`

	// The CIDR used for OpenShift/Kubernetes Pods (immutable).
	PodCIDR string `json:"podCidr,omitempty"`

	// The CIDR used for OpenShift/Kubernetes Services (immutable).
	ServiceCIDR string `json:"serviceCidr,omitempty"`
}

// MasterProfile represents a master profile.
type MasterProfile struct {
	// The size of the master VMs.
	VMSize VMSize `json:"vmSize,omitempty"`
}

// VMSize represents a VM size.
type VMSize string

// VMSize constants
const (
	VMSizeStandardD2sV3 VMSize = "Standard_D2s_v3"
	VMSizeStandardD4sV3 VMSize = "Standard_D4s_v3"
	VMSizeStandardD8sV3 VMSize = "Standard_D8s_v3"
)

// WorkerProfile represents a worker profile.
type WorkerProfile struct {
	// The worker profile name.  Must be "worker".
	Name string `json:"name,omitempty"`

	// The size of the worker VMs.
	VMSize VMSize `json:"vmSize,omitempty"`

	// The disk size of the worker VMs.  Must be 128 or greater.
	DiskSizeGB int `json:"diskSizeGB,omitempty"`

	// The number of worker VMs.  Must be between 3 and 20.
	Count int `json:"count,omitempty"`
}
