package entities

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities/base"
	"gitlab.com/project-quiz/utils/minio"
	"gorm.io/gorm"
)

type Question struct {
	ID              int              `json:"id" gorm:"primaryKey"`
	Code            string           `json:"code"`
	Body            string           `json:"body"`
	MaterialID      int              `json:"material_id"`
	IsImage         bool             `json:"is_image"`
	ImgPath         string           `json:"img_path"`
	QuestionOptions []QuestionOption `json:"question_options,omitempty"`
	IsActive        *bool            `json:"is_active" gorm:"default:true"`
	IsPackOnly      *bool            `json:"is_pack_only" gorm:"default:false"`
	QuestionTags    []QuestionTag    `json:"tags" gorm:"many2many:question_tags;foreginKey:ID;joinForeignKey:QuestionID;references:ID;joinReferences:TagID"`
	ImgPlacementUrl string           `json:"img_placement_url"`
	ContributorID   int              `json:"contributor_id"`
	base.Timestamp
}

type QuestionAdminList struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Body          string `json:"body"`
	Material      string `json:"material"`
	IsImage       bool   `json:"is_image"`
	ImgPath       string `json:"img_path"`
	IsActive      *bool  `json:"is_active" gorm:"default:true"`
	Code          string `json:"code"`
	ContributorID int    `json:"contributor_id"`
	base.Timestamp
}

func (q *Question) GenerateTempStorageUrl(m minio.MinioStorageContract) {
	if q.IsImage {
		url, err := m.GetTemporaryPublicUrl(q.ImgPath)
		if err != nil {
			logrus.Error("[Question Entity]", err.Error())
			return
		}

		q.ImgPath = url.String()
	}
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	var countQuestion int64
	tx.Debug().Model(q).Select("id").Count(&countQuestion)
	questionCode := fmt.Sprintf("kq%v", countQuestion+1)
	q.Code = questionCode
	tx.Statement.SetColumn("code", questionCode)
	return nil
}
