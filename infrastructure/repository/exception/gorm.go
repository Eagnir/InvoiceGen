package exception

import "errors"

var GORM_UnknownError = errors.New("Unknown error occured")

var GORM_ContextDoesNotExist = errors.New("Context does not exist. Please use 'OpenContext' to initilize a context")
