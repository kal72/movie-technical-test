package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Route struct {
	App           *fiber.App
	LogMiddleware fiber.Handler
}

func (r *Route) Setup() {
	r.App.Use(recover.New())
	r.App.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Server running...",
		})
	})

	appV1 := r.App.Group("/v1")
	appV1.Use(r.LogMiddleware)
}
