package usecase

import (
	"fmt"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gitlab.com/project-quiz/utils/random"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type premiumPackage struct {
	premiumPackageRepo repository.PremiumPackageRepository
	name               string
}

type PremiumPackageUsecase interface {
	// Create new package
	Create(param params.PremiumPackageCreateParam) appctx.Response
	// Edit a role
	Update(param params.PremiumPackageUpdateParam) appctx.Response
	// Get role list
	List(param params.PremiumPackageFilterParam) appctx.Response
	// Get detail role
	Detail(ID int) appctx.Response
	// Delete role
	Delete(ID int) appctx.Response
	// // Assign Role to user
	// Assign(userID int, roleName string) appctx.Response
	// // Revoke Role from user
	// Revoke(userID int, roleName string) appctx.Response
}

func NewPremiumPackageUsecase(db *gorm.DB) PremiumPackageUsecase {
	return &premiumPackage{
		premiumPackageRepo: repository.NewPremiumPackageRepository(db),
		name:               "Role Usecase",
	}
}

func (r *premiumPackage) Create(param params.PremiumPackageCreateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", r.name))

	var pp entities.PremiumPackage
	copier.Copy(&pp, &param)

	voucher, err := random.GenerateRandomVoucher(12)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}
	pp.Token = voucher

	pp, err = r.premiumPackageRepo.Create(pp)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(pp)
}

func (r *premiumPackage) Update(param params.PremiumPackageUpdateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Update] is executed", r.name))

	// get role
	var pp entities.PremiumPackage
	pp, err := r.premiumPackageRepo.Get(param.ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	copier.Copy(&pp, &param)

	pp, err = r.premiumPackageRepo.Update(pp)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(pp)
}

func (r *premiumPackage) List(param params.PremiumPackageFilterParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][List] is executed", r.name))

	// get role list
	var pps []entities.PremiumPackage
	pps, count, err := r.premiumPackageRepo.List(param)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(pps).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (r *premiumPackage) Detail(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Detail] is executed", r.name))

	// get role
	var pp entities.PremiumPackage
	pp, err := r.premiumPackageRepo.Get(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Detail] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(pp)
}

func (r *premiumPackage) Delete(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", r.name))

	// get role
	_, err := r.premiumPackageRepo.Delete(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("package deleted sucessfully")
}

// func (r *role) Assign(userID int, roleName string) appctx.Response {
// 	log.Info(fmt.Sprintf("[%s][Assign] is executed", r.name))

// 	var role entities.Role
// 	role, err := r.repo.GetByName(roleName)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	var user entities.User
// 	user, err = r.userRepo.Get(user, userID)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	_, err = r.userRepo.AddRole(user, role)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	return *appctx.NewResponse().WithMessage("role assigned sucessfully")
// }

// func (r *role) Revoke(userID int, roleName string) appctx.Response {
// 	log.Info(fmt.Sprintf("[%s][Revoke] is executed", r.name))

// 	var role entities.Role
// 	role, err := r.repo.GetByName(roleName)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	var user entities.User
// 	user, err = r.userRepo.Get(user, userID)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	_, err = r.userRepo.RemoveRole(user, role)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	return *appctx.NewResponse().WithMessage("role revoked sucessfully")
// }
