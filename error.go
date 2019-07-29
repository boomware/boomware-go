package boomware

import "fmt"

type Error interface {
	error
	fmt.Stringer
	GetCode() ErrorCode
	GetReason() string
}

type boomwareError struct {
	Code   ErrorCode `json:"errorCode"`
	Reason string    `json:"errorReason"`
}

func (e *boomwareError) GetCode() ErrorCode {
	return e.Code
}

func (e *boomwareError) GetReason() string {
	return e.Reason
}

func (e *boomwareError) Error() string {
	return e.String()
}

func (e *boomwareError) String() string {
	return fmt.Sprintf("boomware: error:%d reason:%s", e.Code, e.Reason)
}

func NewError(code ErrorCode, reason string) Error {
	return &boomwareError{Code: code, Reason: reason}
}

type ErrorCode int

const (
	// Api errors
	// For more info please visit https://boomware.com/docs/en/#Errors
	//
	InternalServerErrorCode    ErrorCode = 99
	AuthRequiredErrorCode      ErrorCode = 1
	InactiveAccountErrorCode   ErrorCode = 2
	TooManyRequestsErrorCode   ErrorCode = 4
	InactiveTokenErrorCode     ErrorCode = 5
	InsufficientFundsErrorCode ErrorCode = 6
	InvalidRequestErrorCode    ErrorCode = 10
	InvalidNumberErrorCode     ErrorCode = 11
	SenderNotAllowedErrorCode  ErrorCode = 16
	RequestExpiredErrorCode    ErrorCode = 22

	// System errors
	UnknownErrorCode        ErrorCode = 1000
	MarshalRequestErrorCode ErrorCode = 1001
	CreateRequestErrorCode  ErrorCode = 1002
	DoRequestErrorCode      ErrorCode = 1003
	ReadBodyErrorCode       ErrorCode = 1004
	UnmarshalErrorCode      ErrorCode = 1005
)
