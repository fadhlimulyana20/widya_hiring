package params

import "gitlab.com/project-quiz/internal/params/generics"

type PremiumPackageFilterParam struct {
	UserID   int   `json:"user_id"`
	IsActive *bool `json:"is_active"`
	generics.GenericFilter
}

type PremiumPackageCreateParam struct {
	UserID     int `json:"user_id" validate:"required"`
	LongPeriod int `json:"long_period" validate:"required"`
}

type PremiumPackageUpdateParam struct {
	ID         int `json:"id"`
	UserID     int `json:"user_id" validate:"required"`
	LongPeriod int `json:"long_period" validate:"required"`
}
