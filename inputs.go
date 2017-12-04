package protocol

// Inputs ...
type Inputs map[string]*Input

// Input ...
type Input struct {

	// Type specifies the type of this "Input", either of [file]
	Type string `json:"type" yaml:"type"`

	// Extensions represents file extensions this "Input" can accept.
	// This properties must appear only when "Type" == "file"
	Extensions []string `json:"ext" yaml:"ext"`

	// Following properties are only used for "Parameters".
	// It means they must not appear in "Component" definition.
	URL string `json:"url"  yaml:"url"`
}
