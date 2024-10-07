package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/fampay-task/model"
	"github.com/killerrekt/fampay-task/service"
	"github.com/killerrekt/fampay-task/utils"
	"google.golang.org/api/youtube/v3"
)

type VideoController interface {
	StartService(*fiber.Ctx) error
	SetQueryParameter(*fiber.Ctx) error
}

type videoControllers struct {
	service  service.VideoService
	ytClient *youtube.Service
}

func NewVideoController(
	service service.VideoService,
	ytClient *youtube.Service,
) VideoController {
	return &videoControllers{
		service:  service,
		ytClient: ytClient,
	}
}

func (v videoControllers) StartService(c *fiber.Ctx) error {
	var req struct {
		Mode *bool `json:"mode"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to the json", Status: false})
	}
	if req.Mode == nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass a mode", Status: false})
	}
	utils.State = *req.Mode
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Changed the mode", Status: true})
}

func (v videoControllers) SetQueryParameter(c *fiber.Ctx) error {
	var req struct {
		Query *string `json:"query"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to the json", Status: false})
	}
	if req.Query == nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass a query", Status: false})
	}
	utils.Query = *req.Query
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Changed the query", Status: true})
}
