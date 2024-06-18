package main

import (
	"errors"
	"flag"

	"github.com/sho-he/aws-cloudfront-log-formatter/adapters"
	"github.com/sho-he/aws-cloudfront-log-formatter/adapters/fileoperators"
	"github.com/sho-he/aws-cloudfront-log-formatter/drivers/local"
)

var (
	ipath  = flag.String("f", "", "input file path")
	opath  = flag.String("o", "", "output file path")
	format = flag.String("t", "csv", "conversion format type")
)

func main() {
	flag.Parse()
	if err := valid(); err != nil {
		panic(err)
	}
	call()
}

func valid() error {
	if *ipath == "" {
		return errors.New("file path (-f) option is required")
	}
	if *opath == "" {
		return errors.New("file path (-o) option is required")
	}
	return nil
}

func call() {
	r := adapters.NewRouter()
	r.Run(
		*format,
		fileoperators.NewFileOperator(
			local.NewReader(*ipath),
			local.NewCreater(*opath),
		),
	)
}
