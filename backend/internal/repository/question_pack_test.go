package repository

import (
	"testing"

	"gitlab.com/project-quiz/config"
	"gitlab.com/project-quiz/database"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/params/generics"
)

var (
	dbConfig = config.NewDbConfig().Load().Get()
	db       = database.NewSqlDB(dbConfig.Driver, dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database).ORM()
	repo     = NewQuestionPackRepository(db)
)

func TestGetListQuestionPack(t *testing.T) {
	t.Logf("Get List of Question Pack\n")

	param := params.QuestionPackFilterParam{
		GenericFilter: generics.GenericFilter{
			Q:         "",
			StartDate: "",
			EndDate:   "",
			Page:      1,
			Limit:     25,
		},
	}
	q, n, err := repo.GetList(param)
	if err != nil {
		t.Error(err)
	}

	t.Logf("pack: %v", q)
	t.Logf("count: %v", n)
}
