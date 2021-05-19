package entity

import "InvoiceGen/entity/exception"

type Client struct {
	ClientId int `gorm:"not null;primaryKey"`

	Name          string `gorm:"unique;not null"`
	Address       string `gorm:"not null;"`
	Email         string `gorm:"default:null"`
	ContactNumber string `gorm:"not null;"`
	GSTNumber     string `gorm:"unique"`

	CompanyId         int `gorm:"not null"`
	DefaultCurrencyId int `gorm:"not null"`

	Company  *Company
	Invoices []*Invoice
	Tags     []*Tag `gorm:"many2many:client_tags;"`

	DefaultCurrency *Currency

	DefaultStruct
}

func (obj *Client) Validate() error {
	if obj.Name == "" {
		return exception.Client_RequiredField_Name
	}
	if obj.Address == "" {
		return exception.Client_RequiredField_Address
	}
	if obj.ContactNumber == "" {
		return exception.Client_RequiredField_ContactNumber
	}
	if obj.DefaultCurrency == nil {
		return exception.Client_RequiredField_Currency
	}
	if obj.Company == nil {
		return exception.Client_RequiredField_Company
	}

	return nil
}

func NewClient(name, address, email, contactNumber, gstNumber string, currency *Currency, company *Company) (*Client, error) {
	if name == "" {
		return nil, exception.Client_RequiredField_Name
	}
	if address == "" {
		return nil, exception.Client_RequiredField_Address
	}
	if contactNumber == "" {
		return nil, exception.Client_RequiredField_ContactNumber
	}
	if currency == nil {
		return nil, exception.Client_RequiredField_Currency
	}
	if company == nil {
		return nil, exception.Client_RequiredField_Company
	}

	u := &Client{
		Name:            name,
		Address:         address,
		Email:           email,
		ContactNumber:   contactNumber,
		GSTNumber:       gstNumber,
		Company:         company,
		DefaultCurrency: currency,
	}
	return u, nil
}

func (obj *Client) SwitchCurrency(currency *Currency) error {
	obj.DefaultCurrency = currency
	//obj.CurrencyId = currency.CurrencyId
	return nil
}

func (obj *Client) SwitchCompany(company *Company) error {
	obj.Company = company
	//obj.CurrencyId = currency.CurrencyId
	return nil
}
