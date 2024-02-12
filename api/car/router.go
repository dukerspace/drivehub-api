package car

import "github.com/gofiber/fiber/v2"

func Router(app fiber.Router) {
	v1 := app.Group("api/v1")
	repository := NewRepository()
	service := NewService(repository)

	handler := NewHandler(service)

	car := v1.Group("cars")
	car.Get("/", handler.GetAll)
	car.Post("/", handler.Create)
	car.Get("/:id", handler.GetByID)
	car.Put("/:id", handler.Update)
	car.Delete("/:id", handler.Delete)

}
