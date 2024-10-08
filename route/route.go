package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/fampay-task/controller"
	"github.com/killerrekt/fampay-task/service"
	"google.golang.org/api/youtube/v3"
)

func SetUpRoute(vid_service service.VideoService, yt_client *youtube.Service, incomingRoutes *fiber.App) {
	vid_handler := controller.NewVideoController(vid_service, yt_client)

	incomingRoutes.Post("/set-mode", vid_handler.StartService)
	incomingRoutes.Post("/add-query", vid_handler.AddQueryParameter)
	incomingRoutes.Post("/remove-query", vid_handler.RemoveQueryParameter)
	incomingRoutes.Post("/recover", vid_handler.Recover)

	incomingRoutes.Get("/get-data", vid_handler.GetData)
	incomingRoutes.Get("/settings", vid_handler.CurrentSettings)
	incomingRoutes.Get("/query", vid_handler.GetQuery)
}
