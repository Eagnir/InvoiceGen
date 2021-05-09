package setting

import (
	"runtime"
	"strings"
)

const VersionMajor int = 1
const VersionMinor int = 0

const APIResetToken = "9f6a1c3e-7fe0-41d5-8ee3-52e0610f55cb"

const APITokenKey = "token"
const APIUserEmailKey = "email"

//API Response Status
const (
	StatusSuccess = iota
	StatusWarning
	StatusFailure
	StatusFatalError
)

//API Response Status Text
const (
	StatusSuccessText    = "Success"
	StatusWarningText    = "Warning"
	StatusFailureText    = "Failure"
	StatusFatalErrorText = "Fatal"
)

var DefaultPageSize int = 10

func GetCallerFunctionName() string {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		ar := strings.Split(details.Name(), ".")
		funcName := ar[len(ar)-1]
		return funcName
	}
	return "Unkown"
}
