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
	data := `
{
  "user_roles": {
    "alice": [
      "admin"
    ],
    "bob": [
      "employee"
    ],
    "eve": [
      "customer"
    ]
  },
  "role_grants": {
    "customer": [
      {
        "resource": "customer.get"
      }
    ],
    "employee": [
      {
        "resource": "employee.list"
      },
      {
        "resource": "employee.get"
      }
    ]
  }
}
`
	fmt.Println(oparbac.Allow(context.TODO(), []byte(data), oparbac.Input{User: "alice", Resource: "employee.get"}))
}
```


