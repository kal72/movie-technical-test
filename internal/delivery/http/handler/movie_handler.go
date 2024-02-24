package handler

import (
	"movie-technical-test/internal/model"
	"movie-technical-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type MovieHandler struct {
	UseCase *usecase.MovieUseCase
}

func NewMovieHandler(useCase *usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{
		UseCase: useCase,
	}
}

func (h *MovieHandler) AddNew(ctx *fiber.Ctx) error {
	var newCtx = ctx.UserContext()
	var request model.MovieCreateRequest
	var err error

	err = ctx.BodyParser(&request)
	if err != nil {
		h.UseCase.Log.Error(newCtx, err)
		return fiber.ErrInternalServerError
	}

	h.UseCase.Log.StartRequest(newCtx, request)

	response, err := h.UseCase.AddNew(newCtx, request)
	if err != nil {
		return err
	}

	resp := model.Response[fiber.Map]{
		Status:  "00",
		Data:    response,
		Message: "Success",
	}
	h.UseCase.Log.FinishRequest(newCtx, request, resp)
	return ctx.JSON(resp)
}
