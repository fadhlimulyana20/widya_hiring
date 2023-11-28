package usecase

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gorm.io/gorm"
)

type userPoint struct {
	userPointRepo repository.UserPointRepository
	name          string
}

type UserPointUsecase interface {
	Get(userID int) appctx.Response
	GetList(param params.UserPointFilterParam) appctx.Response
}

func NewUserPointUsecase(db *gorm.DB) UserPointUsecase {
	return &userPoint{
		userPointRepo: repository.NewUserPointRepository(db),
		name:          "User Point Usecase",
	}
}

func (u *userPoint) Get(userID int) appctx.Response {
	up, err := u.userPointRepo.GetByUser(userID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(up)
}

func (u *userPoint) GetList(param params.UserPointFilterParam) appctx.Response {
	ups, count, err := u.userPointRepo.GetListOfUserPoint(param)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetList] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(ups).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}
