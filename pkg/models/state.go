package models

import "fmt"

type State struct {
	Id   int    `json:"stateID"`
	Name string `json:"stateName"`
}

func (s *State) String() string {
	return fmt.Sprintf("Id: %d, Name: %s", s.Id, s.Name)
}
