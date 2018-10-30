// Copyright © 2018 The Pingaling Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pingaling

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"unicode"

	"github.com/ghodss/yaml"
)

const separator = "---"

// external functions
var (
	decoYAMLtoJSON    = yaml.YAMLToJSON
	decoJSONUnmarshal = json.Unmarshal
	decoYAMLUnmarshal = yaml.Unmarshal
)

// Reader interface for Read()
type Reader interface {
	Read() ([]byte, error)
}

// YAMLReader struct for reader
type YAMLReader struct {
	reader Reader
}

// NewYAMLReader returns YAMLReader struct
func NewYAMLReader(r *bufio.Reader) *YAMLReader {
	return &YAMLReader{
		reader: &LineReader{reader: r},
	}
}

// Read returns a full YAML document.
func (r *YAMLReader) Read() ([]byte, error) {
	var buffer bytes.Buffer
	for {
		line, err := r.reader.Read()
		if err != nil && err != io.EOF {
			return nil, err
		}

		sep := len([]byte(separator))
		if i := bytes.Index(line, []byte(separator)); i == 0 {
			// We have a potential document terminator
			i += sep
			after := line[i:]
			if len(strings.TrimRightFunc(string(after), unicode.IsSpace)) == 0 {
				if buffer.Len() != 0 {
					return buffer.Bytes(), nil
				}
				if err == io.EOF {
					return nil, err
				}
			}
		}
		if err == io.EOF {
			if buffer.Len() != 0 {
				// If we're at EOF, we have a final, non-terminated line. Return it.
				return buffer.Bytes(), nil
			}
			return nil, err
		}
		buffer.Write(line)
	}
}

// LineReader struct
type LineReader struct {
	reader *bufio.Reader
}

// Read returns a single line (with '\n' ended) from the underlying reader.
// An error is returned iff there is an error with the underlying reader.
func (r *LineReader) Read() ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line     []byte
		buffer   bytes.Buffer
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.reader.ReadLine()
		buffer.Write(line)
	}
	buffer.WriteByte('\n')
	return buffer.Bytes(), err
}

// SplitYAMLDocuments read multiple YAML docs to []TypeMeta struct
func SplitYAMLDocuments(ymlBytes []byte) ([]TypeMeta, error) {
	buf := bytes.NewBuffer(ymlBytes)
	reader := NewYAMLReader(bufio.NewReader(buf))
	docs := make([]TypeMeta, 0)
	for {
		// Read one YAML document at a time, until io.EOF is returned
		typeMetaInfo := TypeMeta{}
		b, err := reader.Read()
		if err == io.EOF {
			break
		}
		if len(b) == 0 {
			break
		}
		// Deserialize the TypeMeta information of this byte slice

		if err := decoYAMLUnmarshal(b, &typeMetaInfo); err != nil {
			return nil, &ErrNotExpectedYAML{
				OriginalBody: string(b),
				Err:          err,
			}
		}
		docs = append(docs, typeMetaInfo)
	}

	return docs, nil
}

// YAMLDecoder unmarshal []byte to struct
func YAMLDecoder(b []byte, into interface{}) error {
	toJSON, err := decoYAMLtoJSON(b)
	if err != nil {
		return &ErrNotExpectedYAML{
			OriginalBody: string(b),
			Err:          err,
		}
	}
	err = decoJSONUnmarshal(toJSON, into)
	if err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: string(b),
			Err:          err,
		}
	}
	return nil
}

// JSONDecoder decode response into target struct
func JSONDecoder(b bytes.Buffer, into interface{}) error {

	if err := json.NewDecoder(&b).Decode(into); err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: b.String(),
			Err:          err,
		}
	}
	return nil
}
