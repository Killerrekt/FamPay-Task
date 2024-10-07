package service

import (
	"errors"

	"github.com/killerrekt/fampay-task/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if len(videos) == 0 {
		return errors.New("videos is empty")
	}
	err := v.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&videos).Error
	return err
}
