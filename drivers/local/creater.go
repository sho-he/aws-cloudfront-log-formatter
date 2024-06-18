package local

import "os"

type Creater struct {
	output string
}

func NewCreater(output string) *Creater {
	return &Creater{
		output,
	}
}

func (c *Creater) CreateFile() (*os.File, error) {
	f, err := os.Create(c.output)
	return f, err
}
