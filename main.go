package main

import (
	"github.com/dukerspace/drivehub-api/api"
	"github.com/dukerspace/drivehub-api/internal/constant"
	"github.com/dukerspace/drivehub-api/internal/database/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	helper.LoadFile(constant.DB_NAME)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(requestid.New())
	api.Routes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	app.Listen(":3000")
}
