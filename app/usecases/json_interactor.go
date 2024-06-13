package usecases

import "fmt"

type JsonInteractor struct {
	fields *[]string
	rows   *[]string
}

func NewJsonInteractor(
	fields *[]string,
	rows *[]string,
) *JsonInteractor {
	return &JsonInteractor{
		fields,
		rows,
	}
}

func (c *JsonInteractor) Call() error {
	fmt.Printf("Fields: %s\n", *c.fields)
	fmt.Printf("Rows: %s\n", *c.rows)
	return nil
}
