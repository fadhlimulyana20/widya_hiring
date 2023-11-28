package repository

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/minio"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"
	"gorm.io/gorm"
)

type questionRepo struct {
	db    *gorm.DB
	minio minio.MinioStorageContract
	name  string
}

type QuestionRepository interface {
	// Create a new role
	Create(question entities.Question) (entities.Question, error)
	// Update role
	Update(question entities.Question) (entities.Question, error)
	// List role
	List(param params.QuestionFilterParam) ([]entities.Question, int, error)
	// List role
	ListJoinMaterial(param params.QuestionFilterParam) ([]entities.QuestionAdminList, int, error)
	// Get Role
	Get(ID int) (entities.Question, error)
	// Get Total
	GetTotal() (int, error)
	// Delete Role
	Delete(ID int) (entities.Question, error)
	// Add tag
	AddTag(entities.Question, []entities.QuestionTag) (entities.Question, error)
	// Remove tag
	RemoveTag(entities.Question, entities.QuestionTag) (entities.Question, error)
}

func NewQuestionRepository(db *gorm.DB, m minio.MinioStorageContract) QuestionRepository {
	return &questionRepo{
		db:    db,
		minio: m,
		name:  "Question Repository",
	}
}

func (q *questionRepo) Create(question entities.Question) (entities.Question, error) {
	if err := q.db.Create(&question).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return question, err
	}

	question.GenerateTempStorageUrl(q.minio)

	return question, nil
}

func (q *questionRepo) Get(ID int) (entities.Question, error) {
	var question entities.Question
	var questionOptions []entities.QuestionOption

	db := q.db

	if err := db.Debug().Preload("QuestionTags").First(&question, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", q.name, err.Error()))
		return question, err
	}

	if err := q.db.Debug().Select("id", "question_id", "body", "is_image", "img_path", "option_value", "created_at", "updated_at").Where("question_id = ?", ID).Order("created_at asc").Find(&questionOptions).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", q.name, err.Error()))
		return question, err
	}

	question.QuestionOptions = questionOptions

	question.GenerateTempStorageUrl(q.minio)
	for i, v := range question.QuestionOptions {
		v.GenerateTempStorageUrl(q.minio)
		question.QuestionOptions[i] = v
	}

	return question, nil
}

func (q *questionRepo) List(param params.QuestionFilterParam) ([]entities.Question, int, error) {
	var questions []entities.Question

	var count int64
	db := q.db

	if param.MaterialID != 0 {
		db = db.Where("material_id = ?", param.MaterialID)
	}

	if param.Code != "" {
		db = db.Where("LOWER(code) like ?", "%"+strings.ToLower(param.Code)+"%")
	}

	switch param.Show {
	case "active":
		db = db.Where("is_active = ?", true)
	case "inactive":
		db = db.Where("is_active = ?", false)
	}

	if param.IncludePackOnly != nil {
		if *param.IncludePackOnly {
			db = db.Where("is_pack_only = ?", true).Or("is_pack_only = ?", false)
		} else {
			db = db.Where("is_pack_only = ?", false)
		}
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&questions).Count(&count).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return questions, int(count), err
	}

	return questions, int(count), nil
}

func (q *questionRepo) GetTotal() (int, error) {
	var question entities.Question

	var count int64

	if err := q.db.Debug().Model(&question).Count(&count).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return 0, err
	}

	return int(count), nil
}

func (q *questionRepo) ListJoinMaterial(param params.QuestionFilterParam) ([]entities.QuestionAdminList, int, error) {
	var questions []entities.QuestionAdminList

	var count int64
	db := q.db
	hasWhere := false
	var queryValues []interface{}
	var queryValuesPagination []interface{}

	sqlStatment := `
	select q.id, q.body, q.is_image, q.code, m."name" as material, q.img_path, q.is_active, q.contributor_id, q.created_at, q.updated_at from questions as q
	inner join materials as m
	on q.material_id = m.id
	`
	sqlStatmentCount := `
	select count(q.id) from questions as q
	inner join materials as m
	on q.material_id = m.id
	`
	if param.MaterialID != 0 {
		hasWhere = true
		sqlStatment += ` where material_id = ?`
		sqlStatmentCount += ` where material_id = ?`
		queryValues = append(queryValues, param.MaterialID)
	}

	if param.Code != "" {
		if !hasWhere {
			sqlStatment += ` where`
			sqlStatmentCount += ` where`
			sqlStatment += ` LOWER(q.code) like ?`
			sqlStatmentCount += ` LOWER(q.code) like ?`
			hasWhere = true
		} else {
			sqlStatment += ` AND LOWER(q.code) like ?`
			sqlStatmentCount += ` AND LOWER(q.code) like ?`
		}
		queryValues = append(queryValues, param.Code)
	}

	if param.ContributorID != 0 {
		if !hasWhere {
			sqlStatment += ` where`
			sqlStatmentCount += ` where`
			sqlStatment += ` q.contributor_id = ?`
			sqlStatmentCount += ` q.contributor_id = ?`
			hasWhere = true
		} else {
			sqlStatment += ` AND q.contributor_id = ?`
			sqlStatmentCount += ` AND q.contributor_id = ?`
		}
		queryValues = append(queryValues, param.ContributorID)
	}

	sqlStatment += ` order by q.created_at desc offset ? limit ?`
	queryValuesPagination = append(queryValues, (param.Page-1)*param.Limit)
	queryValuesPagination = append(queryValuesPagination, param.Limit)

	if err := db.Debug().Table("questions").Raw(sqlStatment, queryValuesPagination...).Scan(&questions).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List With Join] %s", q.name, err.Error()))
		return questions, int(count), err
	}

	if err := db.Debug().Table("questions").Raw(sqlStatmentCount, queryValues...).Scan(&count).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List With Join] %s", q.name, err.Error()))
		return questions, int(count), err
	}

	return questions, int(count), nil
}

func (q *questionRepo) Update(question entities.Question) (entities.Question, error) {
	if err := q.db.Model(&question).Updates(&question).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return question, err
	}

	return question, nil
}

func (q *questionRepo) Delete(ID int) (entities.Question, error) {
	var question entities.Question

	if err := q.db.Delete(&question, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", q.name, err.Error()))
		return question, err
	}

	return question, nil
}

func (q *questionRepo) AddTag(question entities.Question, tags []entities.QuestionTag) (entities.Question, error) {
	if err := q.db.Debug().Model(&question).Association("QuestionTags").Append(&tags); err != nil {
		log.Error(fmt.Sprintf("[%s][AddTag] %s", q.name, err.Error()))
		return question, err
	}

	return question, nil
}

func (q *questionRepo) RemoveTag(question entities.Question, tag entities.QuestionTag) (entities.Question, error) {
	if err := q.db.Debug().Model(&question).Association("QuestionTags").Delete(&tag); err != nil {
		log.Error(fmt.Sprintf("[%s][AddTag] %s", q.name, err.Error()))
		return question, err
	}

	return question, nil
}
