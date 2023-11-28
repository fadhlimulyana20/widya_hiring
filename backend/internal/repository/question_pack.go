package repository

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/utils/pagination/gorm_pagination"
	"gorm.io/gorm"
)

type questionPackRepo struct {
	db   *gorm.DB
	name string
}

type QuestionPackRepository interface {
	// Create new Question Pack
	Create(pack entities.QuestionPack) (entities.QuestionPack, error)
	// Update question pack
	Update(pack entities.QuestionPack) (entities.QuestionPack, error)
	// Get List of question Pack
	GetList(param params.QuestionPackFilterParam) ([]entities.QuestionPack, int, error)
	// Get question pack detal
	Get(ID int) (entities.QuestionPack, error)
	// Delete Question Pack
	Delete(ID int) error
	// Add Question
	AddQuestions(ID int, questionIDs []int) error
	// Delete Question
	DeleteQuestions(ID int, questionIDs []int) error
}

func NewQuestionPackRepository(db *gorm.DB) QuestionPackRepository {
	return &questionPackRepo{
		db:   db,
		name: "Question Pack Repository",
	}
}

func (q *questionPackRepo) Create(pack entities.QuestionPack) (entities.QuestionPack, error) {
	if err := q.db.Create(&pack).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return pack, err
	}
	return pack, nil
}

func (q *questionPackRepo) Update(pack entities.QuestionPack) (entities.QuestionPack, error) {
	if err := q.db.Debug().Model(&pack).Updates(&pack).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return pack, err
	}
	return pack, nil
}

func (q *questionPackRepo) GetList(param params.QuestionPackFilterParam) ([]entities.QuestionPack, int, error) {
	var packs []entities.QuestionPack
	var count int64
	db := q.db

	db.Model(&packs).Count(&count)

	if param.IsActive != nil {
		db = db.Where("is_active = ?", *param.IsActive)
	}

	if err := db.Debug().Scopes(gorm_pagination.Paginate(param.Page, param.Limit)).Order("created_at desc").Find(&packs).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return packs, 0, err
	}

	return packs, int(count), nil
}

func (q *questionPackRepo) Get(ID int) (entities.QuestionPack, error) {
	var pack entities.QuestionPack

	if err := q.db.Debug().Preload("Questions").First(&pack, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][GET] %s", q.name, err.Error()))
		return pack, err
	}

	return pack, nil
}

func (q *questionPackRepo) AddQuestions(ID int, questionIDs []int) error {
	val := ""

	for i, v := range questionIDs {
		if i != (len(questionIDs) - 1) {
			val += fmt.Sprintf("(%d,%d),", v, ID)
		} else {
			val += fmt.Sprintf("(%d,%d)", v, ID)
		}
	}

	err := q.db.Exec(fmt.Sprintf(`
		INSERT INTO question_pack_items (question_id, question_pack_id)
		select val.qid, val.qpid
		from (values %s) as val(qid, qpid)
		where not exists (
			select 1 from question_pack_items qu
			where qu.question_id = val.qid
			and qu.question_pack_id = val.qpid
		)`,
		val)).Error
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Add Question] %s", q.name, err.Error()))
		return err
	}

	return nil
}

func (q *questionPackRepo) DeleteQuestions(ID int, questionIDs []int) error {
	val := "("

	for i, v := range questionIDs {
		val += strconv.Itoa(v)
		if i != len(questionIDs)-1 {
			val += ","
		} else {
			val += ")"
		}
	}

	err := q.db.Exec(fmt.Sprintf(`
		DELETE FROM question_pack_items qp
		WHERE qp.question_pack_id = %d
		AND qp.question_id IN %s
		`,
		ID, val)).Error
	if err != nil {
		logrus.Error(fmt.Sprintf("[%s][Delete Question] %s", q.name, err.Error()))
		return err
	}

	return nil
}

func (q *questionPackRepo) Delete(ID int) error {
	var pack entities.QuestionPack

	if err := q.db.Debug().Delete(&pack, ID).Error; err != nil {
		logrus.Error(fmt.Sprintf("[%s][DELETE] %s", q.name, err.Error()))
		return err
	}

	return nil
}
