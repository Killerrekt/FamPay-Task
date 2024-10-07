package service

import (
	"errors"
	"time"

	"github.com/killerrekt/fampay-task/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VideoService interface {
	SaveBulkVideo([]model.Video) error
	GetQuery() ([]string, error)
	GetData(string, string) (model.GetVid, error)
	RecoverInfo() ([]model.Recover, error)
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

func (v videoServices) GetData(q string, publish string) (model.GetVid, error) {
	var res model.GetVid
	query := v.DB.Model(&model.Video{}).Where("query = ?", q)
	if publish != "" {
		Time, err := time.Parse(time.RFC3339, publish)
		if err != nil {
			return res, err
		}
		query = query.Where("publishingtime < ?", Time)
	}
	err := query.Order("publishingtime desc").Limit(11).Find(&res.Videos).Error
	if err != nil {
		return res, err
	}
	if len(res.Videos) == 11 {
		res.NextPage = true
		res.Videos = res.Videos[:10]
	} else {
		res.NextPage = false
	}
	return res, err
}

func (v videoServices) RecoverInfo() ([]model.Recover, error) {
	var results []model.Recover
	err := v.DB.Model(&model.Video{}).Select("query, MAX(publishingtime) AS max_publishingtime").Group("query").Find(&results).Error
	return results, err
}
