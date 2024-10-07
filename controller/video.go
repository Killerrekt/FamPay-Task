package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/fampay-task/model"
	"github.com/killerrekt/fampay-task/service"
	"github.com/killerrekt/fampay-task/utils"
	"google.golang.org/api/youtube/v3"
)

type VideoController interface {
	StartService(*fiber.Ctx) error
	AddQueryParameter(*fiber.Ctx) error
	RemoveQueryParameter(*fiber.Ctx) error
	GetData(*fiber.Ctx) error
	GetQuery(*fiber.Ctx) error
	CurrentSettings(*fiber.Ctx) error
	Recover(*fiber.Ctx) error
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

func (v videoControllers) AddQueryParameter(c *fiber.Ctx) error {
	var req struct {
		Query *string `json:"query"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to the json", Status: false})
	}

	if req.Query == nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass a query", Status: false})
	}

	if _, ok := utils.QueryMap[*req.Query]; ok {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "The query is already being run", Status: false})
	}

	utils.QueryMap[*req.Query] = time.Now().Add(-24 * time.Hour).UTC() //just a buffer so that there is some data in the table
	utils.ContinuousFetch(v.ytClient, v.service, *req.Query)

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "added to the queries", Status: true})
}

func (v videoControllers) RemoveQueryParameter(c *fiber.Ctx) error {
	var req struct {
		Query *string `json:"query"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Failed to the json", Status: false})
	}

	if req.Query == nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass a query", Status: false})
	}

	delete(utils.QueryMap, *req.Query)

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Removed the query", Status: true})
}

func (v videoControllers) CurrentSettings(c *fiber.Ctx) error {
	var Data struct {
		Mode  bool     `json:"mode"`
		Query []string `json:"query"`
	}

	Data.Mode = utils.State
	for k, _ := range utils.QueryMap {
		Data.Query = append(Data.Query, k)
	}

	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Successfully got the settings", Data: Data, Status: true})
}

func (v videoControllers) GetQuery(c *fiber.Ctx) error {
	queries, err := v.service.GetQuery()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to get the query in the DB", Status: false})
	}
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Got the list of queries", Status: true, Data: queries})
}

func (v videoControllers) GetData(c *fiber.Ctx) error {
	q := c.Query("query")
	time := c.Query("publish")
	if q == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{Message: "Didn't pass a query", Status: false})
	}
	data, err := v.service.GetData(q, time)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to get the query in the DB", Status: false})
	}
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Got the list of queries", Status: true, Data: data})
}

func (v videoControllers) Recover(c *fiber.Ctx) error {
	res, err := v.service.RecoverInfo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{Message: "Failed to get the recovery started", Status: false})
	}
	for _, val := range res {
		if _, ok := utils.QueryMap[val.Query]; ok { //if already running ignore it and go to next one
			continue
		}
		utils.QueryMap[val.Query] = val.PublishingTime
		utils.ContinuousFetch(v.ytClient, v.service, val.Query)
	}
	return c.Status(fiber.StatusAccepted).JSON(model.Response{Message: "Recovery done", Status: true})
}
