package exception

import "errors"

var Tax_RecordNotFound = errors.New("Invoice item not found")

var Tax_RequiredField_Name = errors.New("Please provide a name for this tax")
var Tax_RequiredField_ShortName = errors.New("Please provide a short name for this tax")
var Tax_RequiredField_Percentage = errors.New("Please provide a positive percentage to be tax")
var Tax_RequiredField_TaxGroup = errors.New("Please provide a tax group for this tax")

var Tax_PrimeryKeyNotZero = errors.New("Tax's primary key is not zero")
