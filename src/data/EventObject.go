package data

type EventPost struct {
	Destination string     `json:"destination" validate:"required"`
	Event       []EventObj `json:"events" validate:"required"`
}

type EventObj struct {
	Type            string             `json:"type" validate:"required"`
	Message         MessageObj         `json:"message" validate:"required"`
	Timestamp       int                `json:"timestamp" validate:"required"`
	Source          sourceObj          `json:"source" validate:"required"`
	ReplyToken      string             `json:"replyToken" validate:"required"`
	Mode            string             `json:"mode" validate:"required"`
	WebhookEventId  string             `json:"webhookEventId" validate:"required"`
	DeliveryContext DeliveryContextObj `json:"deliveryContext" validate:"required"`
}

type DeliveryContextObj struct {
	IsRedelivery bool `json:"isRedelivery" validate:"required"`
}

type MessageObj struct {
	Type string `json:"type" validate:"required"`
	Id   string `json:"id" validate:"required"`
	Text string `json:"text" validate:"required"`
}

type sourceObj struct {
	Type   string `json:"type" validate:"required"`
	UserID string `json:"userID" validate:"required"`
}
