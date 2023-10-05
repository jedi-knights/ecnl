package models

import "fmt"

type MatchEvent struct {
	MatchId        int    `json:"matchID"`
	GameDate       string `json:"gameDate"`
	HomeTeamId     int    `json:"homeTeamID"`
	HomeTeamName   string `json:"homeTeam"`
	HomeTeamClubId int    `json:"homeTeamClubID"`
	HomeTeamScore  int    `json:"homeTeamScore"`
	AwayTeamId     int    `json:"awayTeamID"`
	AwayTeamName   string `json:"awayTeam"`
	AwayTeamClubId int    `json:"awayTeamClubID"`
	AwayTeamScore  int    `json:"awayTeamScore"`
	Flight         string `json:"flight"`
	Division       string `json:"division"`
	EventName      string `json:"eventName"`
	Complex        string `json:"complex"`
	Venue          string `json:"venue"`
}

func (m MatchEvent) String() string {
	return fmt.Sprintf("'%s' vs '%s' at '%s'", m.HomeTeamName, m.AwayTeamName, m.GameDate)
}
