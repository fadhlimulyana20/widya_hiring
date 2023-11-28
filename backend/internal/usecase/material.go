package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gorm.io/gorm"
)

type material struct {
	materialRepo repository.MaterialRepository
	name         string
}

type MaterialUsecase interface {
	// Create
	Create(param params.MaterialCreateParam) appctx.Response
	Update(param params.MaterialEditParam) appctx.Response
	// Get list of material
	List(param params.MaterialFilterParam) appctx.Response
	// Get detail of material
	Detail(ID int) appctx.Response
	Delete(ID int) appctx.Response
}

func NewMaterialUsecase(db *gorm.DB) MaterialUsecase {
	return &material{
		materialRepo: repository.NewMaterialRepository(db),
		name:         "Material Usecase",
	}
}

func (m *material) Create(param params.MaterialCreateParam) appctx.Response {
	var material entities.Material
	material.Name = param.Name
	material.Level = param.Level

	material, err := m.materialRepo.Create(material)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", m.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(material)
}

func (m *material) Update(param params.MaterialEditParam) appctx.Response {
	var material entities.Material
	material, err := m.materialRepo.Get(param.ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", m.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	material.Name = param.Name
	material.Level = param.Level

	material, err = m.materialRepo.Update(material)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", m.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(material)
}

func (m *material) List(param params.MaterialFilterParam) appctx.Response {
	materials, count, err := m.materialRepo.List(param)
	if err != nil {
		logrus.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(materials).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (m *material) Detail(ID int) appctx.Response {
	material, err := m.materialRepo.Get(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Detail] %s", m.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusNotFound)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(material)
}

func (m *material) Delete(ID int) appctx.Response {
	_, err := m.materialRepo.Delete(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete] %s", m.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusNotFound)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithMessage("Material deleted successfully")
}
