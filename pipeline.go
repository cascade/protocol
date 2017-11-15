package protocol

// Pipeline represents root structure of pipeline.yaml.
type Pipeline struct {

	// Machine represents the provider and the specification of calculation instances.
	Machine *Machine `json:"machine" yaml:"machine"`

	// Steps represents steps of jobs to execute sequentially.
	Steps []*Job `json:"steps" yaml:"steps"`

	// Job field is an alias for "steps" with only 1 element.
	// If steps has any elements, "job" field on the root will be ignored.
	Job *Job `json:"job,omitempty" yaml:"job,omitempty"`
}
