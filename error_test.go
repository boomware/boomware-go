package boomware

import (
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
