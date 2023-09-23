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
	return fmt.Sprintf("Name: \"%s\", OrgId: %d, OrgSeasonId: %d, ClubId: %d, City: \"%s\", StateCode: %s, EventId: %d, EventCounts: %d", c.Name, c.OrgId, c.OrgSeasonId, c.ClubId, c.City, c.StateCode, c.EventId, c.EventCounts)
}
