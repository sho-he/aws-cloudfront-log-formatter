package cloudfront

import (
	"errors"
	"regexp"
	"strings"
)

type cloudfrontLogFormat struct {
	data           *[]byte
	versionPattern string
	fieldsPattern  string
}

func NewCloudfrontLogFormat(data *[]byte) *cloudfrontLogFormat {
	return &cloudfrontLogFormat{
		data,
		`#Version:\s\d*`,
		`#Fields:\s.*\n`,
	}
}

func (c *cloudfrontLogFormat) GetLogVersion() (*string, error) {
	re := regexp.MustCompile(c.versionPattern)
	match := re.FindString(string(*c.data))
	if match != "" {
		return &strings.Fields(match)[1], nil
	}
	return &match, errors.New("Log format is invalid")

}

func (c *cloudfrontLogFormat) GetFields() (*[]string, error) {
	re := regexp.MustCompile(c.fieldsPattern)
	match := re.FindString(string(*c.data))
	if match != "" {
		fields := strings.Fields(match)[1:]
		return &fields, nil
	}
	return &[]string{}, errors.New("Log format is invalid")
}

func (c *cloudfrontLogFormat) GetRows() (*[]string, error) {
	rows := strings.Split(string(*c.data), "\n")[2:]
	return &rows, nil
}
