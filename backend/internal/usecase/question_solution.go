package usecase

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"

	e "gitlab.com/project-quiz/utils/error"
	"gitlab.com/project-quiz/utils/minio"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type questionSolution struct {
	questionSolutionRepo repository.QuestionSolutionRepository
	name                 string
	minio                minio.MinioStorageContract
}

type QuestionSolutionUsecase interface {
	// Create new question tag
	Create(param params.QuestionSolutionCreate) appctx.Response
	CreateWithFile(param params.QuestionSolutionWithFileUploadCreate) appctx.Response
	// Edit a question tag
	Update(param params.QuestionSolutionUpdate) appctx.Response
	// // Get question tag list
	// List(param params.QuestionTagFilter) appctx.Response
	// Get detail question tag
	Detail(ID int) appctx.Response
	// Delete question tag
	Delete(ID int) appctx.Response
	// // Assign Role to user
	// Assign(userID int, roleName string) appctx.Response
	// // Revoke Role from user
	// Revoke(userID int, roleName string) appctx.Response
}

func NewQuestionSolutionUsecase(db *gorm.DB, minio minio.MinioStorageContract) QuestionSolutionUsecase {
	return &questionSolution{
		questionSolutionRepo: repository.NewQuestionSolutionRepository(db),
		name:                 "Question Solution Usecase",
		minio:                minio,
	}
}

func (q *questionSolution) Create(param params.QuestionSolutionCreate) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", q.name))

	var solution entities.QuestionSolution
	solution, err := q.questionSolutionRepo.Get(param.QuestionID)
	if err == nil {
		log.Error(fmt.Sprintf("[%s][Detail] %s", q.name, "solution existed"))
		return *appctx.NewResponse().WithErrors("solution existed").WithCode(http.StatusBadRequest)
	}

	// solution = entities.QuestionSolution{
	// 	SolutionImgUrl: param.SolutionImgUrl,
	// 	QuestionID:     param.QuestionID,
	// 	SolutionText:   param.SolutionText,
	// 	SolutionType:   param.SolutionType,
	// 	PdfFileUrl:     param.PdfFileUrl,
	// }
	copier.Copy(&solution, &param)

	data, err := q.questionSolutionRepo.Create(solution)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		if _, ok := err.(*e.ValueValidationError); ok {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(data)
}

func (q *questionSolution) CreateWithFile(param params.QuestionSolutionWithFileUploadCreate) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create With File] is executed", q.name))

	solution, err := q.questionSolutionRepo.Get(param.QuestionID)
	if err == nil {
		log.Error(fmt.Sprintf("[%s][Detail] %s", q.name, "solution existed"))
		return *appctx.NewResponse().WithErrors("solution existed").WithCode(http.StatusBadRequest)
	}

	solution = entities.QuestionSolution{
		SolutionImgUrl: "",
		QuestionID:     param.QuestionID,
		SolutionText:   param.SolutionText,
		SolutionType:   param.SolutionType,
		PdfFileUrl:     "",
	}

	if solution.SolutionType == "img" {
		path := make(chan string)
		e := make(chan error)
		go q.minio.UploadMultipart(path, e, param.SolutionImg, fmt.Sprintf("/question/%d/solution", param.QuestionID))

		if err := <-e; err != nil {
			log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
		}

		solution.SolutionImgUrl = <-path

	}

	if solution.SolutionType == "pdf" {
		path := make(chan string)
		e := make(chan error)
		go q.minio.UploadMultipart(path, e, param.SolutionImg, fmt.Sprintf("/question/%d/solution", param.QuestionID))

		if err := <-e; err != nil {
			log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
		}

		solution.PdfFileUrl = <-path
	}

	solution, err = q.questionSolutionRepo.Create(solution)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		if _, ok := err.(*e.ValueValidationError); ok {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(solution)
}

func (q *questionSolution) Update(param params.QuestionSolutionUpdate) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Update] is executed", q.name))

	// get solution
	var solution entities.QuestionSolution
	solution, err := q.questionSolutionRepo.GetByID(param.ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	logrus.Debug(param)
	copier.Copy(&solution, &param)

	solution, err = q.questionSolutionRepo.Update(solution)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(solution)
}

// func (q *questionSolution) List(param params.QuestionTagFilter) appctx.Response {
// 	log.Info(fmt.Sprintf("[%s][List] is executed", q.name))

// 	// get role list
// 	var tags []entities.QuestionTag
// 	tags, count, err := q.questionTagRepo.List(param)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
// 		return *appctx.NewResponse().WithErrors(err.Error())
// 	}

// 	return *appctx.NewResponse().WithData(tags).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
// }

func (q *questionSolution) Detail(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Detail] is executed", q.name))

	// get question solution
	var solution entities.QuestionSolution
	solution, err := q.questionSolutionRepo.GetByID(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	if solution.SolutionType == "img" && !strings.HasPrefix(solution.SolutionImgUrl, "http") {
		imgUrl, err := q.minio.GetTemporaryPublicUrl(solution.SolutionImgUrl)
		if err != nil {
			log.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
			return *appctx.NewResponse().WithErrorObj(err)
		}
		solution.SolutionImgUrl = imgUrl.String()
	}

	if solution.SolutionType == "pdf" && !strings.HasPrefix(solution.PdfFileUrl, "http") {
		pdfUrl, err := q.minio.GetTemporaryPublicUrl(solution.PdfFileUrl)
		if err != nil {
			log.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
			return *appctx.NewResponse().WithErrorObj(err)
		}
		solution.PdfFileUrl = pdfUrl.String()
	}

	return *appctx.NewResponse().WithData(solution)
}

func (q *questionSolution) Delete(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", q.name))

	// delete question tag
	err := q.questionSolutionRepo.Delete(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("question tag deleted sucessfully")
}
