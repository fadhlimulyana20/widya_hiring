package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/usecase"
	"gitlab.com/project-quiz/utils/boolpointer"
	"gitlab.com/project-quiz/utils/json"
	"gitlab.com/project-quiz/utils/minio"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type question struct {
	handler         Handler
	questionUsecase usecase.QuestionUsecase
	attemptUsecase  usecase.UserQuestionAttemptUsecase
	name            string
}

type QuestionHandler interface {
	// Create Question
	Create(w http.ResponseWriter, r *http.Request)
	// Update QUestion
	Update(w http.ResponseWriter, r *http.Request)
	//Admin Get List
	AdminGetList(w http.ResponseWriter, r *http.Request)
	//Admin Get Detail
	AdminGetDetail(w http.ResponseWriter, r *http.Request)
	// Admin add Option
	AdminAddOption(w http.ResponseWriter, r *http.Request)
	// Admin update Option
	AdminUpdateOption(w http.ResponseWriter, r *http.Request)
	//Admin delete option
	AdminDeleteOption(w http.ResponseWriter, r *http.Request)
	// Get list of questions
	GetList(w http.ResponseWriter, r *http.Request)
	// Get detail of question
	GetDetail(w http.ResponseWriter, r *http.Request)
	// Add and Remove Mark
	AddRemoveMark(w http.ResponseWriter, r *http.Request)
	// Get Solution
	GetSolution(w http.ResponseWriter, r *http.Request)
	// Add tags
	AddTags(w http.ResponseWriter, r *http.Request)
	// Remove tag
	RemoveTag(w http.ResponseWriter, r *http.Request)
	// Contributor Create Question
	CreateByContributor(w http.ResponseWriter, r *http.Request)
	// Get list question created by contributor
	GetListByContributor(w http.ResponseWriter, r *http.Request)
	// Contributor Update Question
	UpdateByContributor(w http.ResponseWriter, r *http.Request)
	//Contributor Get Detail
	GetDetailByContributor(w http.ResponseWriter, r *http.Request)
	// Upload Image Placement
	UploadImagePlacement(w http.ResponseWriter, r *http.Request)
}

func NewQuestionHandler(db *gorm.DB, minio minio.MinioStorageContract) QuestionHandler {
	return &question{
		questionUsecase: usecase.NewQuestionUsecase(db, minio),
		attemptUsecase:  usecase.NewUserQuestionAttemptUsecase(db),
		name:            "Uestion Handler",
	}
}

func (q *question) AdminGetList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionFilterParam
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	questionListResp := q.questionUsecase.ListWithMaterial(param)
	q.handler.Response(w, questionListResp, startTime, time.Now())
}

func (q *question) AdminGetDetail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := q.questionUsecase.Detail(idx)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) AdminAddOption(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionOptionAdd
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.AddOption(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) AdminUpdateOption(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionOptionUpdate
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)
	param.ID = idx

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.UpdateOption(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) AdminDeleteOption(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := q.questionUsecase.DeleteOption(idx)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) GetList(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionFilterParam
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	logrus.Debug(userID)

	param.Show = "active"
	param.IncludePackOnly = boolpointer.BoolPointer(false)

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	questionListResp := q.questionUsecase.List(param)
	questionList, ok := questionListResp.Data.([]entities.Question)
	if !ok {
		ctx = ctx.WithErrors("error parsing question list")
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	var questionListIDs []int
	for _, v := range questionList {
		questionListIDs = append(questionListIDs, v.ID)
	}

	attemptParam := params.AttemptGetLatestAnswersParam{
		QuestionIDs: questionListIDs,
		UserID:      userID,
	}

	attempt := q.attemptUsecase.GetLatestAnswers(attemptParam)
	attemptList, ok := attempt.Data.([]entities.UserQuestionAttempt)
	if !ok {
		ctx = ctx.WithErrors("error parsing attempt list")
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	mark := q.questionUsecase.GetMarks(userID, questionListIDs)
	markList, ok := mark.Data.([]entities.UserQuestionMark)
	if !ok {
		ctx = ctx.WithErrors("error parsing attempt list")
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	var response []params.QuestionListResponse
	for i := 0; i < len(questionList); i++ {
		var temp params.QuestionListResponse
		temp.ID = questionList[i].ID
		temp.Body = questionList[i].Body
		temp.MaterialID = questionList[i].MaterialID
		temp.Code = questionList[i].Code
		for j := 0; j < len(attemptList); j++ {
			if attemptList[j].QuestionID == temp.ID {
				temp.IsSubmitted = attemptList[j].IsSubmitted
				temp.IsAnswered = attemptList[j].QuestionID != 0
				temp.IsAnswerTrue = attemptList[j].AttemptValue
				temp.AnswerID = *attemptList[j].QuestionOptionID
				attemptList = append(attemptList[:j], attemptList[j+1:]...)
			}
		}
		for j := 0; j < len(markList); j++ {
			if markList[j].QuestionID == temp.ID {
				temp.IsMarked = true
				markList = append(markList[:j], markList[j+1:]...)
			}
		}
		response = append(response, temp)
	}

	resp := appctx.NewResponse().WithData(response).WithMetaObj(*questionListResp.Meta)

	q.handler.Response(w, *resp, startTime, time.Now())
}

func (q *question) GetDetail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	var param params.AttemptClearAnswerQuestionParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID
	param.QuestionID = idx

	resp := q.questionUsecase.Detail(idx)
	respAttemp := q.attemptUsecase.GetLatestAnswer(param)
	question := resp.Data.(entities.Question)
	// fmt.Println(question)
	var response params.QuestionDetailResponse
	response.ID = question.ID
	response.Body = question.Body
	response.MaterialID = question.MaterialID
	response.IsImage = question.IsImage
	response.ImgPath = question.ImgPath
	response.QuestionOptions = question.QuestionOptions
	response.QuestionTags = question.QuestionTags
	response.Code = question.Code
	response.ImgPlacementUrl = question.ImgPlacementUrl
	TrueAnswerID := 0
	for i := 0; i < len(response.QuestionOptions); i++ {
		if *response.QuestionOptions[i].OptionValue {
			TrueAnswerID = response.QuestionOptions[i].ID
		}
		*response.QuestionOptions[i].OptionValue = false
	}

	attempt, ok := respAttemp.Data.(entities.UserQuestionAttempt)
	if ok {
		response.IsAnswerTrue = attempt.AttemptValue
		response.IsAnswered = attempt.IsSubmitted
		response.AnswerID = *attempt.QuestionOptionID
		response.TrueAnswerID = TrueAnswerID
	}

	mark := q.questionUsecase.GetMark(userID, param.QuestionID)
	_, ok = mark.Data.(entities.UserQuestionMark)
	if ok {
		response.IsMarked = true
	}

	responseFinal := *appctx.NewResponse().WithData(response)

	q.handler.Response(w, responseFinal, startTime, time.Now())
}

func (q *question) AddRemoveMark(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	var param params.QuestionMarkAddRemoveParam
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.UserID = userID

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if param.Action == "add" {
		resp := q.questionUsecase.AddMark(param.UserID, param.QuestionID)
		q.handler.Response(w, resp, startTime, time.Now())
		return
	} else if param.Action == "remove" {
		resp := q.questionUsecase.DeleteMark(param.UserID, param.QuestionID)
		q.handler.Response(w, resp, startTime, time.Now())
		return
	}
}

func (q *question) GetSolution(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := q.questionUsecase.GetSolution(idx)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) Create(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionCreate
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.ContributorID = userID

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", q.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.Create(param, true)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) Update(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionUpdate
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)
	param.ID = idx

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", q.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.Update(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) AddTags(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionAddTags
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", q.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.AddTags(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) RemoveTag(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionRemoveTag

	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", q.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.RemoveTag(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) CreateByContributor(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionCreate
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.ContributorID = userID

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", q.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := q.questionUsecase.Create(param, false)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) GetListByContributor(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.QuestionFilterParam
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.ContributorID = userID

	logrus.Debug(userID)

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	questionListResp := q.questionUsecase.ListWithMaterial(param)
	q.handler.Response(w, questionListResp, startTime, time.Now())
}

func (q *question) UpdateByContributor(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var param params.QuestionUpdate
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)
	param.ID = idx
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.ContributorID = userID

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", q.name))
		ctx = ctx.WithErrors(err.Error())
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", q.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		q.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	param.IsActive = nil
	resp := q.questionUsecase.Update(param)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) GetDetailByContributor(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	idx, _ := strconv.Atoi(id)

	resp := q.questionUsecase.Detail(idx)
	q.handler.Response(w, resp, startTime, time.Now())
}

func (q *question) UploadImagePlacement(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	// ctx := appctx.NewResponse()

	if err := r.ParseMultipartForm(1024); err != nil {
		d := appctx.NewResponse().WithErrors(err.Error())
		q.handler.Response(w, *d, startTime, time.Now())
		return
	}

	_, fileHeader, err := r.FormFile("file")
	if err != nil {
		d := appctx.NewResponse().WithErrors(err.Error())
		q.handler.Response(w, *d, startTime, time.Now())
		return
	}

	questionID := r.FormValue("question_id")
	questionIDNumber, err := strconv.Atoi(questionID)
	if err != nil {
		d := appctx.NewResponse().WithErrors(err.Error())
		q.handler.Response(w, *d, startTime, time.Now())
		return
	}
	resp := q.questionUsecase.UploadImagePlacement(questionIDNumber, fileHeader)
	q.handler.Response(w, resp, startTime, time.Now())
}
