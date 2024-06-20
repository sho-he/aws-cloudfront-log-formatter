package usecases

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"

	cf "github.com/sho-he/aws-cloudfront-log-formatter/adapters/cloudfront"
	"github.com/sho-he/aws-cloudfront-log-formatter/usecases/ports"
)

type JsonInteractor struct {
	fileoperator ports.FileOperator
}

func NewJsonInteractor(
	fileoperator ports.FileOperator,
) *JsonInteractor {
	return &JsonInteractor{
		fileoperator,
	}
}

func (j *JsonInteractor) Call() error {
	data, err := j.fileoperator.ReadFile()
	if err != nil {
		return err
	}
	format := cf.NewCloudfrontLogFormat(data)
	ver, err := format.GetLogVersion()
	if err != nil {
		return err
	}
	fmt.Printf("Cloudfront Log Format Versin: %s\n", *ver)
	fields, err := format.GetFields()
	if err != nil {
		return err
	}
	rows, err := format.GetRows()
	if err != nil {
		return err
	}
	output, err := j.fileoperator.CreateFile()
	slog.Info("=== JSON Conversion Start ===")
	if err := j.Convert(output, fields, rows); err != nil {
		return err
	}
	slog.Info("=== JSON Conversion Complete ===")
	return nil
}

func (j *JsonInteractor) Convert(output *os.File, fields *[]string, rows *[]string) error {
	_f := *fields
	objs := []map[string]string{}
	for _, row := range *rows {
		_obj := map[string]string{}
		for j, elm := range strings.Fields(row) {
			_obj[_f[j]] = elm

		}
		objs = append(objs, _obj)
	}
	bytes, err := json.Marshal(objs)
	if err != nil {
		return err
	}
	if _, err := output.Write(bytes); err != nil {
		return err
	}
	return nil
}
