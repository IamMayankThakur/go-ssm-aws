package ssm

import (
	"reflect"
	"testing"
)

func TestNewError(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		want    *Client
		wantErr bool
	}{
		{
			name: "disabled SSM test",
			cfg: Config{
				Enabled:     false,
				SecretsPath: "/test/secret",
				Region:      "us-east-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nil config error",
			cfg:     Config{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no region error",
			cfg: Config{
				Enabled:     false,
				SecretsPath: "/test/secret",
				Region:      "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := New(&test.cfg)
			if (err != nil) != test.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("New() got = %v, want %v", got, test.want)
			}
		})
	}
}

func TestNewSuccess(t *testing.T) {
	cfg := Config{
		Enabled:     true,
		SecretsPath: "/test/secret",
		Region:      "us-east-1",
	}
	_, err := New(&cfg)
	if err != nil {
		t.Errorf("error initlializing new client %v", err)
	}
}
