package http

import (
	"context"
	"time"

	"gitlab.com/project-quiz/config"
	"gitlab.com/project-quiz/database"
	h "gitlab.com/project-quiz/internal/server/http"
	mail "gitlab.com/project-quiz/utils/mailer"
	"gitlab.com/project-quiz/utils/minio"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func StartServer(ctx context.Context, port int) {
	dbConfig := config.NewDbConfig().Load().Get()
	db := database.NewSqlDB(dbConfig.Driver, dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database).ORM()

	smtpConfig := config.NewSMTPConfig().Load().Get()
	smtp := mail.NewMailer(smtpConfig.Host, smtpConfig.Port, smtpConfig.AuthEmail, smtpConfig.Password).GetMailer()

	minioConfig := config.NewMinioCfg().Load()
	minio := minio.NewMinioStorage(minioConfig.Endpoint, minioConfig.AccessKeyID, minioConfig.SecretAccessKey, minioConfig.BucketName, minioConfig.UseSSL)

	secretKey := config.NewSecretCfg().Load()
	oauth := config.NewOauthConfig().Load()

	// Sentry
	sentryCfg := config.NewSentryConfig().Load()
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryCfg.SentryDSN,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		logrus.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("It works!")

	ht := h.NewServer(&h.HttpServerCfg{
		DB:             db,
		SMTP:           *smtp,
		Minio:          minio,
		Secret:         secretKey.Key,
		AesSecret:      secretKey.AesKey,
		GoogleClientID: oauth.GoogleClientID,
	})
	defer ht.Done()
	ht.Run(ctx, port)

	// return
	// http.ListenAndServe(":3000", r)
}

func ServerCmd(ctx context.Context) *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start HTTP server",
		Long:  "Start HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")
			if port == 0 {
				port = 3000
			}
			StartServer(ctx, port)
		},
	}

	serverCmd.PersistentFlags().Int("port", 3000, "step for rolling back migration")

	return serverCmd
}
