package entity

import (
	"InvoiceGen/interface/web/api/setting"
	"math"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
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
		context    echo.Context
	}

	APIPagination struct {
		PageSize   int
		Page       int
		TotalPages int
	}

	ResponseError struct {
		ErrorNumber string
		Exceptions  []interface{}
		Messages    []string
	}
)

func NewAPIResponse(command string, c echo.Context) *APIResponse {
	resp := &APIResponse{
		Command:    command,
		Status:     setting.StatusFailure,
		StatusText: setting.StatusFailureText,
		DateTime:   time.Now().UTC(),
		Pagination: &APIPagination{
			PageSize: setting.DefaultPageSize,
			Page:     1,
		},
		context: c,
	}
	return resp
}

func (resp *APIResponse) ReturnWith(status APIResponseStatus, message string) error {
	resp.SetStatus(status)
	resp.Message = message
	return resp.Return()
}

func (resp *APIResponse) Return() error {
	resp.Sanitize()
	return resp.context.JSON(http.StatusOK, resp)
}

func (resp *APIResponse) Sanitize() {
	// Sanitization of Response Object
	if resp.Error != nil {
		if len(resp.Error.Exceptions) <= 0 {
			resp.Error.Exceptions = nil
		}
		if len(resp.Error.Messages) <= 0 {
			resp.Error.Messages = nil
		}
	}
	if resp.Data == nil || len(resp.Data) <= 0 {
		resp.Data = nil
		resp.Pagination = nil
	}

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

func (resp *APIResponse) SetErrorAndReturn(errorNumber string, userMessage string, errorObjects interface{}, status APIResponseStatus) error {
	resp.SetError(errorNumber, userMessage, errorObjects, status)
	return resp.Return()
}

func (resp *APIResponse) SetError(errorNumber string, userMessage string, errorObjects interface{}, status APIResponseStatus) {
	er := &ResponseError{
		ErrorNumber: errorNumber,
		Messages:    []string{},
	}

	if errorObjects != nil {
		var arr []interface{}
		switch reflect.TypeOf(errorObjects).Kind() {
		case reflect.Slice:
			for _, e := range errorObjects.([]error) {
				er.Exceptions = append(er.Exceptions, e)
			}

		default:
			arr = append(arr, errorObjects)
			er.Exceptions = arr
		}

		for _, e := range er.Exceptions {
			if v, ok := e.(error); ok {
				er.Messages = append(er.Messages, v.Error())
			}
		}
	}

	resp.Message = userMessage
	resp.SetStatus(status)
	resp.Error = er
}

func (resp *APIResponse) SetResponse(message string, status APIResponseStatus, data interface{}) error {
	if data != nil {
		rt := reflect.TypeOf(data)
		switch rt.Kind() {
		case reflect.Slice:
			resp.SetData(data.([]interface{}))
		default:
			resp.AddData(data)
		}
	}
	resp.Message = message
	resp.SetStatus(status)
	return nil
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
