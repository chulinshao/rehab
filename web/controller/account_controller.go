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

// GET /account/code/:code
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

func (ctl AccountController) UpdateAlipayAccount(c *gin.Context)  {
	doctorCode := c.PostForm("doctorCode")
	alipayAccount := c.PostForm("alipayAccount")

	result := common.GetResult()
	err := common.IsNull(&result, doctorCode, "doctorCode")
	if err != nil {
		c.JSON(http.StatusOK, result)
		return
	}
	err = common.IsNull(&result, alipayAccount, "alipayAccount")
	if err != nil {
		c.JSON(http.StatusOK, result)
		return
	}
	ctl.Service.UpdateAlipayAccount(doctorCode, alipayAccount)
	c.JSON(http.StatusOK, result)
}

func (ctl AccountController) ApplyWithdrawals(c *gin.Context)  {

}
