package exception

import "errors"

var Company_RecordNotFound = errors.New("Company not found")

var Company_RequiredField_Name = errors.New("Please provide a value for the name field")
var Company_RequiredField_Address = errors.New("Please provide a value for the address field")
var Company_RequiredField_ContactNumber = errors.New("Please provide a value for the contact number field")
var Company_RequiredField_Currency = errors.New("Please provide a default currency for this company")

var Company_PrimeryKeyNotZero = errors.New("Company's primary key is not zero")
