package protocol

// Component is an unit of pipeline, with defining interface of itself.
type Component struct {
	Machine     *Machine     `json:"machine,omitempty" yaml:"machine"`
	Job         *Job         `json:"job"               yaml:"job"`
	Inputs      *Inputs      `json:"inputs"            yaml:"inputs"`
	Outputs     *Outputs     `json:"outputs"           yaml:"outputs"`
	Validations *Validations `json:"validations"       yaml:"validations"`
}
