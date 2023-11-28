package repository

import (
	"fmt"
	"strings"

	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type QuestionTagRepo struct {
	db   *gorm.DB
	name string
}

type QuestionTagRepository interface {
	// Create a new qeuestion tag
	Create(tag entities.QuestionTag) (entities.QuestionTag, error)
	// Update qeuestion tag
	Update(tag entities.QuestionTag) (entities.QuestionTag, error)
	// List qeuestion tag
	List(param params.QuestionTagFilter) ([]entities.QuestionTag, int, error)
	// List of question tag with IDs
	ListIn([]int) ([]entities.QuestionTag, int, error)
	// Get qeuestion tag
	Get(ID int) (entities.QuestionTag, error)
	// Get qeuestion tag By name
	GetByName(name string) (entities.QuestionTag, error)
	// Delete qeuestion tag
	Delete(ID int) (entities.QuestionTag, error)
}

// Create new question tag repository instance
func NewQuestionTagRepository(db *gorm.DB) QuestionTagRepository {
	return &QuestionTagRepo{
		db:   db,
		name: "Question Tag Repository",
	}
}

func (q *QuestionTagRepo) Create(tag entities.QuestionTag) (entities.QuestionTag, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", q.name))

	if err := q.db.Create(&tag).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return tag, err
	}

	return tag, nil
}

func (q *QuestionTagRepo) Get(ID int) (entities.QuestionTag, error) {
	log.Info(fmt.Sprintf("[%s][Get] is executed", q.name))
	var tag entities.QuestionTag

	db := q.db

	if err := db.Debug().First(&tag, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", q.name, err.Error()))
		return tag, err
	}

	return tag, nil
}

func (q *QuestionTagRepo) List(param params.QuestionTagFilter) ([]entities.QuestionTag, int, error) {
	log.Info(fmt.Sprintf("[%s][Update] is executed", q.name))

	var tags []entities.QuestionTag

	var count int64
	q.db.Find(&tags).Count(&count)

	db := q.db

	if param.Q != "" {
		db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(param.Q)+"%")
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&tags).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return tags, int(count), err
	}

	return tags, int(count), nil
}

func (q *QuestionTagRepo) Update(tag entities.QuestionTag) (entities.QuestionTag, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", q.name))

	if err := q.db.Model(&tag).Updates(&tag).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return tag, err
	}

	return tag, nil
}

func (q *QuestionTagRepo) Delete(ID int) (entities.QuestionTag, error) {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", q.name))
	var tag entities.QuestionTag

	if err := q.db.Delete(&tag, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", q.name, err.Error()))
		return tag, err
	}

	return tag, nil
}

func (q *QuestionTagRepo) GetByName(name string) (entities.QuestionTag, error) {
	log.Info(fmt.Sprintf("[%s][GetByName] is executed", q.name))
	var tag entities.QuestionTag

	if err := q.db.Debug().Where("name = ?", name).First(&tag).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GetByName] %s", q.name, err.Error()))
		return tag, err
	}

	return tag, nil
}

func (q *QuestionTagRepo) ListIn(IDs []int) ([]entities.QuestionTag, int, error) {
	var tags []entities.QuestionTag
	var count int64

	if err := q.db.Debug().Where("id IN ?", IDs).Order("created_at DESC").Find(&tags).Count(&count).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][ListIn] %s", q.name, err.Error()))
		return tags, 0, err
	}

	return tags, int(count), nil
}
