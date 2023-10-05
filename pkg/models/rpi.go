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
	return fmt.Sprintf("TeamId: %d, Name: '%s', Ranking: %d, RPI: %f", d.TeamId, d.TeamName, d.Ranking, d.RPI)
}
