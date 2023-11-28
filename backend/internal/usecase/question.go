package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gitlab.com/project-quiz/utils/boolpointer"
	"gitlab.com/project-quiz/utils/minio"
	"gorm.io/gorm"
)

type question struct {
	questionRepo repository.QuestionRepository
	markRepo     repository.UserQuestionMarkRepository
	solutionRepo repository.QuestionSolutionRepository
	optionRepo   repository.QuestionOptionRepository
	tagRepo      repository.QuestionTagRepository
	minio        minio.MinioStorageContract
	name         string
}

type QuestionUsecase interface {
	// Create Question
	Create(param params.QuestionCreate, isAdmin bool) appctx.Response
	// Update Question
	Update(param params.QuestionUpdate) appctx.Response
	// Add Option
	AddOption(param params.QuestionOptionAdd) appctx.Response
	// Add Option
	UpdateOption(param params.QuestionOptionUpdate) appctx.Response
	// Delete Option
	DeleteOption(ID int) appctx.Response
	// Get list of material
	List(param params.QuestionFilterParam) appctx.Response
	// Get list of material
	ListWithMaterial(param params.QuestionFilterParam) appctx.Response
	// Get detail of material
	Detail(ID int) appctx.Response
	// Get mark to question
	GetMark(userID, questionID int) appctx.Response
	// Add mark to question
	AddMark(userID, questionID int) appctx.Response
	// Delete Mark
	DeleteMark(userID, questionID int) appctx.Response
	// Get Marks
	GetMarks(userID int, questionID []int) appctx.Response
	// Get Solution
	GetSolution(questionID int) appctx.Response
	// Add Tags
	AddTags(param params.QuestionAddTags) appctx.Response
	// Remove Tag
	RemoveTag(param params.QuestionRemoveTag) appctx.Response
	// Add Image Placement
	UploadImagePlacement(questionID int, file *multipart.FileHeader) appctx.Response
}

func NewQuestionUsecase(db *gorm.DB, minio minio.MinioStorageContract) QuestionUsecase {
	return &question{
		questionRepo: repository.NewQuestionRepository(db, minio),
		markRepo:     repository.NewUserQuestionMarkRepository(db),
		solutionRepo: repository.NewQuestionSolutionRepository(db),
		optionRepo:   repository.NewQuestionOptionRepository(db),
		tagRepo:      repository.NewQuestionTagRepository(db),
		minio:        minio,
		name:         "Question Usecase",
	}
}

func (q *question) List(param params.QuestionFilterParam) appctx.Response {
	materials, count, err := q.questionRepo.List(param)
	if err != nil {
		logrus.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(materials).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (q *question) ListWithMaterial(param params.QuestionFilterParam) appctx.Response {
	materials, count, err := q.questionRepo.ListJoinMaterial(param)
	if err != nil {
		logrus.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(materials).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (q *question) Detail(ID int) appctx.Response {
	material, err := q.questionRepo.Get(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusNotFound)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	if material.ImgPlacementUrl != "" {
		imgUrl, err := q.minio.GetTemporaryPublicUrl(material.ImgPlacementUrl)
		if err == nil {
			material.ImgPlacementUrl = imgUrl.String()
		}
	}

	return *appctx.NewResponse().WithData(material)
}

func (q *question) AddOption(param params.QuestionOptionAdd) appctx.Response {
	option := entities.QuestionOption{
		Body:        param.Body,
		OptionValue: &param.OptionValue,
		IsImage:     param.IsImage,
		ImgPath:     param.ImgPath,
		QuestionID:  param.QuestionID,
	}
	option, err := q.optionRepo.Create(option)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(option)
}

func (q *question) UpdateOption(param params.QuestionOptionUpdate) appctx.Response {
	option, err := q.optionRepo.Get(param.ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	option.Body = param.Body
	option.ImgPath = param.ImgPath
	option.IsImage = param.IsImage
	option.OptionValue = &param.OptionValue

	option, err = q.optionRepo.Update(option)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(option)
}

func (q *question) DeleteOption(ID int) appctx.Response {
	_, err := q.optionRepo.Delete(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][DeleteOPtion] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithMessage("Option deleted successfully")
}

func (q *question) GetMark(userID, questionID int) appctx.Response {
	mark, err := q.markRepo.Get(userID, questionID)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusNotFound)
	}

	return *appctx.NewResponse().WithData(mark)
}

func (q *question) AddMark(userID, questionID int) appctx.Response {
	mark := entities.UserQuestionMark{
		QuestionID: questionID,
		UserID:     userID,
	}
	mark, err := q.markRepo.Create(mark)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(mark)
}

func (q *question) DeleteMark(userID, questionID int) appctx.Response {
	mark, err := q.markRepo.Delete(userID, questionID)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(mark)
}

func (q *question) GetMarks(userID int, questionIDs []int) appctx.Response {
	marks, err := q.markRepo.GetMany(userID, questionIDs)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(marks)
}

func (q *question) GetSolution(questionID int) appctx.Response {
	solution, err := q.solutionRepo.Get(questionID)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(solution)
}

func (q *question) Create(param params.QuestionCreate, isAdmin bool) appctx.Response {
	// question := entities.Question{
	// 	Body:            param.Body,
	// 	MaterialID:      param.MaterialID,
	// 	IsImage:         param.IsImage,
	// 	ImgPath:         param.ImgPath,
	// 	ImgPlacementUrl: param.ImgPlacementUrl,
	// }
	var question entities.Question
	copier.Copy(&question, param)

	if !isAdmin {
		question.IsActive = boolpointer.BoolPointer(false)
	}

	question, err := q.questionRepo.Create(question)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	// for i := 0; i < len(param.QuestionOptions); i++ {
	// 	questionOption := entities.QuestionOption{
	// 		Body:        param.QuestionOptions[i].Body,
	// 		OptionValue: &param.QuestionOptions[i].OptionValue,
	// 		IsImage:     param.QuestionOptions[i].IsImage,
	// 		ImgPath:     param.QuestionOptions[i].ImgPath,
	// 		QuestionID:  question.ID,
	// 	}

	// 	questionOption, err := q.optionRepo.Create(questionOption)
	// 	if err != nil {
	// 		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	// 	}

	// 	question.QuestionOptions = append(question.QuestionOptions, questionOption)
	// }

	return *appctx.NewResponse().WithData(question)
}

func (q *question) Update(param params.QuestionUpdate) appctx.Response {
	// question := entities.Question{
	// 	ID:              param.ID,
	// 	Body:            param.Body,
	// 	MaterialID:      param.MaterialID,
	// 	IsImage:         param.IsImage,
	// 	ImgPath:         param.ImgPath,
	// 	IsActive:        param.IsActive,
	// 	ImgPlacementUrl: param.ImgPlacementUrl,
	// 	ContributorID:   param.ContributorID,
	// }

	var question entities.Question
	copier.Copy(&question, param)
	question.QuestionOptions = []entities.QuestionOption{}

	question, err := q.questionRepo.Update(question)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	logrus.Debug(question)

	for i := 0; i < len(param.QuestionOptions); i++ {
		var questionOption entities.QuestionOption
		if param.QuestionOptions[i].ID == 0 {
			questionOption = entities.QuestionOption{
				Body:        param.QuestionOptions[i].Body,
				OptionValue: &param.QuestionOptions[i].OptionValue,
				IsImage:     param.QuestionOptions[i].IsImage,
				ImgPath:     param.QuestionOptions[i].ImgPath,
				QuestionID:  question.ID,
			}

			questionOption, err = q.optionRepo.Create(questionOption)
			if err != nil {
				return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
			}
		} else {
			questionOption = entities.QuestionOption{
				ID:          param.QuestionOptions[i].ID,
				Body:        param.QuestionOptions[i].Body,
				OptionValue: &param.QuestionOptions[i].OptionValue,
				IsImage:     param.QuestionOptions[i].IsImage,
				ImgPath:     param.QuestionOptions[i].ImgPath,
				QuestionID:  question.ID,
			}

			questionOption, err = q.optionRepo.Update(questionOption)
			if err != nil {
				return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
			}
		}
		question.QuestionOptions = append(question.QuestionOptions, questionOption)
	}

	return *appctx.NewResponse().WithData(question)
}

func (q *question) AddTags(param params.QuestionAddTags) appctx.Response {
	question, err := q.questionRepo.Get(param.QuestionID)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	tags, count, err := q.tagRepo.ListIn(param.TagIDs)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	logrus.Debug(count)

	if count > 0 {
		question, err := q.questionRepo.AddTag(question, tags)
		if err != nil {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		}
		return *appctx.NewResponse().WithData(question)
	}

	return *appctx.NewResponse().WithData(question)
}

func (q *question) RemoveTag(param params.QuestionRemoveTag) appctx.Response {
	question, err := q.questionRepo.Get(param.QuestionID)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	tag, err := q.tagRepo.Get(param.TagID)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	question, err = q.questionRepo.RemoveTag(question, tag)
	if err != nil {
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	return *appctx.NewResponse().WithData(question)
}

func (q *question) UploadImagePlacement(questionID int, file *multipart.FileHeader) appctx.Response {
	// Find Question
	quest, err := q.questionRepo.Get(questionID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	path := make(chan string)
	e := make(chan error)
	go q.minio.UploadMultipart(path, e, file, fmt.Sprintf("/questions/%d/img_placement", questionID))

	if err := <-e; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
	}

	quest.ImgPlacementUrl = <-path
	quest, err = q.questionRepo.Update(quest)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(quest)
}
