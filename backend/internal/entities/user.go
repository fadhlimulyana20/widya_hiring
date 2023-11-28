package entities

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities/base"
	"gorm.io/gorm"
)

type User struct {
	ID         int       `gorm:"primaryKey" json:"id,omitempty"`
	Name       string    `json:"name,omitempty" validate:"required"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	VerifiedAt time.Time `json:"verified_at"`
	Roles      []Role    `json:"roles" gorm:"many2many:user_roles"`
	base.Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Add role
	var role Role
	if err := tx.Debug().Where("name = ?", "basic").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role.Name = "basic"
			if err := tx.Debug().Create(&role).Error; err != nil {
				logrus.Error(err.Error())
			}
		}
	}

	tx.Debug().Model(u).Association("Roles").Append(&role)
	return nil
}
