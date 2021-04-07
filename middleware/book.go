package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/agniswarm/go-lambda/models"
	"github.com/agniswarm/go-lambda/response"
	"github.com/gofiber/fiber/v2"
)

func BookValidator() fiber.Handler {

	return func(c *fiber.Ctx) error {
		bk := new(models.Book)

		if err := json.Unmarshal(c.Body(), bk); err != nil {
			return response.ReportError(c, err, http.StatusUnprocessableEntity)
		}

		if err := bk.ValidateSchema(); err != nil {
			return response.ReportError(c, err, http.StatusBadRequest)
		}
		return c.Next()
	}
}
