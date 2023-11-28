package usecase

import (
	"errors"
	"net/http"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gorm.io/gorm"
)

type userQuestionAttempt struct {
	attemptRepo   repository.UserQuestionAttemptRepository
	optionRepo    repository.QuestionOptionRepository
	userPointRepo repository.UserPointRepository
	name          string
}

type UserQuestionAttemptUsecase interface {
	AnswerQuestion(param params.AttemptAnswerQuestionParam) appctx.Response
	ClearAnswer(param params.AttemptClearAnswerQuestionParam) appctx.Response
	SubmitAnswer(param params.AttemptSubmitAnswerQuestionParam) appctx.Response
	GetLatestAnswer(param params.AttemptClearAnswerQuestionParam) appctx.Response
	MarkAttempt(param params.AttemptSubmitAnswerQuestionParam) appctx.Response
	GetLatestAnswers(param params.AttemptGetLatestAnswersParam) appctx.Response
}

func NewUserQuestionAttemptUsecase(db *gorm.DB) UserQuestionAttemptUsecase {
	return &userQuestionAttempt{
		attemptRepo:   repository.NewUserQuestionAttemptRepository(db),
		optionRepo:    repository.NewQuestionOptionRepository(db),
		userPointRepo: repository.NewUserPointRepository(db),
		name:          "User Question Attempt Usecase",
	}
}

func (u *userQuestionAttempt) AnswerQuestion(param params.AttemptAnswerQuestionParam) appctx.Response {
	attempt, err := u.attemptRepo.GetLatest(param.QuestionID, param.UserID)
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrorObj(err)
		}
		found = false
	}

	if found {
		attempt.QuestionOptionID = &param.OptionID
		attempt, err = u.attemptRepo.Update(attempt)
		if err != nil {
			return *appctx.NewResponse().WithErrorObj(err)
		}
	} else {
		attempt.QuestionID = param.QuestionID
		attempt.QuestionOptionID = &param.OptionID
		attempt.UserID = param.UserID
		attempt, err = u.attemptRepo.Create(attempt)
		if err != nil {
			return *appctx.NewResponse().WithErrorObj(err)
		}
	}

	return *appctx.NewResponse().WithData(attempt)
}

func (u *userQuestionAttempt) ClearAnswer(param params.AttemptClearAnswerQuestionParam) appctx.Response {
	attempt, err := u.attemptRepo.GetLatest(param.QuestionID, param.UserID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	attempt.QuestionOptionID = nil
	attempt.AttemptValue = false
	attempt, err = u.attemptRepo.UpdateField(attempt, []string{"question_option_id", "attempt_value"})
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}
	return *appctx.NewResponse().WithData(attempt)
}

func (u *userQuestionAttempt) GetLatestAnswer(param params.AttemptClearAnswerQuestionParam) appctx.Response {
	attempt, err := u.attemptRepo.GetLatestSubmitted(param.QuestionID, param.UserID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrorObj(err)
		} else {
			attempt, err = u.attemptRepo.GetLatest(param.QuestionID, param.UserID)
			if err != nil {
				return *appctx.NewResponse().WithErrorObj(err)
			}
		}
	}

	return *appctx.NewResponse().WithData(attempt)
}

func (u *userQuestionAttempt) GetLatestAnswers(param params.AttemptGetLatestAnswersParam) appctx.Response {
	attempt, err := u.attemptRepo.GetLatestSubmittedAnswers(param.QuestionIDs, param.UserID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(attempt)
}

func (u *userQuestionAttempt) MarkAttempt(param params.AttemptSubmitAnswerQuestionParam) appctx.Response {
	attempt, err := u.attemptRepo.GetLatest(param.QuestionID, param.UserID)
	found := true
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrorObj(err)
		}
		found = false
	}

	if !found {
		attempt.QuestionID = param.QuestionID
		attempt.UserID = param.UserID
		attempt, err = u.attemptRepo.Create(attempt)
		if err != nil {
			return *appctx.NewResponse().WithErrorObj(err)
		}
	}

	attempt.IsMarked = true
	attempt, err = u.attemptRepo.Update(attempt)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("Soal berhasil ditandai")
}

func (u *userQuestionAttempt) SubmitAnswer(param params.AttemptSubmitAnswerQuestionParam) appctx.Response {
	attempt, err := u.attemptRepo.GetLatest(param.QuestionID, param.UserID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	if attempt.QuestionOptionID == nil {
		return *appctx.NewResponse().WithErrors("Jawaban kosong").WithCode(http.StatusBadRequest)
	}

	option, err := u.optionRepo.Get(*attempt.QuestionOptionID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	if *option.OptionValue {
		go u.userPointRepo.UpdateOrCreate(param.UserID, 3)
	}

	attempt.AttemptValue = *option.OptionValue
	attempt.IsSubmitted = true
	attempt.IsMarked = false
	attempt, err = u.attemptRepo.Update(attempt)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	trueOption, err := u.optionRepo.GetTrueOption(param.QuestionID)
	if err != nil {
		return *appctx.NewResponse().WithErrorObj(err)
	}

	response := &params.AttemptSubmitAnswerResponse{
		AttemptValue:     *option.OptionValue,
		TrueAnswerID:     trueOption.ID,
		AnswerID:         option.ID,
		TrueAnswerStreak: 0, // Gonna develop it later LOL
	}

	return *appctx.NewResponse().WithData(response)
}
