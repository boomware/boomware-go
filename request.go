package boomware

type Request struct {
	// phone number in e164 format *required
	Number string `json:"number"`

	// callbackUrl http callback for status of the request *optional
	CallbackURL string `json:"callbackUrl,omitempty"`
}

// /v1/sms

type SMSRequest struct {
	Request
	From string `json:"from,omitempty"`
	Text string `json:"text"`
}

// /v1/verify

type VerifyMethod string

const (
	VerifyMethodSMS   VerifyMethod = "sms"
	VerifyMethodCall  VerifyMethod = "call"
	VerifyMethodVoice VerifyMethod = "voice"
)

type VerifyRequest struct {
	Number               string       `json:"number"`
	CodeLength           int          `json:"codeLength,omitempty"`
	Language             string       `json:"language,omitempty"`
	Method               VerifyMethod `json:"method,omitempty"`
	CodeExpiry           int          `json:"codeExpiry,omitempty"`
	CheckAttemptsAllowed int          `json:"checkAttemptsAllowed,omitempty"`
}

type VerifyCheckRequest struct {
	ID   string `json:"requestId"`
	Code string `json:"code"`
}

type VerifyInfoRequest struct {
	ID string `json:"requestId"`
}

// /v1/insight

type InsightRequest struct {
	// phone number in e164 format *required
	Number string
}
