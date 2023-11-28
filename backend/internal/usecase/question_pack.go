package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gorm.io/gorm"
)

type questionPack struct {
	questionPackRepo       repository.QuestionPackRepository
	questionPackAttempRepo repository.QuestionPackAttemptRepository
	name                   string
}

type QuestionPackUsecase interface {
	// Creeate Question pack
	Create(param params.QuestionPackCreateParam) appctx.Response
	// Update Quetion pack
	Update(param params.QuestionPackUpdateParam) appctx.Response
	// Get list of Question Pack
	List(param params.QuestionPackFilterParam) appctx.Response
	// Get detail of question pack
	Detail(ID int) appctx.Response
	// Delete question pack
	Delete(ID int) appctx.Response
	// Add Questions
	AddQuestions(param params.QuestionPackAddQuestionParam) appctx.Response
	// Delete Quetions
	DeleteQuestions(param params.QuestionPackAddQuestionParam) appctx.Response
	// Take question pack
	TakeQuestionPack(QuestionPackID, UserID int) appctx.Response
	// Finish question pack
	FinishQuestionPack(QuestionPackAttemptID, UserID int) appctx.Response
	// Get attempt list
	GetAttemptList(param params.QuestionPackAttemptFilterParam) appctx.Response
	// Get Question Pack Attempt Detail
	// GetQuestionPackAttemptDetail(QuestionPackAttemptID int) appctx.Response
}

func NewQuestionPackUsecase(db *gorm.DB) QuestionPackUsecase {
	return &questionPack{
		questionPackRepo:       repository.NewQuestionPackRepository(db),
		questionPackAttempRepo: repository.NewQuestionPackAttemptRepository(db),
		name:                   "QUestion Pack Usecase",
	}
}

func (q *questionPack) Create(param params.QuestionPackCreateParam) appctx.Response {
	var pack entities.QuestionPack
	copier.Copy(&pack, &param)
	ctx := appctx.NewResponse()

	pack, err := q.questionPackRepo.Create(pack)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithData(pack)
}

func (q *questionPack) Update(param params.QuestionPackUpdateParam) appctx.Response {
	var pack entities.QuestionPack
	ctx := appctx.NewResponse()

	pack, err := q.questionPackRepo.Get(param.ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	copier.Copy(&pack, &param)
	pack, err = q.questionPackRepo.Update(pack)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithData(pack)
}

func (q *questionPack) List(param params.QuestionPackFilterParam) appctx.Response {
	ctx := appctx.NewResponse()
	packs, count, err := q.questionPackRepo.GetList(param)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithData(packs).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (q *questionPack) Detail(ID int) appctx.Response {
	ctx := appctx.NewResponse()
	pack, err := q.questionPackRepo.Get(ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithData(pack)
}

func (q *questionPack) Delete(ID int) appctx.Response {
	ctx := appctx.NewResponse()
	if err := q.questionPackRepo.Delete(ID); err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithMessage("question pack deleted")
}

func (q *questionPack) AddQuestions(param params.QuestionPackAddQuestionParam) appctx.Response {
	ctx := appctx.NewResponse()
	_, err := q.questionPackRepo.Get(param.ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Add Question] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	err = q.questionPackRepo.AddQuestions(param.ID, param.QuestionIDs)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Add Question] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithMessage("Questions has been added")
}

func (q *questionPack) DeleteQuestions(param params.QuestionPackAddQuestionParam) appctx.Response {
	ctx := appctx.NewResponse()
	_, err := q.questionPackRepo.Get(param.ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete Question] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	err = q.questionPackRepo.DeleteQuestions(param.ID, param.QuestionIDs)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete Question] %s", q.name, err.Error()))
		return *ctx.WithErrors(err.Error())
	}

	return *ctx.WithMessage("Questions has been deleted")
}

func (q *questionPack) TakeQuestionPack(questionPackID, userID int) appctx.Response {
	ctx := appctx.NewResponse()

	_, err := q.questionPackRepo.Get(questionPackID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Take Question Pack] %s", q.name, err.Error()))
		return *ctx.WithErrorContexts(err)
	}

	questionPackAttempt := entities.QuestionPackAttempt{
		UserID:         userID,
		QuestionPackID: questionPackID,
		IsFinish:       false,
		StartedAt:      time.Now(),
		Score:          0,
	}

	questionPackAttempt, err = q.questionPackAttempRepo.Create(questionPackAttempt)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Take Question Pack] %s", q.name, err.Error()))
		return *ctx.WithErrorContexts(err)
	}

	return *ctx.WithData(questionPackAttempt)
}

func (q *questionPack) FinishQuestionPack(questionPackAttemptID, userID int) appctx.Response {
	ctx := appctx.NewResponse()

	questionPackAttempt, err := q.questionPackAttempRepo.Get(questionPackAttemptID)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Take Question Pack] %s", q.name, err.Error()))
		return *ctx.WithErrorContexts(err)
	}

	if questionPackAttempt.UserID != userID {
		return *ctx.WithErrors("Invalid user").WithCode(http.StatusBadRequest)
	}

	questionPackAttempt.IsFinish = true
	questionPackAttempt.FinishedAt = time.Now()

	questionPackAttempt, err = q.questionPackAttempRepo.Update(questionPackAttempt)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Take Question Pack] %s", q.name, err.Error()))
		return *ctx.WithErrorContexts(err)
	}

	return *ctx.WithData(questionPackAttempt)
}

func (q *questionPack) GetAttemptList(param params.QuestionPackAttemptFilterParam) appctx.Response {
	ctx := appctx.NewResponse()

	questionPackAttempts, count, err := q.questionPackAttempRepo.GetList(param)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get Attempt List] %s", q.name, err.Error()))
		return *ctx.WithErrorObj(err)
	}

	return *ctx.WithData(questionPackAttempts).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

// func (q *questionPack) GetQuestionPackAttemptDetail(QuestionPackAttemptID int) appctx.Response {
// 	ctx := appctx.NewResponse()

// 	questionPackAttempt, err := q.questionPackAttempRepo.Get(QuestionPackAttemptID)
// 	if err != nil {
// 		logrus.Error(fmt.Sprintf("[%s][Get Attempt List] %s", q.name, err.Error()))
// 		return *ctx.WithErrorObj(err)
// 	}

// 	// Get question pack
// 	questionPack, err := q.questionPackRepo.Get(questionPackAttempt.QuestionPackID)
// 	if err != nil {
// 		logrus.Error(fmt.Sprintf("[%s][Get Attempt List] %s", q.name, err.Error()))
// 		return *ctx.WithErrorObj(err)
// 	}

// 	// Get question option
// 	// Get question attempt

// 	// return
// }
