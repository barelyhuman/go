// Updated version of the following
// https://github.com/tj/front/blob/739be213b0a1c496dccaf9e5df1514150c9548e4/front.go
package front

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

var delim = []byte("---")

func Unmarshal(b []byte, v interface{}) (content []byte, err error) {
	if !bytes.HasPrefix(b, delim) {
		return b, nil
	}
	parts := bytes.SplitN(b, delim, 3)
	content = parts[2]
	err = yaml.Unmarshal(parts[1], v)
	return
}
