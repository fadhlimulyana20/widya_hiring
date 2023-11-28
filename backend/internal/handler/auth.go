package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/usecase"
	"gitlab.com/project-quiz/utils/json"
	"gitlab.com/project-quiz/utils/mailer"
	"gitlab.com/project-quiz/utils/validator"
	"gorm.io/gorm"
)

type auth struct {
	handler     Handler
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
	name        string
}

type AuthHandler interface {
	// Register a new user
	Register(w http.ResponseWriter, r *http.Request)
	// Login user
	Login(w http.ResponseWriter, r *http.Request)
	// Refresh JWT Token
	Refresh(w http.ResponseWriter, r *http.Request)
	// Request Reset Password
	RequestResetPassword(w http.ResponseWriter, r *http.Request)
	// Reset Password using token
	ResetPassword(w http.ResponseWriter, r *http.Request)
	// Request Link sent to email to validate email
	RequestValidationEmail(w http.ResponseWriter, r *http.Request)
	// Validate email using token sent to email
	ValidateEmail(w http.ResponseWriter, r *http.Request)
	// Get Authenticated User
	GetAuthenticatedUser(w http.ResponseWriter, r *http.Request)
	// Update password
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	// Update Account
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	// Auth with Google
	AuthWithGoogle(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(db *gorm.DB, smtp *mailer.Mailer, secret string, googleClientID string) AuthHandler {
	return &auth{
		userUsecase: usecase.NewUserUsecase(db),
		authUsecase: usecase.NewAuthUsecase(db, smtp, secret, googleClientID),
		name:        "AUTH HANDLER",
	}
}

func (a *auth) Register(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.AuthRegistrationParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] Cannot decode json", a.name))
		ctx = ctx.WithErrors(err.Error())
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(fmt.Sprintf("[%s] %s", a.name, err.Error()))
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	// Manual Validation
	if param.Password != param.ConfirmPassword {
		logrus.Error(fmt.Sprintf("[%s] %s", a.name, "password validation error"))
		ctx = ctx.WithErrors("password not match").WithCode(http.StatusBadRequest)

	}

	if len(ctx.Errors) > 0 {
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.Registration(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) Login(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthLoginParam
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.Login(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) Refresh(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthRefreshTokenParam
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.Refresh(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthRequestResetPasswordParams
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.RequestResetPassword(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) ResetPassword(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthResetPasswordParams
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.ResetPassword(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) RequestValidationEmail(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthRequestValidationEmailParams
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.RequestValidationEmail(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) ValidateEmail(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Signin] is executed", a.name))
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthValidateEmailParams
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.ValidateEmail(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Me] is executed", a.name))
	startTime := time.Now()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	resp := a.userUsecase.Get(userID)

	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	userID, _ := strconv.Atoi(r.Header.Get("user"))
	ctx := appctx.NewResponse()

	// Decode data
	var param params.UserUpdatePassword
	param.UserID = userID
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.userUsecase.UpdatePassword(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	var param params.UserUpdateParam
	ctx := appctx.NewResponse()

	userID, _ := strconv.Atoi(r.Header.Get("user"))
	param.ID = userID

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
	}

	if len(ctx.Errors) > 0 {
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.userUsecase.Update(param)
	a.handler.Response(w, resp, startTime, time.Now())
}

func (a *auth) AuthWithGoogle(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	ctx := appctx.NewResponse()

	// Decode data
	var param params.AuthLoginGoogleParam
	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
	}

	// Validate Data
	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error()).WithCode(400)
		a.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := a.authUsecase.AuthenticateGoogleJWT(param.Token)
	a.handler.Response(w, resp, startTime, time.Now())
}
