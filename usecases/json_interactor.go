package usecases

import (
	"fmt"
	"os"
)

type JsonInteractor struct {
	fields *[]string
	rows   *[]string
	output *os.File
}

func NewJsonInteractor(
	fields *[]string,
	rows *[]string,
	output *os.File,
) *JsonInteractor {
	return &JsonInteractor{
		fields,
		rows,
		output,
	}
}

func (j *JsonInteractor) Call() error {
	fmt.Printf("Fields: %s\n", *j.fields)
	fmt.Printf("Rows: %s\n", *j.rows)
	return nil
}

func (j *JsonInteractor) convertFields() error {
	return nil
}
