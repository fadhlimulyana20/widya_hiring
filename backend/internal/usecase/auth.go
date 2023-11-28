package usecase

import (
	"errors"
	"fmt"
	"time"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/config"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"
	"gitlab.com/project-quiz/utils/jwt"
	"gitlab.com/project-quiz/utils/mailer"
	"gitlab.com/project-quiz/utils/oauth"
	"gitlab.com/project-quiz/utils/password"
	"gitlab.com/project-quiz/utils/template"
	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type auth struct {
	userRepo       repository.UserRepository
	tokenRepo      repository.TokenRepository
	name           string
	smtp           *mailer.Mailer
	googleClientID string
}

type AuthUsecase interface {
	// Registration user
	Registration(param params.AuthRegistrationParam) appctx.Response
	// Check registered user and get token
	Login(param params.AuthLoginParam) appctx.Response
	// Refresh Login Token
	Refresh(param params.AuthRefreshTokenParam) appctx.Response
	// Request Reset Password
	RequestResetPassword(param params.AuthRequestResetPasswordParams) appctx.Response
	// Reset Pasword
	ResetPassword(param params.AuthResetPasswordParams) appctx.Response
	// Request link to validate email
	RequestValidationEmail(param params.AuthRequestValidationEmailParams) appctx.Response
	// Validate email
	ValidateEmail(param params.AuthValidateEmailParams) appctx.Response
	// Authenticate google JWT
	AuthenticateGoogleJWT(token string) appctx.Response
}

func NewAuthUsecase(db *gorm.DB, smtp *mailer.Mailer, secret string, googleClientID string) AuthUsecase {
	return &auth{
		userRepo:       repository.NewUserRepository(db),
		tokenRepo:      repository.NewTokenRepository(db, secret),
		name:           "Auth Usecase",
		smtp:           smtp,
		googleClientID: googleClientID,
	}
}

func (a *auth) Registration(param params.AuthRegistrationParam) appctx.Response {
	// Copy from params to entity
	var user entities.User
	copier.Copy(&user, &param)

	// Hash Password
	hp, err := password.HashPassword(user.Password)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Registration] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	user.Password = hp

	// Create Record
	usr, err := a.userRepo.Create(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
	}

	// Generate email verification token
	token, err := a.tokenRepo.Generate(usr.ID, "registration", time.Duration(5)*time.Minute)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", a.name, err.Error()))
	}

	emailData := map[string]interface{}{
		"name": usr.Name,
		"link": fmt.Sprintf("https://kuadran.co/auth/email-confirmation/%s", token.Code),
		"year": time.Now().Year(),
	}
	template, err := template.ProcessHtmlTemplateToStr("./internal/template/email/registration.html", emailData)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", a.name, err.Error()))
	} else {
		go a.smtp.SendMail(config.SMTPFrom, usr.Email, "Pendaftaran Akun Baru", template)
	}

	return *appctx.NewResponse().WithData(usr)
}

func (a *auth) Login(param params.AuthLoginParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Login] is executed", a.name))

	// Get user data
	var user entities.User
	user, err := a.userRepo.GetByEmail(param.Email)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Login] %s", a.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors("User not found").WithCode(401)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	// Check Hash Password
	if match := password.CheckPasswordHash(param.Password, user.Password); !match {
		log.Error(fmt.Sprintf("[%s]Login] %s", a.name, "password not match"))
		return *appctx.NewResponse().WithErrors("Wrong password").WithCode(401)
	}

	// Generate JWT Token
	access, err := jwt.GenerateToken("access", int(user.ID))
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Login] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	refresh, err := jwt.GenerateToken("refresh", int(user.ID))
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Login] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	t := time.Now().UTC().Add(time.Hour*time.Duration(4) - time.Minute*time.Duration(5))
	data := map[string]interface{}{
		"token": map[string]interface{}{
			"access":  access,
			"refresh": refresh,
			"timeout": t,
		},
		"user": user,
	}

	return *appctx.NewResponse().WithData(data)
}

func (a *auth) Refresh(param params.AuthRefreshTokenParam) appctx.Response {
	// Parse Token
	ss, err := jwt.RefreshToken(param.Refresh)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Refresh] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(401)
	}

	t := time.Now().UTC().Add(time.Hour*time.Duration(4) - time.Minute*time.Duration(5))
	data := map[string]interface{}{
		"token": map[string]interface{}{
			"access":  ss,
			"refresh": param.Refresh,
			"timeout": t,
		},
	}

	return *appctx.NewResponse().WithData(data)
}

func (a *auth) RequestResetPassword(param params.AuthRequestResetPasswordParams) appctx.Response {
	// Get user data
	var user entities.User
	user, err := a.userRepo.GetByEmail(param.Email)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Request Reset Password] %s", a.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors("User not found").WithCode(406)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	// Generate email verification token
	token, err := a.tokenRepo.Generate(user.ID, "reset_password", time.Duration(5)*time.Minute)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Request Reset Password] %s", a.name, err.Error()))
	}

	emailData := map[string]interface{}{
		"name": user.Name,
		"link": fmt.Sprintf("https://kuadran.co/auth/forgot-password/reset/%s", token.Code),
		"year": time.Now().Year(),
	}
	template, err := template.ProcessHtmlTemplateToStr("./internal/template/email/reset_password.html", emailData)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Request Reset Password] %s", a.name, err.Error()))
	} else {
		go a.smtp.SendMail(config.SMTPFrom, user.Email, "Permintaan Reset Password", template)
	}

	return *appctx.NewResponse().WithCode(200).WithMessage("Request Reset password has been sent to your email")
}

func (a *auth) ResetPassword(param params.AuthResetPasswordParams) appctx.Response {
	token, valid, err := a.tokenRepo.Validate(param.Token)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(406)
	}

	if !valid {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, "invalid token"))
		return *appctx.NewResponse().WithErrors("invalid token").WithCode(406)
	}

	if token.TokenType != "reset_password" {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, "invalid token"))
		return *appctx.NewResponse().WithErrors("invalid token").WithCode(406)
	}

	var user entities.User
	user, err = a.userRepo.Get(user, token.UserID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(406)
	}

	hp, err := password.HashPassword(param.Password)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(406)
	}

	user.Password = hp
	user, err = a.userRepo.Update(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(406)
	}

	token, err = a.tokenRepo.Close(param.Token)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Reset Password] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(406)
	}

	return *appctx.NewResponse().WithMessage("Reset Password done successfully").WithCode(200)
}

func (a *auth) RequestValidationEmail(param params.AuthRequestValidationEmailParams) appctx.Response {
	// Get user data
	var user entities.User
	user, err := a.userRepo.GetByEmail(param.Email)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Login] %s", a.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *appctx.NewResponse().WithErrors("User not found").WithCode(401)
		}
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	// Generate email verification token
	token, err := a.tokenRepo.Generate(user.ID, "registration", time.Duration(5)*time.Minute)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", a.name, err.Error()))
	}

	emailData := map[string]interface{}{
		"name": user.Name,
		"link": fmt.Sprintf("https://kuadran.co/auth/email-confirmation/%s", token.Code),
		"year": time.Now().Year(),
	}
	template, err := template.ProcessHtmlTemplateToStr("./internal/template/email/registration.html", emailData)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", a.name, err.Error()))
	} else {
		go a.smtp.SendMail(config.SMTPFrom, user.Email, "Validasi Email", template)
	}

	return *appctx.NewResponse().WithCode(200).WithMessage("Request email validation has been sent to your email")
}

func (a *auth) ValidateEmail(param params.AuthValidateEmailParams) appctx.Response {
	token, valid, err := a.tokenRepo.Validate(param.Token)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Validate Email] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
	}

	if !valid {
		log.Error(fmt.Sprintf("[%s][Validate Email] %s", a.name, "invalid token"))
		return *appctx.NewResponse().WithErrors("invalid token").WithCode(400)
	}

	if token.TokenType != "registration" {
		log.Error(fmt.Sprintf("[%s][Validate Email] %s", a.name, "invalid token"))
		return *appctx.NewResponse().WithErrors("invalid token").WithCode(400)
	}

	var user entities.User
	user, err = a.userRepo.Get(user, token.UserID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Validate Email] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
	}

	user.IsVerified = true
	user.VerifiedAt = time.Now()
	user, err = a.userRepo.Update(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Validate Email] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
	}

	token, err = a.tokenRepo.Close(param.Token)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Validate Email] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
	}

	return *appctx.NewResponse().WithMessage("Verification done successfully").WithCode(200)
}

func (a *auth) AuthenticateGoogleJWT(token string) appctx.Response {
	claims, err := oauth.ValidateGoogleJWT(token, a.googleClientID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Authenticate Google JWT] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error()).WithCode(401)
	}

	user, err := a.userRepo.GetByEmail(claims.Email)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Authenticate Google JWT] %s", a.name, err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Hash Password
			hp, err := password.HashPassword("default password")
			if err != nil {
				log.Error(fmt.Sprintf("[%s][Authenticate Google JWT] %s", a.name, err.Error()))
				return *appctx.NewResponse().WithErrors(err.Error())
			}

			user = entities.User{
				Name:       fmt.Sprintf("%s %s", claims.FirstName, claims.LastName),
				Email:      claims.Email,
				IsVerified: true,
				Password:   hp,
			}
			user, err = a.userRepo.Create(user)
			if err != nil {
				log.Error(fmt.Sprintf("[%s][Authenticate Google JWT] %s", a.name, err.Error()))
				return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
			}

		} else {
			return *appctx.NewResponse().WithErrors(err.Error()).WithCode(400)
		}
	}

	// Generate JWT Token
	access, err := jwt.GenerateToken("access", int(user.ID))
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Authenticate Google JWT] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	refresh, err := jwt.GenerateToken("refresh", int(user.ID))
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Authenticate Google JWT] %s", a.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	t := time.Now().UTC().Add(time.Hour*time.Duration(4) - time.Minute*time.Duration(5))
	data := map[string]interface{}{
		"token": map[string]interface{}{
			"access":  access,
			"refresh": refresh,
			"timeout": t,
		},
		"user": user,
	}

	return *appctx.NewResponse().WithData(data)
}
