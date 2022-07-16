# opa-rbac

Role-based access control (RBAC) with the Open Policy Agent.

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


