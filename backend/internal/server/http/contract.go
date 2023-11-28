package http

import (
	"context"

	"gitlab.com/project-quiz/internal/router"
	mail "gitlab.com/project-quiz/utils/mailer"
	"gitlab.com/project-quiz/utils/minio"

	"gorm.io/gorm"
)

type Server interface {
	Run(ctx context.Context, port int)
	Done()
}

type HttpServerCfg struct {
	DB             *gorm.DB
	SMTP           mail.Mailer
	Minio          minio.MinioStorageContract
	Secret         string
	AesSecret      string
	GoogleClientID string
}

func NewServer(h *HttpServerCfg) Server {
	return &httpServer{
		router: router.NewRouter(&router.RouterCfg{
			DB:             h.DB,
			SMTP:           h.SMTP,
			Minio:          h.Minio,
			Secret:         h.Secret,
			AesSecret:      h.AesSecret,
			GoogleClientID: h.GoogleClientID,
		}).Route(),
	}
}
