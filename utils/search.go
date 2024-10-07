package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/killerrekt/fampay-task/model"
	"github.com/killerrekt/fampay-task/service"
	"google.golang.org/api/youtube/v3"
)

var State = false
var QueryMap = make(map[string]time.Time)

func Search(service *youtube.Service, q string, prev time.Time) []model.Video {
	call := service.Search.List([]string{"id", "snippet"}).
		Q(q).
		MaxResults(50).
		PublishedAfter(prev.Format(time.RFC3339)).
		Order("date").
		Type("video")

	fmt.Println(call)

	response, err := call.Do()
	if err != nil {
		log.Fatalln(err.Error())
	}

	var videos []model.Video
	for _, item := range response.Items {
		if item.Id.VideoId == "" {
			continue
		}
		err, temp := YtItemToVideo(item, q)
		if err != nil {
			continue
		}
		videos = append(videos, temp)
	}
	return videos
}

func YtItemToVideo(item *youtube.SearchResult, q string) (error, model.Video) {
	timeparsed, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
	if err != nil {
		return err, model.Video{}
	}

	return nil, model.Video{
		VideoID:        item.Id.VideoId,
		Videotitle:     item.Snippet.Title,
		Description:    item.Snippet.Description,
		Publishingtime: timeparsed,
		Thumbnails:     item.Snippet.Thumbnails.Default.Url,
		Query:          q,
		ChannelId:      item.Snippet.ChannelId,
		ChannelTitle:   item.Snippet.ChannelTitle,
	}
}

func ContinuousFetch(ytClient *youtube.Service, vid_service service.VideoService, query string) {
	q := query
	ticker := time.NewTicker(10 * time.Second)
	go func(ytClient *youtube.Service, vid_service service.VideoService, q string) {
		prev := QueryMap[q]
		for {
			select {
			case <-ticker.C:
				if State {
					if _, ok := QueryMap[q]; !ok {
						log.Println("This goroutine has been stopped as query ", q, " being removed")
						return
					}
					videos := Search(ytClient, q, prev)
					prev = time.Now().UTC()
					err := vid_service.SaveBulkVideo(videos)
					if err != nil {
						log.Println("DB Error :- ", err.Error())
					}
				} else {
					log.Println("The State is currently at False and the query is : ", q)
				}
			}
		}
	}(ytClient, vid_service, q)
}
