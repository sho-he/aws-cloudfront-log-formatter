package adapters

import (
	"github.com/cockroachdb/errors"
	"github.com/sho-he/aws-cloudfront-log-formatter/app/usecases"
)

type Router struct {
}

type InputPort interface {
	Call(
		fields *string,
		rows *string,
	)
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Run(
	format string,
	fields *[]string,
	rows *[]string,
) error {
	switch format {
	case "csv":
		return usecases.NewCsvInteractor(fields, rows).Call()
	case "json":
		return nil
	default:
		return errors.Newf("invalid format: %s", format)
	}
}
