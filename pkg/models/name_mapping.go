package models

import "fmt"

type NameMapping struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type NameMapper []NameMapping

func NewNameMapping(from string, to string) *NameMapping {
	return &NameMapping{
		From: from,
		To:   to,
	}
}

func NewNameMapper() *NameMapper {
	return &NameMapper{}
}

func (m *NameMapping) String() string {
	return fmt.Sprintf("From: '%s', To: '%s'", m.From, m.To)
}
