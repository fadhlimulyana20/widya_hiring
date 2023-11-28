package repository

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	e "gitlab.com/project-quiz/utils/error"
	"gorm.io/gorm"
)

var solutionTypeChoices = []string{"image", "pdf", "text", "embed_youtube"}

func solutionValid(val string, choices []string) bool {
	for _, v := range choices {
		if val == v {
			return true
		}
	}

	return false
}

type questionSolution struct {
	db   *gorm.DB
	name string
}

type QuestionSolutionRepository interface {
	// Get option detail
	Get(questionID int) (entities.QuestionSolution, error)
	// Get question solution by ID
	GetByID(ID int) (entities.QuestionSolution, error)
	// Create question solution
	Create(solution entities.QuestionSolution) (entities.QuestionSolution, error)
	// Update question solution
	Update(solution entities.QuestionSolution) (entities.QuestionSolution, error)
	// Delete question solution
	Delete(ID int) error
}

func NewQuestionSolutionRepository(db *gorm.DB) QuestionSolutionRepository {
	return &questionSolution{
		db:   db,
		name: "Question Solution Repository",
	}
}

func (q *questionSolution) Get(questionID int) (entities.QuestionSolution, error) {
	var questionSolution entities.QuestionSolution

	if err := q.db.Debug().Where("question_id = ?", questionID).First(&questionSolution).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", q.name, err.Error()))
		return questionSolution, err
	}

	return questionSolution, nil
}

func (q *questionSolution) GetByID(ID int) (entities.QuestionSolution, error) {
	var questionSolution entities.QuestionSolution

	if err := q.db.Debug().First(&questionSolution, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Get] %s", q.name, err.Error()))
		return questionSolution, err
	}

	return questionSolution, nil
}

func (q *questionSolution) Create(solution entities.QuestionSolution) (entities.QuestionSolution, error) {
	// Validate solution type
	if valid := solutionValid(solution.SolutionType, solutionTypeChoices); !valid {
		err := &e.ValueValidationError{
			Err: errors.New("invalid solution type"),
		}
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err))
		return solution, err
	}

	if err := q.db.Debug().Create(&solution).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return solution, err
	}

	return solution, nil
}

func (q *questionSolution) Update(solution entities.QuestionSolution) (entities.QuestionSolution, error) {
	// Validate solution type
	if valid := solutionValid(solution.SolutionType, solutionTypeChoices); !valid {
		err := errors.New("invalid solution type")
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err))
		return solution, err
	}

	if err := q.db.Debug().Model(&solution).Updates(&solution).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return solution, err
	}

	return solution, nil
}

func (q *questionSolution) Delete(ID int) error {
	var solution entities.QuestionSolution

	if err := q.db.Debug().Delete(&solution, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return err
	}

	return nil
}
