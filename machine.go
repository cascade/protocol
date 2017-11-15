package protocol

// MachineProvider ...
type MachineProvider string

const (
	// MachineEC2 ...
	MachineEC2 MachineProvider = "ec2"
	// MachineGCE ...
	MachineGCE MachineProvider = "gce"
)

// Machine ...
type Machine struct {
	// Provider represents the platform which provides computing instances.
	Provider MachineProvider `json:"provider" yaml:"provider"`
	// "ec2" field must be given when "provider" is "ec2".
	EC2 *AmazonEC2 `json:"ec2,omitempty" yaml:"ec2,omitempty"`
	// "gce" field must be given when "provider" is "gce".
	GCE *GoogleComputeEngine `json:"gce,omitempty" yaml:"gce,omitempty"`
}

// AmazonEC2 ...
type AmazonEC2 struct {
	// Type represents machine type of EC2, e.g. t2.micro, m4.xlarge, ...
	Type string `json:"type" yaml:"type"`
}

// GoogleComputeEngine ...
type GoogleComputeEngine struct {
	// Type represents machine type of GCE, e.g. n1-standard-1, n1-highmem-2, ...
	Type string `json:"type" yaml:"type"`
}
