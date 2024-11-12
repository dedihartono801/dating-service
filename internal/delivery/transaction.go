package delivery

import (
	"dating-service/internal/app/usecase/transaction"
	"dating-service/internal/entity"
	"dating-service/pkg/constant"
	"dating-service/pkg/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler interface {
	Purchase(ctx *fiber.Ctx) error
	PaymentMethodList(c *fiber.Ctx) error
	PackageList(c *fiber.Ctx) error
}

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) TransactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) Purchase(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	request := new(entity.TransactionRequest)
	if err := c.BodyParser(request); err != nil {
		err = helper.Error(http.StatusBadRequest, constant.MsgInvalidRequest, err)
		return helper.ResponseError(c, err)
	}

	err := h.service.Purchase(ctx, request)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseCreatedOK(c, "success", nil)
}

func (h *transactionHandler) PackageList(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	resp, err := h.service.PackageList(ctx)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseOK(c, "success", resp)
}

func (h *transactionHandler) PaymentMethodList(c *fiber.Ctx) error {
	ctx, cancel := helper.CreateContextWithTimeout()
	defer cancel()
	ctx = helper.SetValueToContext(ctx, c)

	resp, err := h.service.PaymentMethodList(ctx)
	if err != nil {
		return helper.ResponseError(c, err)
	}
	return helper.ResponseOK(c, "success", resp)
}
