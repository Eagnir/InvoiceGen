package exception

import "errors"

var Currency_RecordNotFound = errors.New("Admin user not found")
var Currency_RequiredField_ShortName = errors.New("Please provide a value for the short name field")

var Currency_PrimeryKeyNotZero = errors.New("Currency's primary key is not zero")
