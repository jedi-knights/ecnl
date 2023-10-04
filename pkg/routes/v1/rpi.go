package v1

import "github.com/labstack/echo/v4"

// HandleGetRPIRankings godoc
// @Summary Examines the schedule and calculates the RPI rankings for all teams
// @Description Calculates the RPI rankings for all teams
// @Tags RPI
// @Accept json
// @Produce json
// @Param ageGroup path string true "Age Group"
// @Success 200 {array} models.TeamRPI
// @Router /v1/rpi/{ageGroup} [get]
func HandleGetRPIRankings(c echo.Context) error {
	//gender := c.QueryParam("gender")
	//ageGroup := c.QueryParam("ageGroup")

	return nil
}
