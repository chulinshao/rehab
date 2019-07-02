package controller

import (
	"github.com/chulinshao/rehab/common"
	"github.com/chulinshao/rehab/service"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
)

type AccountController struct {
	Service service.AccountService
}

func (ctl AccountController) GetByCode(c *gin.Context) {
	code := c.Param("code")
	result := common.GetResult()

	common.IsNull(&result, code, "code")
	doctorAccount := ctl.Service.GetByCode(code)

	dptc, tjtc := ctl.Service.CountTotalCommentary(code)
	doctorAccount.EvaluateCommentary = decimal.NewFromFloat(dptc)
	doctorAccount.RefereeCommentary = decimal.NewFromFloat(tjtc)
	result.Data = doctorAccount
	c.JSON(http.StatusOK, result)
}
