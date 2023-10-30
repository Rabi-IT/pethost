package errors_case

import "github.com/gofiber/fiber/v2"

func BadRequest(err error) error {
	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: err.Error(),
	}
}
