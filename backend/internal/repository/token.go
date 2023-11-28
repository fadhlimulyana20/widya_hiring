package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/utils/encryption"
	"gorm.io/gorm"
)

type tokenRepo struct {
	db     *gorm.DB
	name   string
	secret string
}

type TokenRepository interface {
	Generate(userID int, tokenType string, duration time.Duration) (entities.Token, error)
	Validate(token string) (entities.Token, bool, error)
	Close(token string) (entities.Token, error)
}

func NewTokenRepository(db *gorm.DB, secret string) TokenRepository {
	return &tokenRepo{
		db:     db,
		name:   "Token Repository",
		secret: secret,
	}
}

func (t *tokenRepo) Generate(userID int, tokenType string, duration time.Duration) (entities.Token, error) {
	var token entities.Token

	validUntil := time.Now().Add(duration)
	validUntilInt := validUntil.Unix()
	data := fmt.Sprintf("%d#%s#%d", userID, tokenType, validUntilInt)

	aesEncrypt := encryption.NewAESEncrypt(t.secret)
	c, err := aesEncrypt.Encrypt(data)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Generate] %s", t.name, err.Error()))
		return token, err
	}

	token.UserID = userID
	token.ValidUntil = validUntilInt
	token.Code = c
	token.TokenType = tokenType

	if err := t.db.Debug().Create(&token).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Generate] %s", t.name, err.Error()))
		return token, err
	}

	return token, nil
}

func (t *tokenRepo) Validate(tokenCode string) (entities.Token, bool, error) {
	var token entities.Token

	if err := t.db.Debug().Where("code = ?", tokenCode).First(&token).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Validate] %s", t.name, err.Error()))
		return token, false, err
	}

	if token.IsCompleted {
		logrus.Error(fmt.Sprintf("[%s][Validate] %s", t.name, "token is completed"))
		return token, false, errors.New("token is completed")
	}

	aesEncrypt := encryption.NewAESEncrypt(t.secret)
	p, err := aesEncrypt.Decrypt(tokenCode)
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Validate] %s", t.name, err.Error()))
		return token, false, err
	}

	data := strings.Split(p, "#")
	dataUserID, _ := strconv.Atoi(data[0])
	dataTokenType := data[1]
	dataValidUntil, _ := strconv.Atoi(data[2])
	dataValidUntilInt64 := int64(dataValidUntil)

	if token.UserID != dataUserID || token.TokenType != dataTokenType || token.ValidUntil != dataValidUntilInt64 {
		logrus.Error(fmt.Sprintf("[%s][Validate] %s", t.name, "invalid token"))
		return token, false, errors.New("invalid token")
	}

	if token.ValidUntil < time.Now().Unix() {
		logrus.Error(fmt.Sprintf("[%s][Validate] %s", t.name, "token is expired"))
		return token, false, errors.New("token is expired")
	}

	return token, true, nil
}

func (t *tokenRepo) Close(tokenCode string) (entities.Token, error) {
	var token entities.Token

	if err := t.db.Debug().Model(&token).Where("code = ?", tokenCode).Update("is_completed", true).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Validate] %s", t.name, err.Error()))
		return token, err
	}

	return token, nil
}
