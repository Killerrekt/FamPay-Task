package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/killerrekt/fampay-task/model"
	"google.golang.org/api/youtube/v3"
)

var State = false
var Query = ""

func Search(service *youtube.Service) []model.Video {
	q := Query //this is done to prevent possible race condition
	call := service.Search.List([]string{"id", "snippet"}).
		Q(q).
		MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Fatalln(err.Error())
	}

	var videos []model.Video
	for _, item := range response.Items {
		temp := YtItemToVideo(item, q)
		videos = append(videos, temp)
	}
	return videos
}

func YtItemToVideo(item *youtube.SearchResult, q string) model.Video {
	return model.Video{
		VideoID:        item.Id.VideoId,
		Videotitle:     item.Snippet.Title,
		Description:    item.Snippet.Description,
		Publishingtime: item.Snippet.PublishedAt,
		Thumbnails:     item.Snippet.Thumbnails.Default.Url,
		Query:          q,
		ChannelId:      item.Snippet.ChannelId,
		ChannelTitle:   item.Snippet.ChannelTitle,
	}
}

func ContinuousFetch(ytClient *youtube.Service) {
	ticker := time.NewTicker(30 * time.Second)
	go func(ytClient *youtube.Service) {
		for {
			select {
			case <-ticker.C:
				if State {
					videos := Search(ytClient)
					fmt.Println(videos)
				} else {
					log.Println("The State is currently at False and the query is : ", Query)
				}
			}
		}
	}(ytClient)
}
