package models

import "fmt"

type Team struct {
	Id          int    `json:"teamID"`
	Name        string `json:"teamName"`
	ClubId      int    `json:"clubID"`
	InitialSeed int    `json:"initialSeed"`
	ClubLogo    string `json:"clubLogo"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	AgeGroup    string `json:"ageGroup"`
}

func (t *Team) String() string {
	return fmt.Sprintf("Name: \"%s\", Id: %d, ClubId: %d", t.Name, t.Id, t.ClubId)
}
