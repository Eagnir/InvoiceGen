package entity

import "InvoiceGen/entity/exception"

type Tax struct {
	TaxId int `gorm:"not null;primaryKey"`

	Name       string  `gorm:"unique;not null"`
	ShortName  string  `gorm:"unique;not null"`
	Percentage float32 `gorm:"not null"`

	TaxGroupId int `gorm:"not null"`

	TaxGroup *TaxGroup

	DefaultStruct
}

func (obj *Tax) Validate() error {
	if obj.Name == "" {
		return exception.Tax_RequiredField_Name
	}
	if obj.ShortName == "" {
		return exception.Tax_RequiredField_ShortName
	}
	if obj.Percentage <= 0 {
		return exception.Tax_RequiredField_Percentage
	}
	if obj.TaxGroup == nil {
		return exception.Tax_RequiredField_TaxGroup
	}

	return nil
}

func NewTax(name, shortName string, percentage float32, taxGroup *TaxGroup) (*Tax, error) {
	if name == "" {
		return nil, exception.Tax_RequiredField_Name
	}
	if shortName == "" {
		return nil, exception.Tax_RequiredField_ShortName
	}
	if percentage <= 0 {
		return nil, exception.Tax_RequiredField_Percentage
	}
	if taxGroup == nil {
		return nil, exception.Tax_RequiredField_TaxGroup
	}
	u := &Tax{
		Name:       name,
		ShortName:  shortName,
		Percentage: percentage,
		TaxGroup:   taxGroup,
	}
	return u, nil
}

func (tax *Tax) SwitchTaxGroup(taxGroup *TaxGroup) error {
	tax.TaxGroup = taxGroup
	return nil
}

/* func (tax *Tax) SetNewTaxGroup(name, shortName string) error {
	taxGroup, ex := NewTaxGroup(name, shortName)
	if ex == nil {
		tax.SetTaxGroup(taxGroup)
	} else {
		return ex
	}
	return nil
}
*/
