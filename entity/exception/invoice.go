package exception

import "errors"

var Invoice_RecordNotFound = errors.New("Invoice not found")

var Invoice_RequiredField_ClientId = errors.New("Please provide a client for this invoice")
var Invoice_RequiredField_CompanyId = errors.New("Please provide a company for this invoice")

var Invoice_RequiredObject_Tag = errors.New("Please provide a tag object to be added")

var Invoice_PrimeryKeyNotZero = errors.New("Invoice's primary key is not zero")
