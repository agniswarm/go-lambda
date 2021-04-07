package controller

import (
	"encoding/json"
	"net/http"

	"github.com/agniswarm/go-lambda/models"
	"github.com/agniswarm/go-lambda/response"
	"github.com/agniswarm/go-lambda/services"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {

	bk := new(models.Book)
	// This step is already performed in the middleware
	json.Unmarshal(c.Body(), bk)

	err := services.CreateBook(bk)
	if err != nil {
		return response.ReportError(c, err, http.StatusInternalServerError)
	}
	return c.Status(http.StatusCreated).JSON(bk)
}
