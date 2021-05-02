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
		Exception   error
		Message     string
	}
)

func NewAPIResponse() (*APIResponse, error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		ar := strings.Split(details.Name(), ".")
		funcName := ar[len(ar)-1]
		return defaultAPIResponse(funcName), nil
	}
	return nil, nil
}

func defaultAPIResponse(command string) *APIResponse {
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

func (resp *APIResponse) SetError(errorNumber, message string, errorObject error) {
	er := &ResponseError{
		ErrorNumber: errorNumber,
		Message:     message,
		Exception:   errorObject,
	}
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
