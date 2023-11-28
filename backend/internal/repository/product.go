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

type productRepo struct {
	db   *gorm.DB
	name string
}

type ProductRepository interface {
	// Create a new product
	Create(product entities.Product) (entities.Product, error)
	// Update product
	Update(product entities.Product) (entities.Product, error)
	// List product
	List(param params.ProductFilterParam) ([]entities.Product, int, error)
	// Get product
	Get(ID int) (entities.Product, error)
	// Delete product
	Delete(ID int) (entities.Product, error)
}

// Create new role repository instance
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{
		db:   db,
		name: "Product Repository",
	}
}

func (m *productRepo) Create(product entities.Product) (entities.Product, error) {
	if err := m.db.Create(&product).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", m.name, err.Error()))
		return product, err
	}

	return product, nil
}

func (m *productRepo) Get(ID int) (entities.Product, error) {
	var product entities.Product

	db := m.db

	if err := db.Debug().First(&product, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", m.name, err.Error()))
		return product, err
	}

	return product, nil
}

func (m *productRepo) List(param params.ProductFilterParam) ([]entities.Product, int, error) {
	var products []entities.Product

	var count int64
	useFilterCount := false
	db := m.db
	db.Find(&products).Count(&count)

	if param.Q != "" {
		db = db.Where("LOWER(pname) like ?", "%"+strings.ToLower(param.Q)+"%")
		useFilterCount = true
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&products).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", m.name, err.Error()))
		return products, int(count), err
	}

	if useFilterCount {
		db.Count(&count)
	}

	return products, int(count), nil
}

func (m *productRepo) Update(product entities.Product) (entities.Product, error) {
	if err := m.db.Model(&product).Updates(&product).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", m.name, err.Error()))
		return product, err
	}

	return product, nil
}

func (m *productRepo) Delete(ID int) (entities.Product, error) {
	var product entities.Product

	if err := m.db.Debug().Delete(&product, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", m.name, err.Error()))
		return product, err
	}

	return product, nil
}
