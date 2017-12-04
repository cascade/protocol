package protocol

// Job defines the contents of the "Component"
type Job struct {

	// Steps specifies nested components which should be handled in series.
	Steps []*Component `json:"steps,omitempty" yaml:"steps,omitempty"`
	// Pararrell specifies nested components which should be handled in parallel.
	Parallel []*Component `json:"parallel,omitempty" yaml:"parallel,omitempty"`
	// Properties hereafter would be ignored if eigher of "Steps" or "Parallel" is given.

	// Runtime defines a machine runtime such as OS, dependencies and language versions.
	// If you want to use just "local" for the runtime, omit this field.
	Runtime *Runtime `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	// Execute defines the real execution process of this "Job".
	// This field is required, can't be omitted.
	Execute *Execute `json:"execute" yaml:"execute"`
}

// Runtime defines a machine runtime such as OS, dependencies and language versions.
// If you want to use "local" for the runtime, just omit "runtime" field.
type Runtime struct {
	// Image is a OS image on which "Execute" is executed, unless "Execute" doesn't have "Container"
	// If you want to use "local" for the runtime, just omit "runtime" field or keep this field empty.
	Image string `json:"image"   yaml:"image"`
}

// Execute defines the real execution of this "Job".
type Execute struct {

	// Inline shell command to execute on the specified runtime.
	Inline string `json:"inline,omitempty" yaml:"inline,omitempty"`

	// Local script to execute on the specified runtime.
	Script string `json:"script,omitempty" yaml:"script,omitempty"`

	// Container Image to execute as a process,
	// which must have ENTRYPOINT instead of using "script" field.
	Container string `json:"container,omitempty" yaml:"container,omitempty"`
}
