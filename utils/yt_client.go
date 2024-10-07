package utils

import (
	"context"
	"os"

	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var DevKey = os.Getenv("YT_KEY")

func InitYTClient() *youtube.Service {
	log.Println("Dev key being used :- ", DevKey)
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(DevKey))
	if err != nil {
		log.Fatal("Failed to create a YT client")
	}
	return youtubeService
}
