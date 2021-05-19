package entity

import (
	"InvoiceGen/entity/exception"
)

type AdminUser struct {
	AdminUserId int `gorm:"not null;primaryKey"`

	Name     string `gorm:"not null;default:'Admin User'"`
	Email    string `gorm:"not null;unique;"`
	Password string `gorm:"not null;"`

	AuthToken *UUID `gorm:"type:uuid;"`

	CompanyId int `gorm:"default:null"`

	Invoices []*Invoice
	Company  *Company

	DefaultStruct
}

func (obj *AdminUser) Validate() error {
	if obj.Name == "" {
		return exception.AdminUser_RequiredField_Name
	}
	if obj.Email == "" {
		return exception.AdminUser_RequiredField_Email
	}
	if obj.Password == "" {
		return exception.AdminUser_RequiredField_Password
	}
	if obj.Company == nil {
		return exception.AdminUser_RequiredField_Company
	}

	return nil
}

func NewAdminUser(name, email, password string, company *Company) (*AdminUser, error) {
	if name == "" {
		return nil, exception.AdminUser_RequiredField_Name
	}
	if email == "" {
		return nil, exception.AdminUser_RequiredField_Email
	}
	if password == "" {
		return nil, exception.AdminUser_RequiredField_Password
	}
	if company == nil {
		return nil, exception.AdminUser_RequiredField_Company
	}
	u := &AdminUser{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	}
	return u, nil
}

func (obj *AdminUser) SwitchCompany(company *Company) error {
	obj.Company = company
	obj.CompanyId = company.CompanyId
	return nil
}
