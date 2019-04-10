package boomware

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	b := New("user:pass")
	assert.NotNil(t, b)
	assert.Equal(t, "user", b.(*boomware).user)
	assert.Equal(t, "pass", b.(*boomware).pass)
	assert.Equal(t, endpoint, b.(*boomware).endpoint)
	assert.NotNil(t, b.(*boomware).HttpClient)
}

func Test_setCredentials(t *testing.T) {
	b := &boomware{}

	assert.Panics(t, func() {
		b.setCredential("test")
	})

	assert.NotPanics(t, func() {
		b.setCredential("user:pass")
	})

	assert.Equal(t, "user", b.user)
	assert.Equal(t, "pass", b.pass)
}
