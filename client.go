package boomware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type boomware struct {
	httpClient *http.Client
	endpoint   string
	user       string
	pass       string
}

const endpoint = "https://api.boomware.com"

func NewHttp(client *http.Client, credentials string) Boomware {
	b := new(boomware)
	b.httpClient = client
	b.endpoint = endpoint
	b.setCredential(credentials)
	return b
}

func New(credentials string) Boomware {
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}
	return NewHttp(httpClient, credentials)
}

func (b *boomware) setCredential(credentials string) {
	uc := strings.Split(credentials, ":")
	if len(uc) != 2 {
		panic("boomware: invalid credentials. Please make sure that you provide correct access token.")
	}
	b.user = uc[0]
	b.pass = uc[1]
}

func (b *boomware) SMS(r *SMSRequest) *Response {
	response := new(Response)
	err := b.request(http.MethodPost, "/v1/sms", r, response)
	if err != nil {
		response.Error = err
	}
	return response
}

func (b *boomware) CallsFlash(r *CallsFlashRequest) *Response {
	response := new(Response)
	err := b.request(http.MethodPost, "/v1/calls/flash", r, response)
	if err != nil {
		response.Error = err
	}
	return response
}

func (b *boomware) Verify(r *VerifyRequest) *Response {
	response := new(Response)
	err := b.request(http.MethodPost, "/v1/verify", r, response)
	if err != nil {
		response.Error = err
	}
	return response
}

func (b *boomware) VerifyCheck(r *VerifyCheckRequest) *VerifyCheckResponse {
	response := new(VerifyCheckResponse)
	err := b.request(http.MethodPost, "/v1/verify/check", r, response)
	if err != nil {
		response.Error = err
	}
	response.ID = r.ID
	return response
}

func (b *boomware) VerifyInfo(r *VerifyInfoRequest) *VerifyInfoResponse {
	response := new(VerifyInfoResponse)
	err := b.request(http.MethodPost, "/v1/verify/info", r, response)
	if err != nil {
		response.Error = err
	}
	response.ID = r.ID
	return response
}

// /v1/insight

func (b *boomware) Insight(r *InsightRequest) *InsightResponse {
	response := new(InsightResponse)
	path := fmt.Sprintf("/v1/insight/hlr?number=%s", url.QueryEscape(r.Number))
	err := b.request(http.MethodGet, path, nil, response)
	if err != nil {
		response.Error = err
	}
	return response
}

func (b *boomware) MessagingPush(r *MessagingPushRequest) *Response {
	response := new(Response)
	err := b.request(http.MethodPost, "/v1/messaging/push", r, response)
	if err != nil {
		response.Error = err
	}
	return response
}

// Info

func (b *boomware) RequestInfo(requestID string) *InfoResponse {
	response := new(InfoResponse)
	path := fmt.Sprintf("/v1/requests/%s", requestID)
	err := b.request(http.MethodGet, path, nil, response)
	if err != nil {
		response.Error = err
	}
	return response
}

// private

func (b *boomware) request(method, path string, request interface{}, response interface{}) Error {

	urlString := b.endpoint + path

	var buf io.Reader

	// Marshal only requests where data available
	if request != nil {
		jsonBytes, err := json.Marshal(request)
		if err != nil {
			return NewError(MarshalRequestErrorCode, fmt.Sprintf("failed to marshal request:%s error:%s", path, err.Error()))
		}
		buf = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequest(method, urlString, buf)
	if err != nil {
		return NewError(CreateRequestErrorCode, fmt.Sprintf("failed to create http request:%s error:%s", path, err.Error()))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Boomware/1.0")
	req.SetBasicAuth(b.user, b.pass)

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return NewError(DoRequestErrorCode, fmt.Sprintf("http request:%s error %s", path, err.Error()))
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NewError(ReadBodyErrorCode, fmt.Sprintf("failed to read body request:%s error %s", path, err.Error()))
	}

	if b.isCodeSucceeded(resp.StatusCode) == false {
		apiErr := NewError(UnknownErrorCode, "unknown error")
		err = json.Unmarshal(body, apiErr)
		if err != nil {
			return NewError(UnmarshalErrorCode, fmt.Sprintf("failed to unmarshal request:%s error:%s", path, err.Error()))
		}
		return apiErr
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return NewError(UnmarshalErrorCode, fmt.Sprintf("failed to unmarshal request:%s error:%s", path, err.Error()))
	}

	return nil
}

func (boomware) isCodeSucceeded(statusCode int) bool {
	switch statusCode {
	case http.StatusOK, http.StatusCreated:
		return true
	default:
		return false
	}
}
