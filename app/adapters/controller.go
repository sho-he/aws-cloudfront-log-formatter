package adapters

import (
	"github.com/cockroachdb/errors"
	"github.com/sho-he/aws-cloudfront-log-formatter/app/usecases"
	"github.com/sho-he/aws-cloudfront-log-formatter/app/usecases/ports"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Run(
	format string,
	fileoperator ports.FileOperator,
) error {
	switch format {
	case "csv":
		return usecases.NewCsvInteractor(fileoperator).Call()
	case "json":
		return nil
	default:
		return errors.Newf("invalid format: %s", format)
	}
}
