package usecases

import "fmt"

type CsvInteractor struct {
	fields *[]string
	rows   *[]string
}

func NewCsvInteractor(
	fields *[]string,
	rows *[]string,
) *CsvInteractor {
	return &CsvInteractor{
		fields,
		rows,
	}
}

func (c *CsvInteractor) Call() error {
	fmt.Printf("Fields: %s\n", *c.fields)
	fmt.Printf("Rows: %s\n", *c.rows)
	return nil
}
