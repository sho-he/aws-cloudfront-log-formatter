package cloudfront

import (
	"regexp"
	"strings"
)

func GetLogVersion(data *[]byte) *string {
	versionPattern := `#Version:\s\d*`
	re := regexp.MustCompile(versionPattern)
	return &strings.Fields(re.FindString(string(*data)))[1]
}

func GetFields(data *[]byte) *[]string {
	fieldsPattern := `#Fields:\s.*\n`
	re := regexp.MustCompile(fieldsPattern)
	fields := strings.Fields(re.FindString(string(*data)))[1:]
	return &fields
}

func GetRows(data *[]byte) *[]string {
	rows := strings.Split(string(*data), "\n")[2:]
	return &rows
}
