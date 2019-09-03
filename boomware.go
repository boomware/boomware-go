package boomware

type Boomware interface {
	// Sending short text message
	// For more info https://boomware.com/docs/en/#SMS
	SMS(r *SMSRequest) *Response

	// Flash call
	CallsFlash(r *CallsFlashRequest) *Response

	// Number verification api
	// More info https://boomware.com/docs/en/#Verification
	Verify(r *VerifyRequest) *Response
	VerifyCheck(r *VerifyCheckRequest) *VerifyCheckResponse
	VerifyInfo(r *VerifyInfoRequest) *VerifyInfoResponse

	// Insight
	// More info https://boomware.com/docs/en/#Insight
	Insight(r *InsightRequest) *InsightResponse

	// Messaging
	// Push
	MessagingPush(r *MessagingPushRequest) *Response

	// Info
	// For more info https://boomware.com/docs/en/#RequestInfo
	RequestInfo(requestID string) *InfoResponse
}
