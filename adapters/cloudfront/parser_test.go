package cloudfront

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sho-he/aws-cloudfront-log-formatter/adapters/cloudfront"
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
			"the log format is not as expected (#Version: not found)",
			args{
				[]byte("#HogeHoge: 1\n#Fields: field1 field2\ndata1 data2"),
			},
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := cloudfront.NewCloudfrontLogFormat(&tt.args.data)
			got, err := cf.GetLogVersion()
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

func TestGetFields(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"the log format is as expected",
			args{
				[]byte("#Version: 1\n#Fields: field1 field2\ndata1 data2"),
			},
			[]string{"field1", "field2"},
			false,
		},
		{
			"the log format is not as expected (#Fields: not found)",
			args{
				[]byte("#Version: 1\n#HogeHoge: field1 field2\ndata1 data2"),
			},
			[]string{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := cloudfront.NewCloudfrontLogFormat(&tt.args.data)
			got, err := cf.GetFields()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(*got, tt.want); diff != "" {
				t.Errorf("GetFields() got = %v, want %v. diff %v", *got, tt.want, diff)
			}

		})
	}

}

func TestGetRows(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"the log format is as expected",
			args{
				[]byte("#Version: 1\n#Fields: field1 field2\ndata1 data2"),
			},
			[]string{"data1 data2"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := cloudfront.NewCloudfrontLogFormat(&tt.args.data)
			got, err := cf.GetRows()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(*got, tt.want); diff != "" {
				t.Errorf("GetRows() got = %v, want %v. diff %v", *got, tt.want, diff)
			}

		})
	}

}
