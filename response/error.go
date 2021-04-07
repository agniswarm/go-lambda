package response

import (
	"log"
	"net/http"
	"os"

	"github.com/agniswarm/go-lambda/models"
	"github.com/gofiber/fiber/v2"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func ReportError(c *fiber.Ctx, err error, statusCode int) error {
	errorLogger.Println(err.Error())
	return c.Status(statusCode).JSON(models.Error{Message: http.StatusText(statusCode)})
}

func ReportWithoutError(c *fiber.Ctx, statusCode int) error {
	return c.Status(statusCode).JSON(models.Error{Message: http.StatusText(statusCode)})
}
