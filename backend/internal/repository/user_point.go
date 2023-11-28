package repository

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"
	"gorm.io/gorm"
)

type userPoint struct {
	db   *gorm.DB
	name string
}

type UserPointRepository interface {
	// Get User Point
	GetByUser(UserID int) (entities.UserPoint, error)
	// Get List of User Point
	GetListOfUserPoint(params.UserPointFilterParam) ([]entities.UserPoint, int, error)
	// Update User Point
	Update(entities.UserPoint) (entities.UserPoint, error)
	// Create User Point
	Create(entities.UserPoint) (entities.UserPoint, error)
	// Update or Create User Point
	UpdateOrCreate(UserID, AddedPoint int) (entities.UserPoint, error)
}

func NewUserPointRepository(db *gorm.DB) UserPointRepository {
	return &userPoint{
		db:   db,
		name: "User Point Repository",
	}
}

func (u *userPoint) GetByUser(UserID int) (entities.UserPoint, error) {
	var up entities.UserPoint

	if err := u.db.Debug().Preload("User").Where("user_id = ?", UserID).First(&up).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetByUser] %s", u.name, err.Error()))
		return up, err
	}

	return up, nil
}

func (u *userPoint) Update(up entities.UserPoint) (entities.UserPoint, error) {
	if err := u.db.Debug().Model(&up).Updates(&up).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", u.name, err.Error()))
		return up, err
	}

	return up, nil
}

func (u *userPoint) Create(up entities.UserPoint) (entities.UserPoint, error) {
	if err := u.db.Debug().Create(&up).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return up, err
	}

	return up, nil
}

func (u *userPoint) UpdateOrCreate(UserID, AddedPoint int) (entities.UserPoint, error) {
	var up entities.UserPoint

	// Check if user point exist
	up, err := u.GetByUser(UserID)

	// Update user point
	if err == nil {
		up.Point += AddedPoint
		up, er := u.Update(up)
		if er != nil {
			logrus.Error(fmt.Sprintf("[%s][UpdateOrCreate] %s", u.name, er.Error()))
			return up, er
		}
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// logrus.Error(fmt.Sprintf("[%s][UpdateOrCreate] %s", u.name, err.Error()))
		return up, err
	}

	// Create new User Point
	up.UserID = UserID
	up.Point = AddedPoint
	up, err = u.Create(up)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][UpdateOrCreate] %s", u.name, err.Error()))
		return up, err
	}

	return up, nil
}

func (u *userPoint) GetListOfUserPoint(param params.UserPointFilterParam) ([]entities.UserPoint, int, error) {
	var userPoints []entities.UserPoint

	var count int64
	u.db.Find(&userPoints).Count(&count)

	db := u.db

	if err := db.Debug().Preload("User").Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("point desc").Find(&userPoints).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GetListOfUserPoint] %s", u.name, err.Error()))
		return userPoints, int(count), err
	}

	return userPoints, int(count), nil
}
