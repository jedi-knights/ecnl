package models

import "fmt"

type EventType struct {
	Id   int    `json:"eventTypeID"`
	Name string `json:"eventType"`
}

func (e *EventType) ToString() string {
	return fmt.Sprintf("Id: %d, Name: %s", e.Id, e.Name)
}
