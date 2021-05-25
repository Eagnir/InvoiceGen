package dto

type TaxGroup struct {
	TaxGroupId int

	Name      string
	ShortName string

	Taxes []*Tax
}
