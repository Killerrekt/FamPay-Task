package service

import (
	"github.com/killerrekt/fampay-task/model"
	"gorm.io/gorm"
)

type VideoService interface {
	SaveBulkVideo([]model.Video) error
}

type videoServices struct {
	DB *gorm.DB
}

func NewVideoService(DB *gorm.DB) VideoService {
	return &videoServices{
		DB: DB,
	}
}

func (v videoServices) SaveBulkVideo(videos []model.Video) error {
	return v.DB.Create(&videos).Error
}
