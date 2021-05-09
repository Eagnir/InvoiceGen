package message

var AuthError_DTOConversion = "Error occured while trying to serve admin user details"
var AuthError_Unknown = "Error occured, please try again later"

// CatchAll Handler Messages
var AuthError_CatchAll = "Cannot find the API call you are trying to reach"

// Credential Handler Messages
//    Error Messages
var AuthError_Credential = "Error occured while trying to authenticate credential, please try again later"
var AuthError_InvalidToken = "Invalid authentication token"
var AuthError_InvalidCredential = "Invalid credential"
var AuthError_GeneratingToken = "Error occured while trying to generate token"

//    Success Messages
var AuthSuccess_ValidCredential = "Authentication successful"

// ResetPassword Handler Messages
//    Error Messages
var AuthError_MissingPassword = "Please provide a new password"

//    Success Messages
var AuthSuccess_PasswordResetSuccessful = "Password reset successful"

// Invalidate Handler Messages
//    Success Messages
var AuthSuccess_InvalidateSuccessful = "Logout successful"
