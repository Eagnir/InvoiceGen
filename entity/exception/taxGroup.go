package exception

import "errors"

var TaxGroup_RecordNotFound = errors.New("Tax Group not found")

var TaxGroup_RequiredField_Name = errors.New("Please provide a name for this tax group")
var TaxGroup_RequiredField_ShortName = errors.New("Please provide a short name for this tax group")

var TaxGroup_RequiredObject_Tax = errors.New("Please provide a Tax object to add into this tax group")

var TaxGroup_PrimeryKeyNotZero = errors.New("Tax Group's primary key is not zero")
