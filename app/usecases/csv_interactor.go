package usecases

import (
	"fmt"
	"os"
)

type CsvInteractor struct {
	fields *[]string
	rows   *[]string
	output *os.File
}

func NewCsvInteractor(
	fields *[]string,
	rows *[]string,
	output *os.File,
) *CsvInteractor {
	return &CsvInteractor{
		fields,
		rows,
		output,
	}
}

func (c *CsvInteractor) Call() error {
	fmt.Printf("Fields: %s\n", *c.fields)
	fmt.Printf("Rows: %s\n", *c.rows)
	return nil
}

func (c *CsvInteractor) ConvertFields() error {
	return nil
}
