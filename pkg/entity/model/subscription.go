package model

import "time"

type Subscription struct {
	ID                 string      `json:"id"`
	CanceledAt         interface{} `json:"canceled_at"`
	Created            int64       `json:"created"`
	CurrentPeriodEnd   int64       `json:"current_period_end"`
	CurrentPeriodStart int64       `json:"current_period_start"`
	Customer           string      `json:"customer"`
	Livemode           bool        `json:"livemode"`
	Metadata           interface{} `json:"metadata"`
	NextCyclePlan      interface{} `json:"next_cycle_plan"`
	Object             string      `json:"object"`
	PausedAt           interface{} `json:"paused_at"`
	Plan               Plan        `json:"plan"`
	ResumedAt          interface{} `json:"resumed_at"`
	Start              int64       `json:"start"`
	Status             string      `json:"status"`
	TrialEnd           interface{} `json:"trial_end"`
	TrialStart         interface{} `json:"trial_start"`
	Prorate            bool        `json:"prorate"`
	CreatedAt          *time.Time  `json:"-"`
	UpdatedAt          *time.Time  `json:"-"`
	DeletedAt          *time.Time  `json:"-"`
}

type Subscriptions struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
