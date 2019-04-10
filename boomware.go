package boomware

type Boomware interface {
	// Sending short text message
	// For more info https://boomware.com/docs/en/#SMS
	SMS(r *SMSRequest) *Response

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
}
