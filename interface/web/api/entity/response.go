package entity

import (
	"InvoiceGen/interface/web/api/setting"
	"math"
	"runtime"
	"strings"
	"time"
)

type (
	APIResponseStatus     int
	APIResponseStatusText string

	APIResponse struct {
		Command    string
		Pagination *APIPagination
		Data       []interface{}
		Status     APIResponseStatus
		StatusText APIResponseStatusText
		Message    string
		Error      *ResponseError
		DateTime   time.Time
	}

	APIPagination struct {
		PageSize   int
		Page       int
		TotalPages int
	}

	ResponseError struct {
		ErrorNumber string
		Exceptions  []error
		Message     string
	}
)

func newAPIResponse() (*APIResponse, error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		ar := strings.Split(details.Name(), ".")
		funcName := ar[len(ar)-1]
		return NewAPIResponse(funcName), nil
	}
	return nil, nil
}

func NewAPIResponse(command string) *APIResponse {
	resp := &APIResponse{
		Command:    command,
		Status:     setting.StatusFailure,
		StatusText: setting.StatusFailureText,
		DateTime:   time.Now().UTC(),
		Pagination: &APIPagination{
			PageSize: setting.DefaultPageSize,
			Page:     1,
		},
	}
	return resp
}

func (resp *APIResponse) SetStatus(status APIResponseStatus) {
	resp.Status = status
	switch status {
	case setting.StatusSuccess:
		resp.StatusText = setting.StatusSuccessText
	case setting.StatusWarning:
		resp.StatusText = setting.StatusWarningText
	case setting.StatusFailure:
		resp.StatusText = setting.StatusFailureText
	case setting.StatusFatalError:
		resp.StatusText = setting.StatusFatalErrorText
	}
}

func (resp *APIResponse) SetFatalErrors(errorNumber string, messages []string, errorObjects []error) {
	resp.SetErrorWithStatus(errorNumber, strings.Join(messages, "<br>"), errorObjects, setting.StatusFatalError)
}

func (resp *APIResponse) SetFatalError(errorNumber, message string, errorObject error) {
	resp.SetErrorWithStatus(errorNumber, message, []error{errorObject}, setting.StatusFatalError)
}

func (resp *APIResponse) SetFailureErrors(errorNumber string, messages []string, errorObjects []error) {
	resp.SetErrorWithStatus(errorNumber, strings.Join(messages, "<br>"), errorObjects, setting.StatusFailure)
}

func (resp *APIResponse) SetFailureError(errorNumber, message string, errorObject error) {
	resp.SetErrorWithStatus(errorNumber, message, []error{errorObject}, setting.StatusFailure)
}

func (resp *APIResponse) SetErrorWithStatus(errorNumber, message string, errorObjects []error, status APIResponseStatus) {
	er := &ResponseError{
		ErrorNumber: errorNumber,
		Message:     message,
		Exceptions:  errorObjects,
	}
	resp.Message = er.Message
	resp.SetStatus(status)
	resp.Error = er
}

func (resp *APIResponse) SetErrorOnly(errorNumber, message string, errorObjects []error) {
	er := &ResponseError{
		ErrorNumber: errorNumber,
		Message:     message,
		Exceptions:  errorObjects,
	}
	resp.Message = er.Message
	resp.Error = er
}

func (resp *APIResponse) AddData(obj interface{}) error {
	resp.Data = append(resp.Data, obj)
	return nil
}

func (resp *APIResponse) SetData(objs []interface{}) error {
	resp.Data = objs
	return nil
}

func (pagniation *APIPagination) PageIndex() (int, error) {
	if pagniation.Page < 0 {
		return 0, nil
	}
	return pagniation.Page - 1, nil
}
func (pagniation *APIPagination) CalcTotalPages(totalObjects int) error {
	pagniation.TotalPages = int(math.Ceil(float64(totalObjects) / float64(pagniation.PageSize)))
	return nil
}
func (pagination *APIPagination) Reset() error {
	pagination.Page = 1
	pagination.TotalPages = 0
	pagination.PageSize = setting.DefaultPageSize
	return nil
}
