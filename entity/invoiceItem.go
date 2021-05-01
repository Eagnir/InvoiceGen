package entity

import "InvoiceGen/entity/exception"

type InvoiceItem struct {
	InvoiceItemId int `gorm:"not null;primaryKey"`

	InvoiceId int `gorm:"not null;"`

	Title              string  `gorm:"not null"`
	Note               string  `gorm:"default:null"`
	ClassificationCode string  `gorm:"default:null"`
	Quantity           int     `gorm:"not null;default:1"`
	Rate               float32 `gorm:"not null;"`
	TaxableAmount      float32 `gorm:"not null;"`

	Invoice *Invoice `gorm:"references:InvoiceId;"`
	//CustomFields []*CustomField `gorm:"many2many:invoice_items_custom_fields;"`

	DefaultStruct
}

func (obj *InvoiceItem) Validate() error {
	if obj.Title == "" {
		return exception.InvoiceItem_RequiredField_Title
	}
	if obj.Rate <= 0 {
		return exception.InvoiceItem_RequiredField_Rate
	}

	return nil
}

func NewInvoiceItem(title string, rate float32) (*InvoiceItem, error) {
	if title == "" {
		return nil, exception.InvoiceItem_RequiredField_Title
	}
	if rate <= 0 {
		return nil, exception.InvoiceItem_RequiredField_Rate
	}

	u := &InvoiceItem{
		Title: title,
		Rate:  rate,
	}
	return u, nil
}
