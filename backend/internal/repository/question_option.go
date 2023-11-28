package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gorm.io/gorm"
)

type questionOption struct {
	db   *gorm.DB
	name string
}

type QuestionOptionRepository interface {
	// Get option detail
	Get(ID int) (entities.QuestionOption, error)
	// Add Question Option
	Create(param entities.QuestionOption) (entities.QuestionOption, error)
	// Update Question Option
	Update(param entities.QuestionOption) (entities.QuestionOption, error)
	// Delete Question Option
	Delete(ID int) (bool, error)
	GetTrueOption(questionID int) (entities.QuestionOption, error)
}

func NewQuestionOptionRepository(db *gorm.DB) QuestionOptionRepository {
	return &questionOption{
		db:   db,
		name: "Question Option Repository",
	}
}

func (q *questionOption) Get(ID int) (entities.QuestionOption, error) {
	var questionOption entities.QuestionOption

	if err := q.db.Debug().First(&questionOption, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", q.name, err.Error()))
		return questionOption, err
	}

	return questionOption, nil
}

func (q *questionOption) GetTrueOption(questionID int) (entities.QuestionOption, error) {
	var questionOption entities.QuestionOption

	if err := q.db.Debug().Where("question_id = ? AND option_value = ?", questionID, true).First(&questionOption).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetTrueOption] %s", q.name, err.Error()))
		return questionOption, nil
	}

	return questionOption, nil
}

func (q *questionOption) Create(param entities.QuestionOption) (entities.QuestionOption, error) {
	if err := q.db.Debug().Create(&param).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return param, err
	}

	return param, nil
}

func (q *questionOption) Update(param entities.QuestionOption) (entities.QuestionOption, error) {
	if err := q.db.Debug().Model(&param).Updates(&param).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return param, err
	}

	return param, nil
}

func (q *questionOption) Delete(ID int) (bool, error) {
	var option entities.QuestionOption
	if err := q.db.Debug().Delete(&option, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete] %s", q.name, err.Error()))
		return false, err
	}

	return true, nil
}
