package domain

import "fmt"

type GPErrorCode int

const (
	ErrCodeParametersNotCorrect GPErrorCode = 400000
	ErrCodeUserNotFound         GPErrorCode = 400001
	ErrCodeInvalidPassword      GPErrorCode = 400002
	ErrCodeUserEmailExists      GPErrorCode = 400003

	ErrCodeUnauthorized GPErrorCode = 401000

	ErrCodeInternalServer GPErrorCode = 500000

	ErrCodeDatabaseError     GPErrorCode = 501000
	ErrCodeModelConvertError GPErrorCode = 501001
)

func (code GPErrorCode) Message() string {
	switch code {
	case ErrCodeParametersNotCorrect:
		return "Parameters Not Correct"

	case ErrCodeUserNotFound:
		return "User Not Found"

	case ErrCodeInvalidPassword:
		return "Invalid Password"

	case ErrCodeInternalServer:
		return "Internal Server Error"

	default:
		return fmt.Sprintf("Unknown Error Code: %d", code)
	}
}
