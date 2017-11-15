package protocol

// FileProvider ...
type FileProvider string

const (
	// FileLocal ...
	FileLocal FileProvider = "local"
	// FilePublicHTTP ...
	FilePublicHTTP FileProvider = "http"
	// // FileS3 ...
	// FileS3 FileProvider = "s3"
	// // FileGCS ...
	// FileGCS FileProvider = "gcs"
)

// File represents file identity.
type File struct {
	// Provider represents local fs or over http.
	// If omitted, it's gonna determined by parsing "url".
	Provider *FileProvider `json:"provider,omitempty" yaml:"provider,omitempty"`
	// URL (required): is http(s) URL or path to local file.
	// (ex1) https://storage.googleapis.com/hgc-otiai10-test/example.png
	// (ex2) ./path/to/image/example.png
	URL string `json:"url" yaml:"url"`
}
