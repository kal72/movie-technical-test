package handler

import (
	"movie-technical-test/internal/model"
	"movie-technical-test/internal/usecase"
	"strconv"

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

func (h *MovieHandler) List(ctx *fiber.Ctx) error {
	var newCtx = ctx.UserContext()
	var page, size int
	var err error

	h.UseCase.Log.StartRequest(newCtx, ctx.Queries())
	page, err = strconv.Atoi(ctx.Query("page"))
	if err != nil {
		h.UseCase.Log.Error(newCtx, err)
		return fiber.NewError(fiber.StatusBadRequest, "'page' cannot empty and must be number!")
	}

	size, err = strconv.Atoi(ctx.Query("size"))
	if err != nil {
		h.UseCase.Log.Error(newCtx, err)
		return fiber.NewError(fiber.StatusBadRequest, "'size' cannot empty and must be number!")
	}

	responses, pageMetaData, err := h.UseCase.List(newCtx, page, size)
	if err != nil {
		return err
	}

	resp := model.PageResponse[model.MovieResponse]{
		Status:       "00",
		Data:         responses,
		PageMetadata: pageMetaData,
		Message:      "Success",
	}

	h.UseCase.Log.FinishRequest(newCtx, ctx.Queries(), resp)
	return ctx.JSON(resp)
}

func (h *MovieHandler) Detail(ctx *fiber.Ctx) error {
	var newCtx = ctx.UserContext()
	var id int
	var err error
	var message, status = "Success", "00"

	h.UseCase.Log.StartRequest(newCtx, ctx.Queries())
	id, err = strconv.Atoi(ctx.Params("ID"))
	if err != nil {
		h.UseCase.Log.Error(newCtx, err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid parameters")
	}

	response, err := h.UseCase.Detail(newCtx, id)
	if err != nil {
		if err != fiber.ErrNotFound {
			return err
		}

		status = "01"
		message = err.Error()
	}

	resp := model.Response[*model.MovieResponse]{
		Status:  status,
		Data:    response,
		Message: message,
	}

	h.UseCase.Log.FinishRequest(newCtx, ctx.Queries(), resp)
	return ctx.JSON(resp)
}
