package boomware

type Boomware interface {
	// Sending short text message
	SMS(r *SMSRequest) *Response

	// Number verification api
	Verify(r *VerifyRequest) *Response
	VerifyCheck(r *VerifyCheckRequest) *VerifyCheckResponse
	VerifyInfo(r *VerifyInfoRequest) *VerifyInfoResponse
}

type Request struct {
	// phone number in e164 format *required
	Number string `json:"number"`

	// callbackUrl http callback for status of the request *optional
	CallbackURL string `json:"callbackUrl,omitempty"`
}

type Response struct {
	// Error of the request nil if the request was succeeded
	Error *Error `json:"-"`
	// If request succeeded contains request id
	RequestID string `json:"requestId"`
}

type SMSRequest struct {
	Request
	From string `json:"from,omitempty"`
	Text string `json:"text"`
}

// Verification request

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

type VerifyCheckResponse struct {
	// Error of the request nil if the request was succeeded
	// Opposite details of the error
	Error *Error `json:"-"`

	// Verified number in e164 format
	Number string `json:"number"`
	// Method of the verification
	Method VerifyMethod `json:"method"`
	// Unix timestamp of the verification data
	VerifiedAt int64 `json:"verifiedAt"`
}

type VerifyCheckRequest struct {
	RequestID string `json:"requestId"`
	Code      string `json:"code"`
}

type VerifyStatus string

const (
	VerifyStatusProcessing VerifyStatus = "processing"
	VerifyStatusVerified   VerifyStatus = "verified"
)

type VerifyInfoResponse struct {
	// Error of the request nil if the request was succeeded
	// Opposite details of the error
	Error *Error `json:"-"`

	Status                  VerifyStatus `json:"status"`
	Number                  string       `json:"number"`
	Method                  VerifyMethod `json:"method"`
	VerifiedAt              int64        `json:"verifiedAt,omitempty"`
	VerificationRequestedAt int64        `json:"verificationRequestedAt"`
	CheckAttempts           int          `json:"checkAttempts"`
}

type VerifyInfoRequest struct {
	RequestID string `json:"requestId"`
}
