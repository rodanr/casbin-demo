package authz

import (
	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

func NewEnforcer(modelPath, policyPath string) (*casbin.Enforcer, error) {
	adapter := fileadapter.NewAdapter(policyPath)
	e, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, err
	}
	if err := e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}
