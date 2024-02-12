package car

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) GetAll(c *fiber.Ctx) error {
	r, err := h.service.GetAll()
	if err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    r,
	})
}

func (h *handler) Create(c *fiber.Ctx) error {
	data := inputCar{}
	if err := c.BodyParser(&data); err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	r, err := h.service.Create(data)
	if err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    r,
	})
}

func (h *handler) GetByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	r, err := h.service.GetByID(id)
	if err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    r,
	})
}

func (h *handler) Update(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	data := inputCar{}
	if err := c.BodyParser(&data); err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	r, err := h.service.Update(id, data)
	if err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    r,
	})
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := h.service.Delete(id)
	if err != nil {
		msg := err.Error()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": msg,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Deleted",
	})
}
