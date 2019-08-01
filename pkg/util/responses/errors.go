package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// generateResponse combines the status code and the message and returns an echo response
func generateResponse(code int, message string) *echo.HTTPError {
	return echo.NewHTTPError(code, message)
}

// BadRequest generate a bad request response
func BadRequest(message string) *echo.HTTPError {
	return generateResponse(http.StatusBadRequest, message)
}

// InternalServerError generate an internal server error response
func InternalServerError(message string) *echo.HTTPError {
	return generateResponse(http.StatusInternalServerError, message)
}

// NotFound generate a not found response
func NotFound(message string) *echo.HTTPError {
	return generateResponse(http.StatusNotFound, message)
}
