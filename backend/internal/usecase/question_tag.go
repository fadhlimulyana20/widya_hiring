package usecase

import (
	"fmt"

	"gitlab.com/project-quiz/internal/appctx"
	"gitlab.com/project-quiz/internal/entities"
	"gitlab.com/project-quiz/internal/params"
	"gitlab.com/project-quiz/internal/repository"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type questionTag struct {
	questionTagRepo repository.QuestionTagRepository
	name            string
}

type QuestionTagUsecase interface {
	// Create new question tag
	Create(param params.QuestionTagCreateParam) appctx.Response
	// Edit a question tag
	Update(param params.QuestionTagUpdateParam) appctx.Response
	// Get question tag list
	List(param params.QuestionTagFilter) appctx.Response
	// Get detail question tag
	Detail(ID int) appctx.Response
	// Delete question tag
	Delete(ID int) appctx.Response
	// // Assign Role to user
	// Assign(userID int, roleName string) appctx.Response
	// // Revoke Role from user
	// Revoke(userID int, roleName string) appctx.Response
}

func NewQuestionTagUsecase(db *gorm.DB) QuestionTagUsecase {
	return &questionTag{
		questionTagRepo: repository.NewQuestionTagRepository(db),
		name:            "Question Tag Usecase",
	}
}

func (q *questionTag) Create(param params.QuestionTagCreateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", q.name))

	tag := entities.QuestionTag{
		Name: param.Name,
	}

	data, err := q.questionTagRepo.Create(tag)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(data)
}

func (q *questionTag) Update(param params.QuestionTagUpdateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Update] is executed", q.name))

	// get tag
	var tag entities.QuestionTag
	tag, err := q.questionTagRepo.Get(param.ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	copier.Copy(&tag, &param)

	usr, err := q.questionTagRepo.Update(tag)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}

func (q *questionTag) List(param params.QuestionTagFilter) appctx.Response {
	log.Info(fmt.Sprintf("[%s][List] is executed", q.name))

	// get role list
	var tags []entities.QuestionTag
	tags, count, err := q.questionTagRepo.List(param)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(tags).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (q *questionTag) Detail(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Detail] is executed", q.name))

	// get question tag
	var tag entities.QuestionTag
	tag, err := q.questionTagRepo.Get(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Detail] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(tag)
}

func (q *questionTag) Delete(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", q.name))

	// delete question tag
	_, err := q.questionTagRepo.Delete(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", q.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("question tag deleted sucessfully")
}

// func (r *role) Assign(userID int, roleName string) appctx.Response {
// 	log.Info(fmt.Sprintf("[%s][Assign] is executed", r.name))

// 	var role entities.Role
// 	role, err := r.repo.GetByName(roleName)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	var user entities.User
// 	user, err = r.userRepo.Get(user, userID)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	_, err = r.userRepo.AddRole(user, role)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	return *appctx.NewResponse().WithMessage("role assigned sucessfully")
// }

// func (r *role) Revoke(userID int, roleName string) appctx.Response {
// 	log.Info(fmt.Sprintf("[%s][Revoke] is executed", r.name))

// 	var role entities.Role
// 	role, err := r.repo.GetByName(roleName)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	var user entities.User
// 	user, err = r.userRepo.Get(user, userID)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	_, err = r.userRepo.RemoveRole(user, role)
// 	if err != nil {
// 		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
// 		return *appctx.NewResponse().WithErrorObj(err)
// 	}

// 	return *appctx.NewResponse().WithMessage("role revoked sucessfully")
// }
