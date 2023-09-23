package models

import "fmt"

type Organization struct {
	Id            int    `json:"orgID"`
	SeasonId      int    `json:"orgSeasonID"`
	Name          string `json:"orgName"`
	SeasonGroupId int    `json:"orgSeasonGroupID"`
}

func (o *Organization) ToString() string {
	return fmt.Sprintf("Name: \"%s\", Id: %d, SeasonId: %d, SeasonGroupId: %d", o.Name, o.Id, o.SeasonId, o.SeasonGroupId)
}
