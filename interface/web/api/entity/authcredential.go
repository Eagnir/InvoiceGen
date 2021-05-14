package entity

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type (
	AuthCredential struct {
		Email    string `json:"email" validate:"required,omitempty,email"`
		Password string `json:"password" validate:"required"`
	}
)

func (cre *AuthCredential) ValidateSelf() (errMessages []string, errs []error) {
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
