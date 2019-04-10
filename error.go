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
