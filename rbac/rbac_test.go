package rbac_test

import (
	"context"
	_ "embed"
	"testing"

	"github.com/linzhengen/rbacmw/rbac"
)

//go:embed sample_data.json
var dataBytes []byte

func TestAllow(t *testing.T) {
	tests := []struct {
		name  string
		input rbac.Input
		want  bool
	}{
		{
			name: "admin allowed payment.*",
			input: rbac.Input{
				User:     "alice",
				Resource: "payment.delete",
			},
			want: true,
		},
		{
			name: "bob allowed user.get",
			input: rbac.Input{
				User:     "bob",
				Resource: "user.get",
			},
			want: true,
		},
		{
			name: "bob allowed user.list",
			input: rbac.Input{
				User:     "bob",
				Resource: "user.list",
			},
			want: true,
		},
		{
			name: "bob allowed user.delete",
			input: rbac.Input{
				User:     "bob",
				Resource: "user.delete",
			},
			want: true,
		},
		{
			name: "bob not allowed user.create",
			input: rbac.Input{
				User:     "bob",
				Resource: "user.create",
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := rbac.Allow(context.TODO(), dataBytes, test.input)
			if err != nil {
				t.Error(err)
			}
			if got != test.want {
				t.Errorf("want: %v, got: %v", test.want, got)
			}
		})
	}
}
