package middleware

import (
	"fmt"
	"net/http"
	"time"

	"gitlab.com/project-quiz/internal/appctx"
	h "gitlab.com/project-quiz/internal/handler"

	"github.com/sirupsen/logrus"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			startTime := time.Now()
			err := recover()
			if err != nil {
				hd := &h.Handler{}
				logrus.Error(err) // May be log this error? Send to sentry?
				resp := *appctx.NewResponse().WithMessage("There was an internal server error").WithCode(http.StatusInternalServerError).WithErrors(fmt.Sprintf("%v", err))
				hd.Response(w, resp, startTime, time.Now())
			}

		}()

		next.ServeHTTP(w, r)
	})
}
