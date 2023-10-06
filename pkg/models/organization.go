package models

import "fmt"

type Organization struct {
	Id            int    `json:"orgID"`
	SeasonId      int    `json:"orgSeasonID"`
	Name          string `json:"orgName"`
	SeasonGroupId int    `json:"orgSeasonGroupID"`
}

type OrganiationDivision struct {
	Id   int    `json:"divisionID"`
	Name string `json:"divisionName"`
	Text string `json:"divisionText"`
}

func (o *Organization) String() string {
	return fmt.Sprintf("Name: \"%s\", Id: %d, SeasonId: %d, SeasonGroupId: %d", o.Name, o.Id, o.SeasonId, o.SeasonGroupId)
}

func (o OrganiationDivision) String() string {
	return fmt.Sprintf("Name: \"%s\", Id: %d, Text: \"%s\"", o.Name, o.Id, o.Text)
}
