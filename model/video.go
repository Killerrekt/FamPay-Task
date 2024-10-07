package model

import "time"

// index on pub_time for better read performance
type Video struct {
	VideoID        string `gorm:"primaryKey"`
	Videotitle     string
	Description    string
	Publishingtime time.Time `gorm:"type:timestamptz;index:,sort:desc"`
	Thumbnails     string
	Query          string
	ChannelId      string
	ChannelTitle   string
}
