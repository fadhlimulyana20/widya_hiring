package params

import "gitlab.com/project-quiz/internal/params/generics"

type UserCreateParam struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateParam struct {
	ID    int
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UserUpdatePassword struct {
	UserID      int    `json:"user_id" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
}

type UserListParams struct {
	generics.GenericFilter
}
