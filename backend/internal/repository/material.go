package repository

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"
	"gorm.io/gorm"
)

type materialRepo struct {
	db   *gorm.DB
	name string
}

type MaterialRepository interface {
	// Create a new role
	Create(material entities.Material) (entities.Material, error)
	// Update role
	Update(material entities.Material) (entities.Material, error)
	// List role
	List(param params.MaterialFilterParam) ([]entities.Material, int, error)
	// Get Total
	GetTotal() (int, error)
	// Get Role
	Get(ID int) (entities.Material, error)
	// Delete Role
	Delete(ID int) (entities.Material, error)
}

// Create new role repository instance
func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepo{
		db:   db,
		name: "Material Repository",
	}
}

func (m *materialRepo) Create(material entities.Material) (entities.Material, error) {
	if err := m.db.Create(&material).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", m.name, err.Error()))
		return material, err
	}

	return material, nil
}

func (m *materialRepo) Get(ID int) (entities.Material, error) {
	var material entities.Material

	db := m.db

	if err := db.Debug().First(&material, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", m.name, err.Error()))
		return material, err
	}

	return material, nil
}

func (m *materialRepo) List(param params.MaterialFilterParam) ([]entities.Material, int, error) {
	var materials []entities.Material

	var count int64
	db := m.db

	if param.Q != "" {
		db = db.Where("LOWER(name) like ?", "%"+strings.ToLower(param.Q)+"%")
	}

	if param.Level != "" {
		db = db.Where("level = ?", param.Level)
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&materials).Count(&count).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", m.name, err.Error()))
		return materials, int(count), err
	}

	return materials, int(count), nil
}

func (m *materialRepo) GetTotal() (int, error) {
	log.Info(fmt.Sprintf("[%s][Get Total] is executed", m.name))

	var count int64
	var material entities.Material

	if err := m.db.Model(&material).Count(&count).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Get Total] %s", m.name, err.Error()))
		return 0, err
	}

	return int(count), nil
}

func (m *materialRepo) Update(material entities.Material) (entities.Material, error) {
	if err := m.db.Model(&material).Updates(&material).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", m.name, err.Error()))
		return material, err
	}

	return material, nil
}

func (m *materialRepo) Delete(ID int) (entities.Material, error) {
	var material entities.Material

	if err := m.db.Debug().Delete(&material, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", m.name, err.Error()))
		return material, err
	}

	return material, nil
}
