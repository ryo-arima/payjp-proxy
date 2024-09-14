package model

import "time"

type Transfers struct {
    ID          uint `gorm:"primaryKey,autoIncrement"`
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}