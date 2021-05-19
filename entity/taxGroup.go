package entity

import "InvoiceGen/entity/exception"

type TaxGroup struct {
	TaxGroupId int `gorm:"not null;primaryKey;"`

	Name      string `gorm:"unique;not null"`
	ShortName string `gorm:"unique;not null"`

	Invoices []*Invoice //`gorm:"foreignKey:TaxGroupId;"`
	Taxes    []*Tax     //`gorm:"foreignKey:TaxGroupId;"`

	DefaultStruct
}

func (obj *TaxGroup) Validate() error {
	if obj.Name == "" {
		return exception.TaxGroup_RequiredField_Name
	}
	if obj.ShortName == "" {
		return exception.TaxGroup_RequiredField_ShortName
	}
	return nil
}

func NewTaxGroup(name, shortName string) (*TaxGroup, error) {
	if name == "" {
		return nil, exception.TaxGroup_RequiredField_Name
	}
	if shortName == "" {
		return nil, exception.TaxGroup_RequiredField_ShortName
	}
	u := &TaxGroup{
		Name:      name,
		ShortName: shortName,
	}
	return u, nil
}

func (taxGroup *TaxGroup) AddTax(tax *Tax) (*TaxGroup, error) {
	if tax != nil {
		taxGroup.Taxes = append(tax.TaxGroup.Taxes, tax)
		return taxGroup, nil
	}
	return nil, exception.TaxGroup_RequiredObject_Tax
}
