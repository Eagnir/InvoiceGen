package dto

type Invoice struct {
	InvoiceNumber string

	TaxableAmount float32
	TaxPercentage float32
	TaxPayable    float32
	InvoiceAmount float32
	Status        InvoiceStatus

	InvoiceItems []*InvoiceItem

	Client   *Client
	Currency *Currency
	TaxGroup *TaxGroup

	DefaultStruct
}
type InvoiceStatus int

const (
	InvoiceCreated = iota
	InvoicePending
	InvoicePaid
	InvoiceCancelled
)
