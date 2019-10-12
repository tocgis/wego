package controller

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/services"
)

type SqlController struct {
}

func (it *SqlController) Handle(ctx contracts.Context) (interface{}, error) {

	chain := services.Chain(
		&service.SqlService{},
	)
	_ = chain.Handle(ctx)

	ret := ctx.GetValue("user")
	return ret, nil

}
func (it *SqlController) Valid(ctx contracts.Context) error {
	return nil
}