package api

import (
	"github.com/dukerspace/drivehub-api/api/car"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {

	car.Router(app)
}
