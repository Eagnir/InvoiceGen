package exception

import "errors"

var AuthHeader_InvalidAuthenticationHeaders = errors.New("Invalid authentication headers")
var AuthHeader_MissingAuthenticationToken = errors.New("Missing or invalid authentication token")
var AuthHeader_MissingEmail = errors.New("Missing or invalid email address")
