package exception

import "errors"

var Client_RecordNotFound = errors.New("Client not found")

var Client_RequiredField_Name = errors.New("Please provide a value for the name field")
var Client_RequiredField_Address = errors.New("Please provide a value for the address field")
var Client_RequiredField_ContactNumber = errors.New("Please provide a value for the contact number field")
var Client_RequiredField_Country = errors.New("Please provide a value for the country field")
var Client_RequiredField_Currency = errors.New("Please provide a default currency for this client")
var Client_RequiredField_Company = errors.New("Please provide a company for this client")

var Client_PrimeryKeyNotZero = errors.New("Client's primary key is not zero")
