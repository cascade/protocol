package protocol

// Job ...
type Job struct {
	Machine    *Machine         `json:"machine"         yaml:"machine"`
	Definition *Definition      `json:"definition"      yaml:"definition"`
	Inputs     map[string]Input `json:"inputs"          yaml:"inputs"`
	Steps      []Job            `json:"steps,omitempty" yaml:"steps,omitempty"`
}

// Definition ...
type Definition struct {
	Image  *string `json:"image"  yaml:"image"`
	Script *File   `json:"script" yaml:"script"`
	Pod    *string `json:"pod"    yaml:"pod"`
}
