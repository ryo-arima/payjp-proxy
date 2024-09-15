package model

type Customers struct {
	Cards         Card         `json:"cards"`
	Created       int64        `json:"created"`
	DefaultCard   interface{}  `json:"default_card"`
	Description   string       `json:"description"`
	Email         interface{}  `json:"email"`
	ID            string       `json:"id"`
	Livemode      bool         `json:"livemode"`
	Metadata      interface{}  `json:"metadata"`
	Object        string       `json:"object"`
	Subscriptions Subscription `json:"subscriptions"`
}
