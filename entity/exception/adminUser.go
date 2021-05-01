package exception

import "errors"

var AdminUser_RecordNotFound = errors.New("Admin user not found")
var AdminUser_RequiredField_Name = errors.New("Please provide a value for the name field")
var AdminUser_RequiredField_Email = errors.New("Please provide a value for the email field")
var AdminUser_RequiredField_Password = errors.New("Please provide a value for the password field")

var AdminUser_PrimeryKeyNotZero = errors.New("Admin user's primary key is not zero")
