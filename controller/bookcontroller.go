package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Hello(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("ok")
}
