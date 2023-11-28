package params

type AuthRegistrationParam struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type AuthLoginParam struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginGoogleParam struct {
	Token string `json:"token" validate:"required"`
}

type AuthRefreshTokenParam struct {
	Refresh string `json:"refresh" validate:"required"`
}

type AuthValidateEmailParams struct {
	Token string `json:"token" validate:"required"`
}

type AuthRequestValidationEmailParams struct {
	Email string `json:"email" validate:"required"`
}

type AuthRequestResetPasswordParams struct {
	Email string `json:"email" validate:"required"`
}

type AuthResetPasswordParams struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" schema:"password" validate:"required"`
}
