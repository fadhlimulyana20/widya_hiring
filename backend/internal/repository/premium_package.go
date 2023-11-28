package repository

import (
	"fmt"

	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PremiumPackageRepo struct {
	db   *gorm.DB
	name string
}

type PremiumPackageRepository interface {
	// Create a new Premium Package
	Create(pp entities.PremiumPackage) (entities.PremiumPackage, error)
	// Update role
	Update(pp entities.PremiumPackage) (entities.PremiumPackage, error)
	// List role
	List(param params.PremiumPackageFilterParam) ([]entities.PremiumPackage, int, error)
	// Get Role
	Get(ID int) (entities.PremiumPackage, error)
	// Delete Role
	Delete(ID int) (entities.PremiumPackage, error)
}

// Create new role repository instance
func NewPremiumPackageRepository(db *gorm.DB) PremiumPackageRepository {
	return &PremiumPackageRepo{
		db:   db,
		name: "Premium Package Repository",
	}
}

func (r *PremiumPackageRepo) Create(pp entities.PremiumPackage) (entities.PremiumPackage, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", r.name))

	if err := r.db.Create(&pp).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return pp, err
	}

	return pp, nil
}

func (r *PremiumPackageRepo) Get(ID int) (entities.PremiumPackage, error) {
	log.Info(fmt.Sprintf("[%s][Get] is executed", r.name))
	var pp entities.PremiumPackage

	db := r.db

	if err := db.Debug().First(&pp, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", r.name, err.Error()))
		return pp, err
	}

	return pp, nil
}

func (r *PremiumPackageRepo) List(param params.PremiumPackageFilterParam) ([]entities.PremiumPackage, int, error) {
	log.Info(fmt.Sprintf("[%s][List] is executed", r.name))

	var pps []entities.PremiumPackage

	var count int64
	r.db.Find(&pps).Count(&count)

	db := r.db

	if param.UserID != 0 {
		db = db.Where("user_id = ?", param.UserID)
	}

	if *param.IsActive {
		db.Where("is_active", *param.IsActive)
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&pps).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", r.name, err.Error()))
		return pps, int(count), err
	}

	return pps, int(count), nil
}

func (r *PremiumPackageRepo) Update(pp entities.PremiumPackage) (entities.PremiumPackage, error) {
	log.Info(fmt.Sprintf("[%s][Update] is executed", r.name))

	if err := r.db.Model(&pp).Updates(&pp).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", r.name, err.Error()))
		return pp, err
	}

	return pp, nil
}

func (r *PremiumPackageRepo) Delete(ID int) (entities.PremiumPackage, error) {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", r.name))
	var pp entities.PremiumPackage

	if err := r.db.Delete(&pp, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", r.name, err.Error()))
		return pp, err
	}

	return pp, nil
}
