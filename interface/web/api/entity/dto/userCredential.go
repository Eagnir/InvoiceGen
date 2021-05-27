package dto

import (
	"InvoiceGen/entity"
)

type UserCredential struct {
	AdminUserId int `json:"ucid"`

	Name  string `json:"name"`
	Email string `json:"email"`

	AuthToken *entity.UUID `json:"token"`

	Company *company `json:"company"`
}

type company struct {
	CompanyId       int              `json:"cid"`
	Name            string           `json:"name"`
	ContactNumber   string           `json:"tel"`
	DefaultCurrency *defaultCurrency `json:"defaultCurrency"`
}

type defaultCurrency struct {
	ShortName  string  `json:"shortName"`
	LongName   string  `json:"longName"`
	Conversion float32 `json:"conversion"`
}
