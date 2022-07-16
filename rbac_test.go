package oparbac

import (
	"context"
	_ "embed"
	"testing"
)

//go:embed sample_data.json
var dataBytes []byte

func TestAllow(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  bool
	}{
		{
			name: "admin allowed payment.*",
			input: Input{
				User:     "alice",
				Resource: "payment.DELETE",
			},
			want: true,
		},
		{
			name: "bob allowed user.get",
			input: Input{
				User:     "bob",
				Resource: "user.GET",
			},
			want: true,
		},
		{
			name: "bob allowed user.list",
			input: Input{
				User:     "bob",
				Resource: "user.list",
			},
			want: true,
		},
		{
			name: "bob allowed user.delete",
			input: Input{
				User:     "bob",
				Resource: "user.delete",
			},
			want: true,
		},
		{
			name: "bob not allowed user.create",
			input: Input{
				User:     "bob",
				Resource: "user.create",
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Allow(context.TODO(), dataBytes, test.input)
			if err != nil {
				t.Error(err)
			}
			if got != test.want {
				t.Errorf("want: %v, got: %v", test.want, got)
			}
		})
	}
}
