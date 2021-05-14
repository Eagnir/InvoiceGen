package entity

import (
	"InvoiceGen/entity"
	"InvoiceGen/interface/web/api/entity/exception"
	"InvoiceGen/interface/web/api/setting"
	"InvoiceGen/usecase/adminUser"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type (
	AuthHeader struct {
		Email       string            `json:"email" validate:"email"`
		AuthToken   *entity.UUID      `json:"token" validator:"uuid"`
		AdminUserId int               `json:"-"`
		AdminUser   *entity.AdminUser `json:"-"`
	}
)

func NewAuthHeader(header http.Header, auService *adminUser.AdminUserService) (*AuthHeader, error) {
	email := strings.Trim(header.Get(setting.APIUserEmailKey), " ")
	token, err := entity.StringToUUID(strings.Trim(header.Get(setting.APITokenKey), " "))
	if err != nil {
		return nil, exception.AuthHeader_MissingAuthenticationToken
	}

	if email == "" {
		return nil, exception.AuthHeader_MissingEmail
	}

	authHeader := &AuthHeader{
		Email:     email,
		AuthToken: &token,
	}
	if auService != nil {
		usr, err := auService.VerifyAuthTokenAndEmail(authHeader.AuthToken, authHeader.Email)
		if err != nil {
			return nil, err
		}
		authHeader.AdminUser = usr
	}
	return authHeader, nil
}

func (cre *AuthHeader) ValidateSelf() (errMessages []string, errs []error) {
	valid := validator.New()
	if vErr := valid.Struct(cre); vErr != nil {
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		_ = enTranslations.RegisterDefaultTranslations(valid, trans)

		validatorErrs := vErr.(validator.ValidationErrors)
		for _, e := range validatorErrs {
			translatedError := fmt.Errorf(e.Translate(trans))
			errs = append(errs, translatedError)
			errMessages = append(errMessages, translatedError.Error())
		}
		return errMessages, errs
	}
	return nil, nil
}
