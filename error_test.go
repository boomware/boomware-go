package boomware

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Error(t *testing.T) {
	e := NewError(99, "Internal server error")
	assert.Error(t, e)
}

func TestError_String(t *testing.T) {
	e := NewError(99, "Internal server error")
	assert.Equal(t, "boomware: error:99 reason:Internal server error", e.String())
	t.Log(e)
}

func TestNewError(t *testing.T) {
	e := NewError(UnknownErrorCode, "unknown")
	assert.NotNil(t, e)
	assert.Equal(t, UnknownErrorCode, e.GetCode())
	assert.Equal(t, "unknown", e.GetReason())

	err := json.Unmarshal([]byte(`{"errorCode":99, "errorReason": "InternalServer"}`), e)
	assert.NoError(t, err)

	assert.Equal(t, InternalServerErrorCode, e.GetCode())
	assert.Equal(t, "InternalServer", e.GetReason())
}
