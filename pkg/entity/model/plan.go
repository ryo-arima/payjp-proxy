package model

type Plan struct {
	Amount     int         `json:"amount"`
	BillingDay interface{} `json:"billing_day"`
	Created    int64       `json:"created"`
	Currency   string      `json:"currency"`
	ID         string      `json:"id"`
	Interval   string      `json:"interval"`
	Livemode   bool        `json:"livemode"`
	Metadata   interface{} `json:"metadata"`
	Name       interface{} `json:"name"`
	Object     string      `json:"object"`
	TrialDays  int         `json:"trial_days"`
}
