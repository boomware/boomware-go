package boomware

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Error(t *testing.T) {
	e := new(Error)
	e.Code = 99
	e.Reason = "Internal server error"
	assert.Error(t, e)
}

func TestError_String(t *testing.T) {
	e := new(Error)
	e.Code = 99
	e.Reason = "Internal server error"
	t.Log(e)
}

func TestNewError(t *testing.T) {
	e := NewError(UnknownErrorCode, "unknown")
	assert.NotNil(t, e)
	assert.Equal(t, UnknownErrorCode, e.Code)
	assert.Equal(t, "unknown", e.Reason)

	err := json.Unmarshal([]byte(`{"errorCode":99, "errorReason": "InternalServer"}`), e)
	assert.NoError(t, err)

	assert.Equal(t, InternalServerErrorCode, e.Code)
	assert.Equal(t, "InternalServer", e.Reason)
}
