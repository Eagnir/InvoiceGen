package entity

import "InvoiceGen/entity/exception"

type Client struct {
	ClientId int `gorm:"not null;primaryKey"`

	Name          string `gorm:"unique;not null"`
	Address       string `gorm:"not null;"`
	Email         string `gorm:"default:null"`
	ContactNumber string `gorm:"not null;"`
	GSTNumber     string `gorm:"unique"`

	CompanyId int `gorm:"not null"`

	Company  *Company   `gorm:"references:CompanyId"`
	Invoices []*Invoice `gorm:"references:ClientId"`
	Tags     []*Tag     `gorm:"many2many:client_tags;"`

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

	return nil
}

func NewClient(name, address, email, contactNumber, gstNumber string) (*Client, error) {
	if name == "" {
		return nil, exception.Client_RequiredField_Name
	}
	if address == "" {
		return nil, exception.Client_RequiredField_Address
	}
	if contactNumber == "" {
		return nil, exception.Client_RequiredField_ContactNumber
	}

	u := &Client{
		Name:          name,
		Address:       address,
		Email:         email,
		ContactNumber: contactNumber,
		GSTNumber:     gstNumber,
	}
	return u, nil
}
