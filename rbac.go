package oparbac

import (
	"bytes"
	_ "embed"
	"encoding/json"

	"github.com/open-policy-agent/opa/storage/inmem"

	"context"
	"github.com/open-policy-agent/opa/rego"
)

//go:embed rbac.rego
var rbacRego []byte

type Input struct {
	User     string `json:"user"`
	Resource string `json:"resource"`
}

type Data struct {
	UserRoles  map[string][]string            `json:"user_roles"`
	RoleGrants map[string][]map[string]string `json:"role_grants"`
}

func Allow(ctx context.Context, data Data, input Input) (bool, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return false, err
	}

	q := rego.New(
		rego.Query("x := data.rbac.allow"),
		rego.Module("rbac.rego", string(rbacRego)),
		rego.Store(inmem.NewFromReader(&buf)),
		rego.Input(input),
	)

	rs, err := q.Eval(ctx)
	if err != nil {
		return false, err
	}
	return rs[0].Bindings["x"].(bool), nil
}
