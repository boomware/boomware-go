package boomware

import (
	"net/http"
	"strings"
)

type boomware struct {
	HttpClient *http.Client
	endpoint   string
	user       string
	pass       string
}

const endpoint = "https://api.boomware.com"

func New(credentials string) Boomware {
	b := new(boomware)
	b.HttpClient = &http.Client{}
	b.endpoint = endpoint
	b.setCredential(credentials)
	return b
}

func (b *boomware) setCredential(credentials string) {
	uc := strings.Split(credentials, ":")
	if len(uc) != 2 {
		panic("boomware: invalid credentials")
	}
	b.user = uc[0]
	b.pass = uc[1]
}
