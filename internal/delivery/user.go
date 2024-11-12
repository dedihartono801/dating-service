package delivery

import (
	"dating-service/internal/app/usecase/user"
	"dating-service/internal/entity"
	"dating-service/pkg/constant"
	"dating-service/pkg/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Signup(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	UserList(ctx *fiber.Ctx) error
	Swipe(ctx *fiber.Ctx) error
}

type userHandler struct {
	service user.Service
}

func NewSignupHandler(service user.Service) UserHandler {
	return &userHandler{service}
}

func (h *userHandler) Signup(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.SignupRequest)
	if err := c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return helper.ResponseError(c, err)
	}

	err := h.service.Signup(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseCreatedOK(c, "success", nil)
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.LoginRequest)
	if err := c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return helper.ResponseError(c, err)
	}

	resp, err := h.service.Login(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseCreatedOK(c, "success", resp)
}

func (h *userHandler) UserList(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	resp, err := h.service.UserList(ctx)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseCreatedOK(c, "success", resp)
}

func (h *userHandler) Swipe(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.SwipeRequest)
	if err := c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return helper.ResponseError(c, err)
	}

	err := h.service.Swipe(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseCreatedOK(c, "success", nil)
}
