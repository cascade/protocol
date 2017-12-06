package protocol

// Runtime defines a machine runtime such as OS, dependencies and language versions.
// If you want to use "local" for the runtime, just omit "runtime" field.
type Runtime struct {
	// Image is a OS image on which "Execute" is executed, unless "Execute" doesn't have "Container"
	// If you want to use "local" for the runtime, just omit "runtime" field or keep this field empty.
	Image string `json:"image"   yaml:"image"`
}
