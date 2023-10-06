package v1

import (
	"github.com/jedi-knights/ecnl/pkg/controllers"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"strconv"
)

// HandleGetRPIRankings godoc
// @Summary Examines the schedule and calculates the RPI rankings for all teams
// @Description Calculates the RPI rankings for all teams
// @Tags RPI
// @Accept json
// @Produce json
// @Param division path string true "Division" Enums(G2006/2005,G2008,G2009,G2010,G2011,B2006/2005,B2008,B2009,B2010,B2011)
// @Success 200 {array} models.RPIRankingData
// @Router /v1/rpi/{division} [get]
func HandleGetRPIRankings(c echo.Context) error {
	// read the query parameters
	var err error
	var rankingData []models.RPIRankingData

	// read path parameters
	division := c.Param("division")

	if division, err = url.QueryUnescape(division); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rpiController := controllers.NewRPI()

	if rankingData, err = rpiController.GenerateRankings(division); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("X-Element-Count", strconv.Itoa(len(rankingData)))

	return c.JSON(200, rankingData)
}
