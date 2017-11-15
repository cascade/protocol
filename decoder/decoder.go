package decoder

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// NewDecoder ...
func NewDecoder(r io.Reader) Decoder {
	switch v := r.(type) {
	case *os.File:
		switch filepath.Ext(v.Name()) {
		case "json":
			return json.NewDecoder(r)
		default:
			return YamlDecoder{r}
		}
	default:
		return YamlDecoder{r}
	}
}

// Decode ...
func (d YamlDecoder) Decode(v interface{}) error {
	b, err := ioutil.ReadAll(d.r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, v)
}

// Decoder ...
type Decoder interface {
	Decode(interface{}) error
}

// YamlDecoder ...
type YamlDecoder struct {
	r io.Reader
}
