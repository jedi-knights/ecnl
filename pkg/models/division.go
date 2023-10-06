package models

import "fmt"

type Division struct {
	Id   int    `json:"divisionID"`
	Name string `json:"divisionName"`
}

func (d *Division) String() string {
	return fmt.Sprintf("Name: \"%s\", Id: %d", d.Name, d.Id)
}
