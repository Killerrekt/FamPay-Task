package service

import (
	"errors"

	"github.com/killerrekt/fampay-task/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VideoService interface {
	SaveBulkVideo([]model.Video) error
	GetQuery() ([]string, error)
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

func (v videoServices) GetQuery() ([]string, error) {
	var queries []string
	err := v.DB.Model(&model.Video{}).Select("DISTINCT(LOWER(query))").Find(&queries).Error
	return queries, err
}
