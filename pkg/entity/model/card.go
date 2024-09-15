package model

type Card struct {
	AddressCity        interface{}            `json:"address_city"`
	AddressLine1       interface{}            `json:"address_line1"`
	AddressLine2       interface{}            `json:"address_line2"`
	AddressState       interface{}            `json:"address_state"`
	AddressZip         interface{}            `json:"address_zip"`
	AddressZipCheck    string                 `json:"address_zip_check"`
	Brand              string                 `json:"brand"`
	Country            interface{}            `json:"country"`
	Created            int64                  `json:"created"`
	Customer           interface{}            `json:"customer"`
	CvcCheck           string                 `json:"cvc_check"`
	ExpMonth           int                    `json:"exp_month"`
	ExpYear            int                    `json:"exp_year"`
	Fingerprint        string                 `json:"fingerprint"`
	ID                 string                 `json:"id"`
	Last4              string                 `json:"last4"`
	Livemode           bool                   `json:"livemode"`
	Metadata           map[string]interface{} `json:"metadata"`
	Name               string                 `json:"name"`
	Object             string                 `json:"object"`
	ThreeDSecureStatus string                 `json:"three_d_secure_status"`
	Email              string                 `json:"email"`
	Phone              string                 `json:"phone"`
}
