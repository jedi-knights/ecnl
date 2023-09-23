package models

import "fmt"

type Club struct {
	OrgId       int    `json:"orgID"`
	OrgSeasonId int    `json:"orgSeasonID"`
	ClubId      int    `json:"clubID"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ClubLogo    string `json:"clubLogo"`
	StateCode   string `json:"stateCode"`
	EventId     int    `json:"eventID"`
	EventCounts int    `json:"eventCounts"`
}

func (c *Club) ToString() string {
	return fmt.Sprintf("OrgId: %d, OrgSeasonId: %d, ClubId: %d, Name: %s, City: %s, ClubLogo: %s, StateCode: %s, EventId: %d, EventCounts: %d", c.OrgId, c.OrgSeasonId, c.ClubId, c.Name, c.City, c.ClubLogo, c.StateCode, c.EventId, c.EventCounts)
}
