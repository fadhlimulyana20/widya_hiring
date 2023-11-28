package usecase

import (
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/repository"
	"gitlab.com/project-quiz/utils/minio"
	"gorm.io/gorm"
)

type analytic struct {
	name              string
	attemptRepository repository.UserQuestionAttemptRepository
	userRepo          repository.UserRepository
	materialRepo      repository.MaterialRepository
	questionRepo      repository.QuestionRepository
}

type AnalyticUsecase interface {
	GetCreatorAnalytic() appctx.Response
	GetTotalUserAttempt(userID int) appctx.Response
}

func NewAnaliticUsecase(db *gorm.DB, m minio.MinioStorageContract) AnalyticUsecase {
	return &analytic{
		name:              "Analytic Usecase",
		attemptRepository: repository.NewUserQuestionAttemptRepository(db),
		userRepo:          repository.NewUserRepository(db),
		materialRepo:      repository.NewMaterialRepository(db),
		questionRepo:      repository.NewQuestionRepository(db, m),
	}
}

func (a *analytic) GetCreatorAnalytic() appctx.Response {
	ctx := appctx.NewResponse()

	userCount, err := a.userRepo.GetTotal()
	if err != nil {
		return *ctx.WithErrorObj(err)
	}

	materialCount, err := a.materialRepo.GetTotal()
	if err != nil {
		return *ctx.WithErrorObj(err)
	}

	questionCount, err := a.questionRepo.GetTotal()
	if err != nil {
		return *ctx.WithErrorObj(err)
	}

	data := map[string]int{
		"user":     userCount,
		"material": materialCount,
		"question": questionCount,
	}

	return *ctx.WithData(data)
}

func (a *analytic) GetTotalUserAttempt(userID int) appctx.Response {
	ctx := appctx.NewResponse()

	totalAttempt, err := a.attemptRepository.GetTotalAttempt(userID)
	if err != nil {
		return *ctx.WithErrorObj(err)
	}

	totalTrueAttempt, err := a.attemptRepository.GetTotalAttemptWithValueType(userID, true)
	if err != nil {
		return *ctx.WithErrorObj(err)
	}

	totalFalseAttempt, err := a.attemptRepository.GetTotalAttemptWithValueType(userID, false)
	if err != nil {
		return *ctx.WithErrorObj(err)
	}

	data := map[string]int{
		"total":       totalAttempt,
		"total_true":  totalTrueAttempt,
		"total_false": totalFalseAttempt,
	}

	return *ctx.WithData(data)
}
