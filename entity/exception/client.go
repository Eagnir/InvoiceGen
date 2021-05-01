package exception

import "errors"

var Client_RecordNotFound = errors.New("Client not found")

var Client_RequiredField_Name = errors.New("Please provide a value for the name field")
var Client_RequiredField_Address = errors.New("Please provide a value for the address field")
var Client_RequiredField_ContactNumber = errors.New("Please provide a value for the contact number field")

var Client_PrimeryKeyNotZero = errors.New("Client's primary key is not zero")
