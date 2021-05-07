package setting

import (
	"runtime"
	"strings"
)

const VersionMajor int = 1
const VersionMinor int = 0

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
