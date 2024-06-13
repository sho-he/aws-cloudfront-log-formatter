package usecases

import (
	"encoding/csv"
	"os"
	"strings"
)

type CsvInteractor struct {
	fields *[]string
	rows   *[]string
	output *csv.Writer
}

func NewCsvInteractor(
	fields *[]string,
	rows *[]string,
	output *os.File,
) *CsvInteractor {
	return &CsvInteractor{
		fields,
		rows,
		csv.NewWriter(output),
	}
}

func (c *CsvInteractor) Call() error {
	if err := c.ConvertFields(); err != nil {
		return err
	}
	if err := c.ConvertRows(); err != nil {
		return err
	}
	c.output.Flush()
	return nil
}

func (c *CsvInteractor) ConvertFields() error {
	if err := c.output.Write(*c.fields); err != nil {
		return err
	}
	return nil
}

func (c *CsvInteractor) ConvertRows() error {
	for _, row := range *c.rows {
		if err := c.output.Write(strings.Fields(row)); err != nil {
			return err
		}
	}
	return nil
}
