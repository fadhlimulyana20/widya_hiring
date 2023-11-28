package entities

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/project-quiz/internal/entities/base"
	"gitlab.com/project-quiz/utils/minio"
)

type QuestionOption struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Body        string `json:"body"`
	OptionValue *bool  `json:"option_value"`
	IsImage     bool   `json:"is_image"`
	ImgPath     string `json:"img_path"`
	QuestionID  int    `json:"question_id"`
	base.Timestamp
}

func (q *QuestionOption) GenerateTempStorageUrl(m minio.MinioStorageContract) {
	if q.IsImage {
		url, err := m.GetTemporaryPublicUrl(q.ImgPath)
		if err != nil {
			logrus.Error("[Question Entity]", err.Error())
			return
		}

		q.ImgPath = url.String()
	}
}
