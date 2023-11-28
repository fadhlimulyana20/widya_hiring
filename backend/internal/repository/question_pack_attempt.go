package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"
	"gorm.io/gorm"
)

type questionPackAttemptRepo struct {
	db   *gorm.DB
	name string
}

type QuestionPackAttemptRepository interface {
	// Create new question pack attempt
	Create(pack entities.QuestionPackAttempt) (entities.QuestionPackAttempt, error)
	// Update question pack
	Update(pack entities.QuestionPackAttempt) (entities.QuestionPackAttempt, error)
	// Get question packd detail
	Get(ID int) (entities.QuestionPackAttempt, error)
	// Get Lis Question packet attemp
	GetList(param params.QuestionPackAttemptFilterParam) ([]entities.QuestionPackAttempt, int, error)
}

func NewQuestionPackAttemptRepository(db *gorm.DB) QuestionPackAttemptRepository {
	return &questionPackAttemptRepo{
		db:   db,
		name: "Question Pack Attempt REpos",
	}
}

func (q *questionPackAttemptRepo) Create(pack entities.QuestionPackAttempt) (entities.QuestionPackAttempt, error) {
	if err := q.db.Create(&pack).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return pack, err
	}

	return pack, nil
}

func (q *questionPackAttemptRepo) Update(pack entities.QuestionPackAttempt) (entities.QuestionPackAttempt, error) {
	if err := q.db.Model(&pack).Updates(&pack).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return pack, err
	}

	return pack, nil
}

func (q *questionPackAttemptRepo) Get(ID int) (entities.QuestionPackAttempt, error) {
	var pack entities.QuestionPackAttempt

	if err := q.db.Debug().First(&pack, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GET] %s", q.name, err.Error()))
		return pack, err
	}

	return pack, nil
}

func (q *questionPackAttemptRepo) GetList(param params.QuestionPackAttemptFilterParam) ([]entities.QuestionPackAttempt, int, error) {
	var packs []entities.QuestionPackAttempt
	var count int64
	db := q.db
	useFilterCount := false

	db.Model(&packs).Count(&count)

	if param.IsFinish != nil {
		db = db.Where("is_finish = ?", *param.IsFinish)
		useFilterCount = true
	}

	if param.UserID != 0 {
		db = db.Where("user_id = ?", param.UserID)
		useFilterCount = true
	}

	if param.QuestionPackID != 0 {
		db = db.Where("question_pack_id = ?", param.QuestionPackID)
		useFilterCount = true
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&packs).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return packs, 0, err
	}

	if useFilterCount {
		db.Count(&count)
	}

	return packs, int(count), nil
}
