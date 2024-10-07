package controller

import (
	"github.com/killerrekt/fampay-task/service"
	"google.golang.org/api/youtube/v3"
)

type VideoHandler interface {
}

type videoHandlers struct {
	service  service.VideoService
	ytClient *youtube.Service
}

func NewVideoService(
	service service.VideoService,
	ytClient *youtube.Service,
) VideoHandler {
	return &videoHandlers{
		service:  service,
		ytClient: ytClient,
	}
}
