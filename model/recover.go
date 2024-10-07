package model

import "time"

type Recover struct {
	Query          string    `gorm:"column:query"`
	PublishingTime time.Time `gorm:"column:max_publishingtime"`
}
