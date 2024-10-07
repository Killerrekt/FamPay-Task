package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/fampay-task/controller"
	"github.com/killerrekt/fampay-task/service"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
)

func SetUpRoute(ytClient *youtube.Service, DB *gorm.DB, incomingRoutes *fiber.App) {
	vid_service := service.NewVideoService(DB)
	vid_handler := controller.NewVideoController(vid_service, ytClient)
	incomingRoutes.Post("/set-mode", vid_handler.StartService)
}
