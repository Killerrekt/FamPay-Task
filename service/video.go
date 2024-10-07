package service

import "gorm.io/gorm"

type VideoService interface {
}

type videoServices struct {
	DB *gorm.DB
}

func NewVideoService(DB *gorm.DB) VideoService {
	return &videoServices{
		DB: DB,
	}
}
