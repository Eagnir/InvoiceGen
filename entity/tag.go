package entity

import (
	"InvoiceGen/entity/exception"
	"strings"
)

type Tag struct {
	Name string `gorm:"unique;not null;primaryKey"`

	Invoices []*Invoice `gorm:"many2many:invoice_tags;"`
	Clients  []*Client  `gorm:"many2many:client_tags;"`

	DefaultStruct
}

func (obj *Tag) Validate() error {
	if obj.Name == "" {
		return exception.Tag_RequiredField_Name
	}

	return nil
}

func NewTag(name string) (*Tag, error) {
	if name == "" {
		return nil, exception.Tag_RequiredField_Name
	}

	u := &Tag{
		Name: strings.TrimSpace(strings.Title(name)),
	}
	return u, nil
}
