package models

import "fmt"

type Country struct {
	Id   int    `json:"countryID"`
	Name string `json:"countryName"`
}

func (c *Country) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s", c.Id, c.Name)
}
