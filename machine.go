package protocol

// Machine defines on which machine "Job" is gonna be handled.
type Machine struct {
	// Provider specifies machine provider, either of [local, ec2, gce, k8s]
	Provider string `json:"provider" yaml:"provider"`
	// CPU specifies how many CPU cores are required for the "Job"
	CPU int `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	// Memory specifies how much memory are required for the "Job"
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
