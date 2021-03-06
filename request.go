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

// /v1/calls/flash

type CallsFlashRequest struct {
	Request
	// Optional param e164 number
	From string `json:"from,omitempty"`
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

// /v1/messaging/push

type MessagingPushTarget string

const (
	MessagingPushTargetAll  MessagingPushTarget = "all"
	MessagingPushTargetLast MessagingPushTarget = "last"
)

type MessagingPushPriority string

const (
	MessagingPushPriorityHigh   MessagingPushPriority = "high"
	MessagingPushPriorityNormal MessagingPushPriority = "normal"
)

type MessagingPushAndroid struct {
	Data         map[string]interface{} `json:"data,omitempty"`
	Notification *AndroidNotification   `json:"notification,omitempty"`
}

type MessagingPushRequest struct {
	Request

	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`

	Target      MessagingPushTarget   `json:"target,omitempty"`
	TimeToLive  int                   `json:"ttl,omitempty"`
	Priority    MessagingPushPriority `json:"priority,omitempty"`
	CollapseKey string                `json:"collapseKey,omitempty"`

	Android *MessagingPushAndroid  `json:"android,omitempty"`
	IOS     map[string]interface{} `json:"ios,omitempty"`
}

type AndroidNotification struct {
	Title            string `json:"title,omitempty"`
	Body             string `json:"body,omitempty"`
	Icon             string `json:"icon,omitempty"`
	Sound            string `json:"sound,omitempty"`
	Badge            string `json:"badge,omitempty"`
	Tag              string `json:"tag,omitempty"`
	Color            string `json:"color,omitempty"`
	ClickAction      string `json:"click_action,omitempty"`
	BodyLocKey       string `json:"body_loc_key,omitempty"`
	BodyLocArgs      string `json:"body_loc_args,omitempty"`
	TitleLocKey      string `json:"title_loc_key,omitempty"`
	TitleLocArgs     string `json:"title_loc_args,omitempty"`
	AndroidChannelID string `json:"android_channel_id,omitempty"`
}
