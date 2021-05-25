package dto

import "time"

type Client struct {
	ClientId      int
	Name          string
	Address       string
	Email         string
	ContactNumber string
	GSTNumber     string
	Country       string

	DefaultCurrency *Currency

	Invoices []*Invoice

	DefaultStruct

	// Calc Fields
	InvoiceStats *InvoiceStatistics
}

type InvoiceStatistics struct {
	TotalAmount         float32 // In Invoice Currency
	LastInvoiceDate     *time.Time
	PendingInvoiceCount int
}

func (obj *Client) CalcInvoices() {
	obj.InvoiceStats = &InvoiceStatistics{}
	for _, invoice := range obj.Invoices {
		if invoice.Status != InvoiceCancelled {
			obj.InvoiceStats.TotalAmount += invoice.InvoiceAmount
			obj.InvoiceStats.LastInvoiceDate = &invoice.CreatedAt
			if invoice.Status == InvoicePending {
				obj.InvoiceStats.PendingInvoiceCount += 1
			}
		}
	}
}
