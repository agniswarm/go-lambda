package controller

import (
	"net/http"
	"regexp"

	"github.com/agniswarm/go-lambda/response"
	"github.com/agniswarm/go-lambda/services"
	"github.com/gofiber/fiber/v2"
)

var isbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)

func Show(c *fiber.Ctx) error {
	isbn := c.Query("isbn")
	if !isbnRegexp.MatchString(isbn) {
		return response.ReportWithoutError(c, http.StatusBadRequest)
	}
	bk, err := services.GetBook(isbn)
	if err != nil {
		return response.ReportError(c, err, http.StatusBadRequest)
	}
	if bk == nil {
		return response.ReportWithoutError(c, http.StatusNotFound)
	}
	return c.JSON(&bk)
}
