package entity

import (
	"InvoiceGen/entity/exception"
)

type Currency struct {
	CurrencyId int `gorm:"not null;primaryKey"`

	ShortName  string  `gorm:"not null;unique"`
	LongName   string  `gorm:"unique;default:null"`
	Conversion float32 `gorm:"not null;default:1"`

	Invoices  []*Invoice //default = foreignKey:InvoiceId = Currency.<FieldName>Id -> Invoice.<PrimaryKey>
	Companies []*Company `gorm:"foreignKey:DefaultCurrencyId;"`
	Clients   []*Client  `gorm:"foreignKey:DefaultCurrencyId"`

	DefaultStruct
}

func (obj *Currency) Validate() error {
	if obj.ShortName == "" {
		return exception.Currency_RequiredField_ShortName
	}

	return nil
}

func NewCurrency(shortName string, conversion float32) (*Currency, error) {
	if shortName == "" {
		return nil, exception.Currency_RequiredField_ShortName
	}
	u := &Currency{
		ShortName:  shortName,
		Conversion: conversion,
	}
	return u, nil
}
