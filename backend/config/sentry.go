package config

import (
	"os"
)

type Sentry struct {
	SentryDSN string
}

type SentryConfig interface {
	Load() *Sentry
}

func NewSentryConfig() SentryConfig {
	return &Sentry{}
}

func (s *Sentry) Load() *Sentry {
	s.SentryDSN = os.Getenv("SENTRY_DSN")
	return s
}
