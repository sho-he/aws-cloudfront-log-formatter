package local

import "os"

type Reader struct {
	input string
}

func NewReader(input string) *Reader {
	return &Reader{
		input,
	}
}

func (r *Reader) ReadFile() (*[]byte, error) {
	data, err := os.ReadFile(r.input)
	return &data, err
}
