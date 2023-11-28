package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Oauth struct {
	GoogleClientID string
}

type OauthConfig interface {
	Load() *Oauth
}

func NewOauthConfig() OauthConfig {
	return &Oauth{}
}

func (o *Oauth) Load() *Oauth {
	o.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	logrus.Debug(o.GoogleClientID)
	return o
}

func (o *Oauth) Get() *Oauth {
	return o
}
