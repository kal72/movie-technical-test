package router

import (
	"movie-technical-test/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Route struct {
	App           *fiber.App
	LogMiddleware fiber.Handler
	MovieHandler  *handler.MovieHandler
}

func (r *Route) Setup() {
	r.App.Use(recover.New())
	r.App.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Server running...",
		})
	})

	r.App.Use(r.LogMiddleware)
	r.App.Post("/Movies", r.MovieHandler.AddNew)
	r.App.Get("/Movies", r.MovieHandler.List)
	r.App.Get("/Movies/:ID", r.MovieHandler.Detail)
	r.App.Patch("/Movies/:ID", r.MovieHandler.Update)
	r.App.Delete("/Movies/:ID", r.MovieHandler.Delete)
}
