package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sho-he/aws-cloudfront-log-formatter/app/adapters"
)

var (
	path   = flag.String("f", "", "file path")
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
	if *path == "" {
		return errors.New("file path (-f) option is required")
	}
	return nil
}

func call() {
	data := readFile(*path)
	fmt.Printf("Log Versin: %s\n", *getLogVersion(&data))
	output := createFile(genOutputPath())
	r := adapters.NewRouter()
	r.Run(*format, getFields(&data), getRows(&data), output)
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func createFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func genOutputPath() string {
	return *path + "." + *format
}

func getLogVersion(data *string) *string {
	versionPattern := `#Version:\s\d*`
	re := regexp.MustCompile(versionPattern)
	return &strings.Fields(re.FindString(*data))[1]
}

func getFields(data *string) *[]string {
	fieldsPattern := `#Fields:\s.*\n`
	re := regexp.MustCompile(fieldsPattern)
	fields := strings.Fields(re.FindString(*data))[1:]
	return &fields
}

func getRows(data *string) *[]string {
	rows := strings.Split(*data, "\n")[2:]
	return &rows
}
