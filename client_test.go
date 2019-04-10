package boomware

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func TestBoomware_setCredentials(t *testing.T) {
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

func TestBoomware_request(t *testing.T) {

	// Set fake server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "PATCH", r.Method)
		user, pass, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "user", user)
		assert.Equal(t, "pass", pass)
		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, "{\"number\":\"+1800000000\"}", string(body))
		_, _ = w.Write([]byte(`{"requestId":"aa7d0899-38bd-4fba-a270-a76f4d7e8b5d"}`))
	}))
	defer server.Close()

	b := New("user:pass").(*boomware)
	// set url of a fake server
	b.endpoint = server.URL

	// Make a request data struct
	request := &struct {
		Number string `json:"number"`
	}{
		Number: "+1800000000",
	}

	// Make a response data struct
	response := &struct {
		ID string `json:"requestId"`
	}{}

	apiErr := b.request("PATCH", "/example/api", request, response)
	assert.NoError(t, apiErr)
	assert.Equal(t, "aa7d0899-38bd-4fba-a270-a76f4d7e8b5d", response.ID)
}
