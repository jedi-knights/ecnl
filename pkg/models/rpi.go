package models

import (
	"fmt"
)

type RPIRankingData struct {
	TeamId   int
	TeamName string
	RPI      float64
	Ranking  int
}

func (d RPIRankingData) String() string {
	return fmt.Sprintf("#%d: '%s' (%f)", d.Ranking, d.TeamName, d.RPI)
}
