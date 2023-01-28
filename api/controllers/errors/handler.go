package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ApiErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	errorResponse := make(map[string]interface{})
	errorResponse["message"] = err.Error()

	// Send custom error page
	ctx.Status(code).JSON(errorResponse)

	// Return from handler
	return nil
}
