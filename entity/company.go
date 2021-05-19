package entity

import "InvoiceGen/entity/exception"

type Company struct {
	CompanyId int `gorm:"not null;primaryKey"`

	DefaultCurrencyId int `gorm:"default:null;"`

	Name          string `gorm:"unique;not null"`
	Address       string `gorm:"not null;"`
	Email         string `gorm:"default:null"`
	ContactNumber string `gorm:"not null"`
	GSTNumber     string `gorm:"unique"`

	Invoices   []*Invoice
	Clients    []*Client
	AdminUsers []*AdminUser

	DefaultCurrency *Currency //default = foreignKey:DefaultCurrencyId = Company.<FieldName>Id -> Currency.PrimaryKey

	DefaultStruct
}

func (obj *Company) Validate() error {
	if obj.Name == "" {
		return exception.Company_RequiredField_Name
	}
	if obj.Address == "" {
		return exception.Company_RequiredField_Address
	}
	if obj.ContactNumber == "" {
		return exception.Company_RequiredField_ContactNumber
	}
	if obj.DefaultCurrency == nil {
		return exception.Company_RequiredField_Currency
	}
	return nil
}

func NewCompany(name, address, email, contactNumber, gstNumber string, currency *Currency) (*Company, error) {
	if name == "" {
		return nil, exception.Company_RequiredField_Name
	}
	if address == "" {
		return nil, exception.Company_RequiredField_Address
	}
	if contactNumber == "" {
		return nil, exception.Company_RequiredField_ContactNumber
	}
	if currency == nil {
		return nil, exception.Company_RequiredField_Currency
	}

	u := &Company{
		Name:            name,
		Address:         address,
		Email:           email,
		ContactNumber:   contactNumber,
		GSTNumber:       gstNumber,
		DefaultCurrency: currency,
	}
	return u, nil
}

func (obj *Company) SwitchCurrency(currency *Currency) error {
	obj.DefaultCurrency = currency
	return nil
}
