package rbac

import (
	_ "embed"

	"github.com/open-policy-agent/opa/storage/inmem"

	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/util"

	"context"
)

//go:embed rbac.rego
var rbacRego []byte

type Input struct {
	User     string `json:"user"`
	Resource string `json:"resource"`
}

func Allow(ctx context.Context, data []byte, input Input) (bool, error) {
	var json map[string]interface{}
	if err := util.Unmarshal(data, &json); err != nil {
		return false, err
	}
	q := rego.New(
		rego.Query("x := data.rbac.allow"),
		rego.Module("rbac.rego", string(rbacRego)),
		rego.Store(inmem.NewFromObject(json)),
		rego.Input(input),
	)

	rs, err := q.Eval(ctx)
	if err != nil {
		return false, err
	}
	return rs[0].Bindings["x"].(bool), nil
}
