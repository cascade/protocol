package protocol

// Execute defines the real execution of this "Job".
type Execute struct {

	// Inline shell command to execute on the specified runtime.
	Inline *string `json:"inline,omitempty" yaml:"inline,omitempty"`

	// Local script to execute on the specified runtime.
	Script *string `json:"script,omitempty" yaml:"script,omitempty"`

	// Container Image to execute as a process,
	// which must have ENTRYPOINT instead of using "script" field.
	Container *string `json:"container,omitempty" yaml:"container,omitempty"`
}
