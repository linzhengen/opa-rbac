# opa-rbac

Role-based access control (RBAC) with the Open Policy Agent.

[![Test](https://github.com/linzhengen/opa-rbac/actions/workflows/test.yml/badge.svg)](https://github.com/linzhengen/opa-rbac/actions/workflows/test.yml)
[![golangci-lint](https://github.com/linzhengen/opa-rbac/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/linzhengen/opa-rbac/actions/workflows/golangci-lint.yml)

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/linzhengen/opa-rbac"
)

func main() {
	var data = oparbac.Data{
		UserRoles: map[string][]string{
			"alice": {
				"admin",
			},
			"bob": {
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
			},
		},
	}
	input := oparbac.Input{User: "alice", Resource: "employee.get"}
	fmt.Println(oparbac.Allowed(context.TODO(), data, input))
}
```


