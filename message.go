package fcm

import (
	"errors"
	"go/types"
)

var (
	// ErrInvalidMessage occurs if push notitication message is nil.
	ErrInvalidMessage = errors.New("message is invalid")

	// ErrInvalidTarget occurs if message topic is empty.
	ErrInvalidTarget = errors.New("must be filled only topic | token | condition")
)

type Message struct {
	Name          string                 `json:"name,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty"`
	Notification  *Notification          `json:"notification,omitempty"`
	AndroidConfig *AndroidConfig         `json:"android,omitempty"`
	WebpushConfig *WebpushConfig         `json:"webpush,omitempty"`
	ApnsConfig    *ApnsConfig            `json:"apns,omitempty"`
	FcmOptions    *FcmOptions            `json:"fcm_options,omitempty"`
	Token         string                 `json:"token,omitempty"`
	Topic         string                 `json:"topic,omitempty"`
	Condition     string                 `json:"condition,omitempty"`
}

func (msg *Message) Validate() error {
	if msg == nil {
		return ErrInvalidMessage
	}
	if msg.Token != "" && msg.Topic != "" {
		return ErrInvalidTarget
	}
	if msg.Token != "" && msg.Condition != "" {
		return ErrInvalidTarget
	}
	if msg.Topic != "" && msg.Condition != "" {
		return ErrInvalidTarget
	}
	return nil
}

type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Image string `json:"image,omitempty"`
}

type AndroidConfig struct {
	CollapseKey            string                 `json:"collapse_key,omitempty"`
	AndroidMessagePriority string                 `json:"priority,omitempty"`
	Ttl                    string                 `json:"ttl,omitempty"`
	RestrictedPackageName  string                 `json:"restricted_package_name,omitempty"`
	Data                   map[string]interface{} `json:"data,omitempty"`
	Notification           *AndroidNotification   `json:"notification,omitempty"`
	FcmOptions             *AndroidFcmOptions     `json:"fcm_options,omitempty"`
}

type WebpushConfig struct {
	Headers      map[string]interface{} `json:"headers,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Notification types.Object           `json:"Notification,omitempty"`
	FcmOptions   *WebpushFcmOptions     `json:"fcm_options,omitempty"`
}

type ApnsConfig struct {
	Headers    map[string]interface{} `json:"headers,omitempty"`
	Payload    types.Object           `json:"payload,omitempty"`
	FcmOptions *ApnsFcmOptions        `json:"fcm_options,omitempty"`
}

type FcmOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
}

type AndroidFcmOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
}

type WebpushFcmOptions struct {
	Link string `json:"link,omitempty"`
}

type ApnsFcmOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
	Link           string `json:"link,omitempty"`
}

type AndroidNotification struct {
	Title                 string                `json:"title,omitempty"`
	Body                  string                `json:"body,omitempty"`
	Icon                  string                `json:"icon,omitempty"`
	Color                 string                `json:"color,omitempty"`
	Sound                 string                `json:"sound,omitempty"`
	Tag                   string                `json:"tag,omitempty"`
	ClickAction           string                `json:"click_action,omitempty"`
	BodyLocKey            string                `json:"body_loc_key,omitempty"`
	BodyLocArgs           []string              `json:"body_loc_args,omitempty"`
	TitleLocKey           string                `json:"title_loc_key,omitempty"`
	TitleLocArgs          []string              `json:"title_loc_args,omitempty"`
	ChannelID             string                `json:"android_channel_id,omitempty"`
	Ticker                string                `json:"ticker,omitempty"`
	Sticky                bool                  `json:"sticky,omitempty"`
	EventTime             string                `json:"event_time,omitempty"`
	LocalOnly             string                `json:"local_only,omitempty"`
	NotificationPriority  *NotificationPriority `json:"notification_priority,omitempty"`
	DefaultSound          bool                  `json:"default_sound,omitempty"`
	DefaultVibrateTimings bool                  `json:"default_vibrate_timings,omitempty"`
	DefaultLightSettings  string                `json:"default_light_settings,omitempty"`
	VibrateTimings        []string              `json:"vibrate_timings,omitempty"`
	Visibility            *Visibility           `json:"visibility,omitempty"`
	NotificationCount     uint                  `json:"notification_count,omitempty"`
	LightSettings         *LightSettings        `json:"light_settings,omitempty"`
	Image                 string                `json:"image,omitempty"`
}

type LightSettings struct {
	Color            *Color `json:"color,omitempty"`
	LightOnDuration  string `json:"light_on_duration,omitempty"`
	LightOffDuration string `json:"light_off_duration,omitempty"`
}

type Color struct {
	Red   uint8 `json:"red,omitempty"`
	Green uint8 `json:"green,omitempty"`
	Blue  uint8 `json:"blue,omitempty"`
	Alpha uint8 `json:"alpha,omitempty"`
}

type MessagePriorityValue string

type NotificationPriorityValue string

type VisibilityValue string

const (
	MessagePriorityNormal MessagePriorityValue = "NORMAL"
	MessagePriorityHigh   MessagePriorityValue = "HIGH"

	NotificationPriorityUnspecified NotificationPriorityValue = "PRIORITY_UNSPECIFIED"
	NotificationPriorityMin         NotificationPriorityValue = "PRIORITY_MIN"
	NotificationPriorityLow         NotificationPriorityValue = "PRIORITY_LOW"
	NotificationPriorityDefault     NotificationPriorityValue = "PRIORITY_DEFAULT"
	NotificationPriorityHigh        NotificationPriorityValue = "PRIORITY_HIGH"
	NotificationPriorityMax         NotificationPriorityValue = "PRIORITY_MAX"

	VisibilityUnspecified VisibilityValue = "VISIBILITY_UNSPECIFIED"
	VisibilityPrivate     VisibilityValue = "PRIVATE"
	VisibilityPublic      VisibilityValue = "PUBLIC"
	VisibilitySecret      VisibilityValue = "SECRET"
)

type AndroidMessagePriority struct {
	Value MessagePriorityValue
}

type NotificationPriority struct {
	Value NotificationPriorityValue
}

type Visibility struct {
	Value VisibilityValue
}
