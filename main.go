package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/killerrekt/fampay-task/utils"
)

func main() {
	app := fiber.New()

	yt_client := utils.InitYTClient()
	log.Println(yt_client)

	app.Use(logger.New())
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Pong :- Welcome to the fampay task by killerrekt",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
