package model

import "time"

type Video struct {
	VideoID        string `gorm:"primaryKey"`
	Videotitle     string
	Description    string
	Publishingtime time.Time `gorm:"type:timestamptz"`
	Thumbnails     string
	Query          string
	ChannelId      string
	ChannelTitle   string
}
