package model

type Video struct {
	VideoID        string `gorm:"primaryKey"`
	Videotitle     string
	Description    string
	Publishingtime string
	Thumbnails     string
}
