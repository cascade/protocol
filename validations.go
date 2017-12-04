package protocol

// Validations ...
type Validations []*Validation

// Validation ...
type Validation struct {
	Inline  string `json:"inline" yaml:"inline"`
	Message string `json:"message" yaml:"message"`
}
