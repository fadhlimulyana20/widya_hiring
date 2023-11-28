package config

import (
	"os"

	_ "gitlab.com/project-quiz/utils/env"
)

func init() {
	db := NewDbConfig()
	db.Load()
}

func Env() string {
	return os.Getenv("ENV")
}
