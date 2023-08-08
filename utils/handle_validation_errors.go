package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandleValidationErrors(c *fiber.Ctx, err error) error {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		validationErrors := make(map[string]string)

		for _, fieldErr := range validationErr {
			fieldName := fieldErr.Field()
			formattedFieldName := strings.ToLower(fieldName)

			tag := fieldErr.Tag()
			message := fmt.Sprintf("%s is invalid", formattedFieldName)

			switch tag {
			case "required":
				message = fmt.Sprintf("%s is required", formattedFieldName)
			case "min":
				param := fieldErr.Param()
				message = fmt.Sprintf("%s must be at least %s", formattedFieldName, param)
			case "max":
				param := fieldErr.Param()
				message = fmt.Sprintf("%s must be at most %s", formattedFieldName, param)
			}

			validationErrors[formattedFieldName] = message
		}

		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal server error",
	})
}
