package oparbac

import (
	"context"
	_ "embed"
	"testing"
)

var data = Data{
	UserRoles: map[string][]string{
		"alice": {
			"admin",
		},
		"bob": {
			"customer",
			"billing",
			"employee",
		},
		"eve": {
			"customer",
		},
	},
	RoleGrants: map[string][]map[string]string{
		"customer": {
			{
				"resource": "user.get",
			},
		},
		"employee": {
			{
				"resource": "user.list",
			},
			{
				"resource": "user.get",
			},
			{
				"resource": "user.delete",
			},
		},
		"billing": {
			{
				"resource": "payment.*",
			},
		},
	},
}

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
			got, err := Allow(context.TODO(), data, test.input)
			if err != nil {
				t.Error(err)
			}
			if got != test.want {
				t.Errorf("want: %v, got: %v", test.want, got)
			}
		})
	}
}
