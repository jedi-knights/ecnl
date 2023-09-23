package models

import "fmt"

type Event struct {
	Id            int    `json:"eventID"`
	Name          string `json:"eventName"`
	OrgId         int    `json:"orgID"`
	OrgName       string `json:"orgName"`
	OrgSeasonId   int    `json:"orgSeasonID"`
	OrgSeasonName string `json:"orgSeasonName"`
}

func (e *Event) ToString() string {
	return fmt.Sprintf("Name: \"%s\", Id: %d, OrgId: %d, OrgName: %s, OrgSeasonId: %d, OrgSeasonName: \"%s\"", e.Name, e.Id, e.OrgId, e.OrgName, e.OrgSeasonId, e.OrgSeasonName)
}
