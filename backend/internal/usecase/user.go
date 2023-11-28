package usecase

import (
	"errors"
	"fmt"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gitlab.com/project-quiz/utils/password"
	"gitlab.com/project-quiz/utils/postgres"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type user struct {
	repo repository.UserRepository
	name string
}

type UserUsecase interface {
	// Create a new user record
	Create(param params.UserCreateParam) appctx.Response

	// List and filter user record
	List(params.UserListParams) appctx.Response

	// Update user record
	Update(params.UserUpdateParam) appctx.Response

	// Update user record
	Get(int) appctx.Response

	// Delete user record
	Delete(int) appctx.Response

	// Update user password
	UpdatePassword(params.UserUpdatePassword) appctx.Response
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &user{
		repo: repository.NewUserRepository(db),
		name: "USER USECASE",
	}
}

func (u *user) Create(param params.UserCreateParam) appctx.Response {
	var user entities.User
	copier.Copy(&user, &param)

	// Generate Hash Password
	passwd, _ := password.HashPassword(user.Password)
	user.Password = passwd

	usr, err := u.repo.Create(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		httpStatus := postgres.TranslatePostgresError(err)
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(httpStatus)
	}

	return *appctx.NewResponse().WithData(usr)
}

func (u *user) List(param params.UserListParams) appctx.Response {
	var usrs []entities.User
	users, count, err := u.repo.List(usrs, param)
	if err != nil {
		log.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(users).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (u *user) Update(param params.UserUpdateParam) appctx.Response {
	var user entities.User
	user, err := u.repo.Get(user, param.ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	copier.Copy(&user, &param)

	usr, err := u.repo.Update(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}

func (u *user) Get(ID int) appctx.Response {
	var user entities.User
	user, err := u.repo.Get(user, ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Get] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(user)
}

func (u *user) Delete(ID int) appctx.Response {
	var user entities.User
	user, err := u.repo.Get(user, ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(user)
}

func (u *user) UpdatePassword(param params.UserUpdatePassword) appctx.Response {
	// Get user data
	var user entities.User
	user, err := u.repo.Get(user, param.UserID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][UpdatePassword] %s", u.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors("User not found").WithCode(401)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	// Check Hash Password
	if match := password.CheckPasswordHash(param.OldPassword, user.Password); !match {
		log.Error(fmt.Sprintf("[%s]UpdatePassword] %s", u.name, "password not match"))
		return *appctx.NewResponse().WithErrors("Wrong password").WithCode(401)
	}

	// Generate Hash Password
	passwd, _ := password.HashPassword(param.NewPassword)
	user.Password = passwd

	usr, err := u.repo.Update(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][UpdatePassword] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}
