package repository

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gorm.io/gorm"
)

type userQuestionMark struct {
	db   *gorm.DB
	name string
}

type UserQuestionMarkRepository interface {
	Get(userID, questionID int) (entities.UserQuestionMark, error)
	Create(mark entities.UserQuestionMark) (entities.UserQuestionMark, error)
	Delete(userID, questionID int) (entities.UserQuestionMark, error)
	GetMany(userID int, questionID []int) ([]entities.UserQuestionMark, error)
}

func NewUserQuestionMarkRepository(db *gorm.DB) UserQuestionMarkRepository {
	return &userQuestionMark{
		db:   db,
		name: "UserQuestionMarkRepository",
	}
}

func (u *userQuestionMark) Get(userID, questionID int) (entities.UserQuestionMark, error) {
	var mark entities.UserQuestionMark

	if err := u.db.Debug().Where("user_id = ? AND question_id = ?", userID, questionID).First(&mark).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", u.name, err.Error()))
		return mark, err
	}

	return mark, nil
}

func (u *userQuestionMark) Create(mark entities.UserQuestionMark) (entities.UserQuestionMark, error) {
	m, err := u.Get(mark.UserID, mark.QuestionID)
	found := false
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
			return mark, err
		}
	} else {
		found = true
	}

	if found {
		return m, err
	}

	if err := u.db.Debug().Create(&mark).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return mark, err
	}

	return mark, nil
}

func (u *userQuestionMark) Delete(userID, questionID int) (entities.UserQuestionMark, error) {
	var mark entities.UserQuestionMark

	if err := u.db.Debug().Where("user_id = ? AND question_id = ?", userID, questionID).Delete(&mark).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete] %s", u.name, err.Error()))
		return mark, err
	}

	return mark, nil
}

func (u *userQuestionMark) GetMany(userID int, questionID []int) ([]entities.UserQuestionMark, error) {
	var marks []entities.UserQuestionMark

	if err := u.db.Debug().Where("user_id = ? AND question_id in ?", userID, questionID).Find(&marks).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get Many] %s", u.name, err.Error()))
		return marks, err
	}

	return marks, nil
}
