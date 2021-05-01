package entity

import "InvoiceGen/entity/exception"

type Company struct {
	CompanyId int `gorm:"not null;primaryKey"`

	Name          string `gorm:"unique;not null"`
	Address       string `gorm:"not null;"`
	Email         string `gorm:"default:null"`
	ContactNumber string `gorm:"not null"`
	GSTNumber     string `gorm:"unique"`

	Invoices []*Invoice `gorm:"references:CompanyId"`
	Clients  []*Client  `gorm:"references:CompanyId"`

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

	return nil
}

func NewCompany(name, address, email, contactNumber, gstNumber string) (*Company, error) {
	if name == "" {
		return nil, exception.Company_RequiredField_Name
	}
	if address == "" {
		return nil, exception.Company_RequiredField_Address
	}
	if contactNumber == "" {
		return nil, exception.Company_RequiredField_ContactNumber
	}

	u := &Company{
		Name:          name,
		Address:       address,
		Email:         email,
		ContactNumber: contactNumber,
		GSTNumber:     gstNumber,
	}
	return u, nil
}
