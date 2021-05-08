package entity

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type APIErrorTag struct {
	TagCode   string
	TagNumber int
}

func (tag *APIErrorTag) String(Sequence int) string {
	return tag.TagCode + strconv.Itoa(tag.TagNumber) + "." + fmt.Sprintf("%02d", Sequence) + ".00"
}

func (tag *APIErrorTag) StringWithHttpError(Sequence int, err error) string {
	var code int
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	return tag.TagCode + strconv.Itoa(tag.TagNumber) + "." + strconv.Itoa(Sequence) + "." + strconv.Itoa(code)
}

func (tag *APIErrorTag) StringWithCode(Sequence int, code int) string {
	return tag.TagCode + strconv.Itoa(tag.TagNumber) + "." + strconv.Itoa(Sequence) + "." + strconv.Itoa(code)
}
