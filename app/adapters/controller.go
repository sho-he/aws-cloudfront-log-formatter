package adapters

import (
	"os"

	"github.com/cockroachdb/errors"
	"github.com/sho-he/aws-cloudfront-log-formatter/app/usecases"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Run(
	format string,
	fields *[]string,
	rows *[]string,
	output *os.File,
) error {
	switch format {
	case "csv":
		return usecases.NewCsvInteractor(fields, rows, output).Call()
	case "json":
		return usecases.NewJsonInteractor(fields, rows, output).Call()
	default:
		return errors.Newf("invalid format: %s", format)
	}
}
