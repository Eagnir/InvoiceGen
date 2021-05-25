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

	Invoice *Invoice

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

func NewInvoiceItem(title string, quantity int, rate float32) (*InvoiceItem, error) {
	if title == "" {
		return nil, exception.InvoiceItem_RequiredField_Title
	}
	if quantity <= 0 {
		return nil, exception.InvoiceItem_RequiredField_Quantity
	}
	if rate <= 0 {
		return nil, exception.InvoiceItem_RequiredField_Rate
	}

	u := &InvoiceItem{
		Title:    title,
		Rate:     rate,
		Quantity: quantity,
	}
	u.CalcAmounts()
	return u, nil
}

func (obj *InvoiceItem) CalcAmounts() {
	obj.TaxableAmount = obj.Rate * float32(obj.Quantity)
}

func (obj *InvoiceItem) ChangeQuantity(quantity int) {
	obj.Quantity = quantity
	obj.CalcAmounts()
}

func (obj *InvoiceItem) ChangeRate(rate float32) {
	obj.Rate = rate
	obj.CalcAmounts()
}

func (obj *InvoiceItem) SetNote(note string) {
	obj.Note = note
}

func (obj *InvoiceItem) SetClassificationCode(code string) {
	obj.ClassificationCode = code
}
