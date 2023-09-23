package models

import "fmt"

type Organization struct {
	Id            int    `json:"orgID"`
	SeasonId      int    `json:"orgSeasonID"`
	Name          string `json:"orgName"`
	SeasonGroupId int    `json:"orgSeasonGroupID"`
}

func (o *Organization) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s, SeasonId: %d, SeasonGroupId: %d", o.Id, o.Name, o.SeasonId, o.SeasonGroupId)
}
