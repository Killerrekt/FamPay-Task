package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/fampay-task/controller"
	"github.com/killerrekt/fampay-task/service"
)

func SetUpRoute(vid_service service.VideoService, incomingRoutes *fiber.App) {
	vid_handler := controller.NewVideoController(vid_service)

	incomingRoutes.Post("/set-mode", vid_handler.StartService)
	incomingRoutes.Post("/set-query", vid_handler.SetQueryParameter)

	incomingRoutes.Get("/settings", vid_handler.CurrentSettings)
}
