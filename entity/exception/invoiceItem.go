package exception

import "errors"

var InvoiceItem_RecordNotFound = errors.New("Invoice item not found")

var InvoiceItem_RequiredField_Title = errors.New("Please provide a title for this line item")
var InvoiceItem_RequiredField_Rate = errors.New("Please provide a rate for this line item")
var InvoiceItem_RequiredField_Quantity = errors.New("Please provide a quantity for this line item")

var InvoiceItem_PrimeryKeyNotZero = errors.New("Invoice Item's primary key is not zero")
