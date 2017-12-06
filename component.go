package protocol

// Component is an unit of pipeline, with defining interface of itself.
type Component struct {
	// Machine defines machine physical specs on which this component
	// or nested components, if not specified for them, would be executed.
	Machine *Machine `json:"machine,omitempty" yaml:"machine"`

	// Steps can define child-components which will be handled in series.
	// If "Steps" are specified, this component is regarded as "Components Wrapper",
	// therefore any of other properties such as "Execute" or "Runtime" will be ignored.
	Steps []*Component `json:"steps,omitempty" yaml:"steps"`

	// Parallel can define child-components which will be handled in parallel.
	// If "Parallel" are specified, this component is regarded as "Components Wrapper",
	// therefore any of other properties such as "Execute" or "Runtime" will be ignored.
	Parallel []*Component `json:"parallel,omitempty" yaml:"parallel"`

	// Runtime defines runtime such as OS, languages, and dependencies
	// on which this component would be executed.
	Runtime *Runtime `json:"runtime,omitempty" yaml:"runtime"`

	// Execute defines what to be executed on this component.
	Execute *Execute `json:"execute,omitmepty" yaml:"execute"`

	// Inputs defines what kind of inputs are required for this component.
	Inputs Inputs `json:"inputs" yaml:"inputs"`
	// Validation can define validations of this runtime.
	// Validations *Validations `json:"validations"       yaml:"validations"`

	// Outputs can explicitly specify what kind of outputs this component will create.
	Outputs Outputs `json:"outputs" yaml:"outputs"`
}
