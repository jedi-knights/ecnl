package v1

import (
	"github.com/jedi-knights/ecnl/pkg/controllers"
	"github.com/jedi-knights/ecnl/pkg/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

// HandleHealthCheck godoc
// @Summary Health Check
// @Description Check if the API is up and running
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} responses.HealthCheckResponse
// @Router /v1/health [get]
func HandleHealthCheck(e echo.Context) error {
	var (
		err     error
		message string
	)
	e.Logger().Debug("HandleHealthCheck")

	if message, err = controllers.NewHealth().Check(); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, responses.HealthCheckResponse{
		Message: message,
	})
}
