package dto

type InvoiceItem struct {
	InvoiceItemId int

	Title              string
	Note               string
	ClassificationCode string
	Quantity           int
	Rate               float32
	TaxableAmount      float32
}
