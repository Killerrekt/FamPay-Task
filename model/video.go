package model

import "time"

type Video struct {
	VideoID        string `gorm:"primaryKey"`
	Videotitle     string
	Description    string
	Publishingtime time.Time
	Thumbnails     string
	Query          string
	ChannelId      string
	ChannelTitle   string
}
