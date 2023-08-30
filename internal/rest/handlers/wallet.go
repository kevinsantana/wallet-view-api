package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)


func GetWalletBalance(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}