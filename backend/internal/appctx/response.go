package appctx

import (
	"errors"
	"math"
	"net/http"
	"time"

	apperror "gitlab.com/project-quiz/utils/error"

	"gorm.io/gorm"
)

type Response struct {
	Code        int           `json:"code"`
	Message     interface{}   `json:"message,omitempty"`
	Errors      []interface{} `json:"errors,omitempty"`
	Data        interface{}   `json:"data,omitempty"`
	Meta        *MetaData     `json:"meta,omitempty"`
	ProcessTime int64         `json:"process_time"`
}

type MetaData struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPage  int64 `json:"total_page"`
	TotalCount int64 `json:"total_count"`
}

func NewResponse() *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "Success retrieving data",
	}
}

func (r *Response) WithErrorContexts(err error) *Response {
	r.Code = http.StatusInternalServerError
	r.Message = "Failed retrieving data"

	errCtx := err.(*apperror.ErrorWithContext)
	r.Errors = append(r.Errors, errCtx.Ctx)
	return r
}

func (r *Response) WithErrors(err string) *Response {
	r.Code = http.StatusInternalServerError
	r.Message = "Failed retrieving data"
	r.Errors = append(r.Errors, err)
	return r
}

func (r *Response) WithErrorObj(err error) *Response {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.Code = http.StatusNotFound
		r.Message = "Failed retrieving data"
		r.Errors = append(r.Errors, err.Error())
	} else {
		r.Code = http.StatusBadRequest
		r.Message = "Failed retrieving data"
		r.Errors = append(r.Errors, err.Error())
	}

	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) WithMeta(page int64, limit int64, totalCount int64) *Response {
	countf := float64(totalCount)
	limitf := float64(limit)
	r.Meta = &MetaData{
		Page:       page,
		Limit:      limit,
		TotalPage:  int64(math.Ceil(countf / limitf)),
		TotalCount: totalCount,
	}
	return r
}

func (r *Response) WithMetaObj(meta MetaData) *Response {
	r.Meta = &meta
	return r
}

func (r *Response) WithProcessTime(timeStart time.Time, timeEnd time.Time) *Response {
	timeDuration := timeEnd.Sub(timeStart)
	r.ProcessTime = timeDuration.Milliseconds()
	return r
}
