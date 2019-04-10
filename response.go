package boomware

type Response struct {
	// Error of the request nil if the request was succeeded
	Error *Error `json:"-"`
	// If request succeeded contains request id
	ID string `json:"requestId"`
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
