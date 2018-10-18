package pingaling

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"

	"github.com/ghodss/yaml"
)

const separator = "---"

type Reader interface {
	Read() ([]byte, error)
}

type YAMLReader struct {
	reader Reader
}

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

type LineReader struct {
	reader *bufio.Reader
}

// Read returns a single line (with '\n' ended) from the underlying reader.
// An error is returned iff there is an error with the underlying reader.
func (r *LineReader) Read() ([]byte, error) {
	var (
		isPrefix bool  = true
		err      error = nil
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
		if err := yaml.Unmarshal(b, &typeMetaInfo); err != nil {
			return nil, err
		}
		docs = append(docs, typeMetaInfo)
	}

	return docs, nil
}
