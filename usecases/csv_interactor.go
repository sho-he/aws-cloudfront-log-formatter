package usecases

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"strings"

	cf "github.com/sho-he/aws-cloudfront-log-formatter/adapters/cloudfront"
	"github.com/sho-he/aws-cloudfront-log-formatter/usecases/ports"
)

type CsvInteractor struct {
	fileoperator ports.FileOperator
}

func NewCsvInteractor(
	fileoperator ports.FileOperator,
) *CsvInteractor {
	return &CsvInteractor{
		fileoperator,
	}
}

func (c *CsvInteractor) Call() error {
	data, err := c.fileoperator.ReadFile()
	if err != nil {
		return err
	}
	format := cf.NewCloudfrontLogFormat(data)
	ver, err := format.GetLogVersion()
	if err != nil {
		return err
	}
	fmt.Printf("Log Versin: %s\n", *ver)
	fields, err := format.GetFields()
	if err != nil {
		return err
	}
	slog.Info("[csv] Complete Extract header fields")
	rows, err := format.GetRows()
	if err != nil {
		return err
	}
	slog.Info("[csv] Complete Extract rows")

	output, err := c.fileoperator.CreateFile()
	if err := c.Convert(csv.NewWriter(output), fields, rows); err != nil {
		return err
	}
	return nil
}

func (c *CsvInteractor) Convert(output *csv.Writer, fields *[]string, rows *[]string) error {
	defer output.Flush()
	slog.Info("[csv] Conversion start")
	if err := output.Write(*fields); err != nil {
		return err
	}
	slog.Info("[csv] Header was inserted.")
	for _, row := range *rows {
		if err := output.Write(strings.Fields(row)); err != nil {
			return err
		}
	}
	slog.Info("[csv] All rows were inserted.")
	return nil
}
