package cloudfront

import (
	"errors"
	"regexp"
	"strings"
)

func GetLogVersion(data *[]byte) (*string, error) {
	versionPattern := `#Version:\s\d*`
	re := regexp.MustCompile(versionPattern)
	match := re.FindString(string(*data))
	if match != "" {
		return &strings.Fields(match)[1], nil
	}
	return &match, errors.New("Log format is invalid")

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
