package boomware

type Boomware interface {
	// Sending short text message
	SMS(r *SMSRequest) *Response

	// Number verification api
	Verify(r *VerifyRequest) *Response
	VerifyCheck(r *VerifyCheckRequest) *VerifyCheckResponse
	VerifyInfo(r *VerifyInfoRequest) *VerifyInfoResponse
}
