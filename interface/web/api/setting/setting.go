package setting

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
