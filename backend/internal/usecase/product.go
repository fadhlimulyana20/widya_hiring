package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gorm.io/gorm"
)

type product struct {
	productRepo repository.ProductRepository
	name        string
}

type ProductUsecase interface {
	// Create
	Create(param params.ProductCreateParam) appctx.Response
	Update(param params.ProductUpdateParam) appctx.Response
	// Get list of material
	List(param params.ProductFilterParam) appctx.Response
	// Get detail of material
	Detail(ID int) appctx.Response
	Delete(ID int) appctx.Response
}

func NewProductUsecase(db *gorm.DB) ProductUsecase {
	return &product{
		productRepo: repository.NewProductRepository(db),
		name:        "Product Usecase",
	}
}

func (p *product) Create(param params.ProductCreateParam) appctx.Response {
	var product entities.Product
	copier.Copy(&product, &param)

	product, err := p.productRepo.Create(product)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", p.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(product)
}

func (p *product) Update(param params.ProductUpdateParam) appctx.Response {
	var product entities.Product
	copier.Copy(&product, &param)

	product, err := p.productRepo.Update(product)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", p.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(product)
}

func (p *product) List(param params.ProductFilterParam) appctx.Response {
	products, count, err := p.productRepo.List(param)
	if err != nil {
		logrus.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(products).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (p *product) Detail(ID int) appctx.Response {
	product, err := p.productRepo.Get(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Detail] %s", p.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusNotFound)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(product)
}

func (p *product) Delete(ID int) appctx.Response {
	_, err := p.productRepo.Delete(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete] %s", p.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusNotFound)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithMessage("Product deleted successfully")
}
