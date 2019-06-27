package controller

import (
	"github.com/chulinshao/rehab/common"
	"github.com/chulinshao/rehab/service"
)

type AccountController struct {
	Service service.AccountService
}

func (c AccountController) GetByCode(code string) common.Result {
	result := common.GetResult()

	common.IsNull(&result, code, "code")

	result.Data = c.Service.GetByCode(code)
	return result
}
