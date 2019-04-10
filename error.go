package boomware

import "fmt"

type Error struct {
	Code   int    `json:"errorCode"`
	Reason string `json:"errorReason"`
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf("boomware: error:%d reason:%s", e.Code, e.Reason)
}

func NewError(code int, reason string) *Error {
	return &Error{Code: code, Reason: reason}
}

const (
	// Api errors
	// For more info please visit https://boomware.com/docs/en/#Errors
	//
	InternalServerErrorCode    = 99
	AuthRequiredErrorCode      = 1
	InactiveAccountErrorCode   = 2
	TooManyRequestsErrorCode   = 4
	InactiveTokenErrorCode     = 5
	InsufficientFundsErrorCode = 6
	InvalidRequestErrorCode    = 10
	InvalidNumberErrorCode     = 11
	SenderNotAllowedErrorCode  = 16

	// System errors
	UnknownErrorCode        = 1000
	MarshalRequestErrorCode = 1001
	CreateRequestErrorCode  = 1002
	DoRequestErrorCode      = 1003
	ReadBodyErrorCode       = 1004
	UnmarshalErrorCode      = 1005
)
