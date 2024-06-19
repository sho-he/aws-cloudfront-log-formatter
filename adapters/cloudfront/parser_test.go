package cloudfront

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	cf "github.com/sho-he/aws-cloudfront-log-formatter/adapters/cloudfront"
)

func TestGetLogVersion(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"the log format is as expected",
			args{
				[]byte("#Version: 1\n#Fields: field1 field2\ndata1 data2"),
			},
			"1",
			false,
		},
		{
			"the log format is not as expected",
			args{
				[]byte("#HogeHoge: 1\n#Fields: field1 field2\ndata1 data2"),
			},
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cf.GetLogVersion(&tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLogVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(*got, tt.want); diff != "" {
				t.Errorf("GetLogVersion() got = %v, want %v. diff %v", *got, tt.want, diff)
			}

		})
	}

}
