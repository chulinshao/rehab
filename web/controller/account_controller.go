package controller

import (
	"github.com/chulinshao/rehab/common"
	"github.com/chulinshao/rehab/service"
	"github.com/shopspring/decimal"
)

type AccountController struct {
	Service service.AccountService
}

func (c AccountController) GetByCode(code string) common.Result {
	result := common.GetResult()

	common.IsNull(&result, code, "code")
	doctorAccount := c.Service.GetByCode(code)

	dptc, tjtc := c.Service.CountTotalCommentary(code)
	doctorAccount.EvaluateCommentary = decimal.NewFromFloat(dptc)
	doctorAccount.RefereeCommentary = decimal.NewFromFloat(tjtc)
	result.Data = doctorAccount
	return result
}
