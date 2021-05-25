package entity

import (
	"InvoiceGen/entity/exception"
	"strconv"
	"strings"
	"time"
)

type Invoice struct {
	InvoiceId int `gorm:"not null;primaryKey"`

	InvoiceNumber string `gorm:"unique;not null"`

	AdminUserId int `gorm:"not null"`
	ClientId    int `gorm:"not null"`
	CompanyId   int `gorm:"not null"`
	CurrencyId  int `gorm:"not null"`
	TaxGroupId  *int

	TaxableAmount float32       `gorm:"not null"`
	TaxPercentage float32       `gorm:"default:0"`
	TaxPayable    float32       `gorm:"not null"`
	InvoiceAmount float32       `gorm:"not null"`
	Status        InvoiceStatus `gorm:"not null"`

	InvoiceItems []*InvoiceItem `gorm:"not null;foreignKey:InvoiceId"`

	AdminUser *AdminUser
	Client    *Client
	Company   *Company
	Currency  *Currency
	TaxGroup  *TaxGroup

	Tags []*Tag `gorm:"many2many:invoice_tags;"`

	DefaultStruct
}

type InvoiceStatus int

const (
	InvoiceCreated = iota
	InvoicePending
	InvoicePaid
	InvoiceCancelled
)

func (status InvoiceStatus) String() string {
	switch status {
	case InvoiceCreated:
		return "Created"
	case InvoicePending:
		return "Pending"
	case InvoicePaid:
		return "Paid"
	case InvoiceCancelled:
		return "Cancelled"
	}
	return ""
}

func (obj *Invoice) Validate() error {
	if obj.AdminUserId == 0 {
		return exception.Invoice_RequiredField_CompanyId
	}

	return nil
}

func NewInvoice(status InvoiceStatus, adminUser *AdminUser) (*Invoice, error) {
	u := &Invoice{
		AdminUser: adminUser,
		Status:    status,
	}
	u.AutoFillInvoiceNumber(0)
	return u, nil
}

func (invoice *Invoice) AutoFillInvoiceNumber(todaysInvoiceCount int) *Invoice {
	invoice.InvoiceNumber = time.Now().Format("020106") + strconv.Itoa(todaysInvoiceCount+1)
	return invoice
}

func (invoice *Invoice) SetInvoiceNumber(invoiceNumber string) *Invoice {
	invoice.InvoiceNumber = invoiceNumber
	return invoice
}

func (invoice *Invoice) SetDetails(company *Company, client *Client, currency *Currency, taxGroup *TaxGroup) *Invoice {
	invoice.SetCompany(company).SetClient(client).SetCurrency(currency).SetTaxGroup(taxGroup)
	return invoice
}

func (invoice *Invoice) SwitchAdminUser(adminUser *AdminUser) *Invoice {
	invoice.AdminUser = adminUser
	return invoice
}

func (invoice *Invoice) SetClient(client *Client) *Invoice {
	invoice.Client = client
	return invoice
}

func (invoice *Invoice) SetCompany(company *Company) *Invoice {
	invoice.Company = company
	return invoice
}

func (invoice *Invoice) SetCurrency(currency *Currency) *Invoice {
	invoice.Currency = currency
	return invoice
}

func (invoice *Invoice) SetTaxGroup(taxGroup *TaxGroup) *Invoice {
	invoice.TaxGroup = taxGroup
	invoice.CalcAmounts()
	return invoice
}

func (invoice *Invoice) SetInvoiceItems(invoiceItems []*InvoiceItem) *Invoice {
	invoice.InvoiceItems = invoiceItems
	invoice.CalcAmounts()
	return invoice
}

func (invoice *Invoice) AddInvoiceItem(invoiceItem *InvoiceItem) *Invoice {
	invoice.InvoiceItems = append(invoice.InvoiceItems, invoiceItem)
	invoice.CalcAmounts()
	return invoice
}

func (invoice *Invoice) CalcAmounts() *Invoice {
	invoice.TaxableAmount = 0.0
	for _, item := range invoice.InvoiceItems {
		invoice.TaxableAmount += item.TaxableAmount
	}
	invoice.TaxPercentage = 0.0
	if invoice.TaxGroup != nil {
		for _, item := range invoice.TaxGroup.Taxes {
			invoice.TaxPercentage += item.Percentage
		}
	}
	invoice.TaxPayable = invoice.TaxableAmount * invoice.TaxPercentage / 100
	invoice.InvoiceAmount = invoice.TaxableAmount + invoice.TaxPayable
	return invoice
}

func (invoice *Invoice) AddTagByName(name string) error {
	t, ex := NewTag(name)
	if ex == nil {
		invoice.Tags = append(invoice.Tags, t)
		return nil
	}
	return ex
}

func (invoice *Invoice) AddTag(tag *Tag) error {
	if tag != nil {
		invoice.Tags = append(invoice.Tags, tag)
		return nil
	}
	return exception.Invoice_RequiredObject_Tag
}

func (invoice *Invoice) RemoveTagByName(name string) *Invoice {
	removeIndex := -1
	for tagIndex, tagItem := range invoice.Tags {
		if tagItem.Name == strings.TrimSpace(strings.Title(name)) {
			removeIndex = tagIndex
			break
		}
	}
	invoice.Tags = append(invoice.Tags[:removeIndex], invoice.Tags[removeIndex+1:]...)
	return invoice
}

func (invoice *Invoice) RemoveTag(tag *Tag) *Invoice {
	removeIndex := -1
	for tagIndex, tagItem := range invoice.Tags {
		if tagItem.Name == strings.TrimSpace(strings.Title(tag.Name)) {
			removeIndex = tagIndex
			break
		}
	}
	invoice.Tags = append(invoice.Tags[:removeIndex], invoice.Tags[removeIndex+1:]...)
	return invoice
}
