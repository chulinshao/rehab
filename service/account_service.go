package service

import (
	"github.com/chulinshao/rehab/common"
	"github.com/chulinshao/rehab/dao"
	"github.com/chulinshao/rehab/datasource"
	"github.com/chulinshao/rehab/models"
)

type AccountService interface {

	// 查询医生详情
	GetByCode(code string) models.DoctorAccount

	// 查询总提成
	// 返回点评提成，推荐提成
	CountTotalCommentary(doctorCode string) (float64, float64)
}

func NewAccountService() AccountService {
	return &accountService{
		dao:           dao.NewAccountDao(datasource.GetEngine()),
		accountLogDao: dao.NewDoctorAccountLogDao(datasource.GetEngine()),
	}
}

type accountService struct {
	dao           *dao.AccountDao
	accountLogDao *dao.DoctorAccountLogDao
}

func (s accountService) GetByCode(code string) models.DoctorAccount {
	return s.dao.GetByCode(code)
}

func (s accountService) CountTotalCommentary(doctorCode string) (float64, float64) {
	evaluateCommentary := s.accountLogDao.CountTotalCommentary(doctorCode, common.EVALUATE)
	refereeCommentary := s.accountLogDao.CountTotalCommentary(doctorCode, common.REFEREE)
	return evaluateCommentary, refereeCommentary
}
