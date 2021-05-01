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

	Invoices []*Invoice `gorm:"references:AdminUserId"`

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

	return nil
}

func NewAdminUser(name, email, password string) (*AdminUser, error) {
	if name == "" {
		return nil, exception.AdminUser_RequiredField_Name
	}
	if email == "" {
		return nil, exception.AdminUser_RequiredField_Email
	}
	if password == "" {
		return nil, exception.AdminUser_RequiredField_Password
	}
	u := &AdminUser{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return u, nil
}
