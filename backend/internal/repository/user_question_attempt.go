package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gorm.io/gorm"
)

type userQuestionAttempt struct {
	db   *gorm.DB
	name string
}

type UserQuestionAttemptRepository interface {
	// Create
	Create(attempt entities.UserQuestionAttempt) (entities.UserQuestionAttempt, error)
	// Update
	Update(attempt entities.UserQuestionAttempt) (entities.UserQuestionAttempt, error)
	// Update Field
	UpdateField(attempt entities.UserQuestionAttempt, fields []string) (entities.UserQuestionAttempt, error)
	// Get Latest
	GetLatest(questionID int, userID int) (entities.UserQuestionAttempt, error)
	// Get Latets answered
	GetLatestSubmitted(questionID int, userID int) (entities.UserQuestionAttempt, error)
	GetLatestSubmittedAnswers(questionIDs []int, userID int) ([]entities.UserQuestionAttempt, error)
	GetTotalAttempt(userID int) (int, error)
	GetTotalAttemptWithValueType(userID int, valueType bool) (int, error)
}

func NewUserQuestionAttemptRepository(db *gorm.DB) UserQuestionAttemptRepository {
	return &userQuestionAttempt{
		db:   db,
		name: "User Question Attempt Repository",
	}
}

func (uqa *userQuestionAttempt) Create(attempt entities.UserQuestionAttempt) (entities.UserQuestionAttempt, error) {
	if err := uqa.db.Debug().Create(&attempt).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", uqa.name, err.Error()))
		return attempt, err
	}

	return attempt, nil
}

func (uqa *userQuestionAttempt) Update(attempt entities.UserQuestionAttempt) (entities.UserQuestionAttempt, error) {
	if err := uqa.db.Debug().Model(&attempt).Updates(&attempt).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", uqa.name, err.Error()))
		return attempt, err
	}

	return attempt, nil
}

func (uqa *userQuestionAttempt) UpdateField(attempt entities.UserQuestionAttempt, fields []string) (entities.UserQuestionAttempt, error) {
	if err := uqa.db.Debug().Model(&attempt).Select(fields).Updates(&attempt).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", uqa.name, err.Error()))
		return attempt, err
	}

	return attempt, nil
}

func (uqa *userQuestionAttempt) GetLatest(questionID int, userID int) (entities.UserQuestionAttempt, error) {
	var attempt entities.UserQuestionAttempt
	if err := uqa.db.Where("question_id = ? AND user_id = ? AND is_submitted = ?", questionID, userID, false).Order("created_at desc").First(&attempt).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetLatest] %s", uqa.name, err.Error()))
		return attempt, err
	}

	return attempt, nil
}

func (uqa *userQuestionAttempt) GetLatestSubmitted(questionID int, userID int) (entities.UserQuestionAttempt, error) {
	var attempt entities.UserQuestionAttempt
	if err := uqa.db.Where("question_id = ? AND user_id = ? AND is_submitted = ?", questionID, userID, true).Order("created_at desc").First(&attempt).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetLatest] %s", uqa.name, err.Error()))
		return attempt, err
	}

	return attempt, nil
}
func (uqa *userQuestionAttempt) GetLatestSubmittedAnswers(questionIDs []int, userID int) ([]entities.UserQuestionAttempt, error) {
	var attempts []entities.UserQuestionAttempt
	sqlStatement := `select id, question_id, question_option_id, user_id, attempt_value, is_submitted
	from (
		select id, question_id, question_option_id, user_id, attempt_value, is_submitted, row_number() over(partition by question_id, user_id order by is_submitted desc, created_at desc) as rn
		from user_question_attempts uqa
	) as R
	where R.rn = 1 and R.user_id = ? and R.question_id in ?`

	if err := uqa.db.Debug().Raw(sqlStatement, userID, questionIDs).Scan(&attempts).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetLatestSubmittedAnswers] %s", uqa.name, err.Error()))
		return attempts, err
	}

	return attempts, nil
}

func (uqa *userQuestionAttempt) GetTotalAttempt(userID int) (int, error) {
	var count int
	sqlStatement := `select count(id) from user_question_attempts where user_id = ?`

	if err := uqa.db.Debug().Raw(sqlStatement, userID).Scan(&count).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get Total Attempt] %s", uqa.name, err.Error()))
		return 0, err
	}

	return count, nil
}

func (uqa *userQuestionAttempt) GetTotalAttemptWithValueType(userID int, valueType bool) (int, error) {
	var count int
	sqlStatement := `select count(id) from user_question_attempts where user_id = ? and attempt_value = ?`

	if err := uqa.db.Debug().Raw(sqlStatement, userID, valueType).Scan(&count).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get Total Attempt With Value Type] %s", uqa.name, err.Error()))
		return 0, err
	}

	return count, nil
}
