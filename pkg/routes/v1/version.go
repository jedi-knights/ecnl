package v1

import (
	"github.com/jedi-knights/ecnl/pkg/controllers"
	"github.com/jedi-knights/ecnl/pkg/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

// HandleGetVersion godoc
// @Summary Get the API's current version
// @Description Get the current version of the API
// @Tags Admin
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.VersionResponse
// @Router /v1/version [get]
func HandleVersion(e echo.Context) error {
	var (
		err     error
		version string
	)
	e.Logger().Debug("HandleVersion")

	if version, err = controllers.NewVersion().Get(); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, responses.VersionResponse{
		Version: version,
	})
}
