package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RPIEvent struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	TeamId    int                `bson:"team_id" json:"team_id"`
	TeamName  string             `bson:"team_name" json:"team_name"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Ranking   int                `bson:"ranking" json:"ranking"`
	Value     float64            `bson:"rpi" json:"rpi"`
}

func NewRPIEvent(timestamp time.Time, data RPIRankingData) *RPIEvent {
	return &RPIEvent{
		Timestamp: timestamp,
		TeamId:    data.TeamId,
		TeamName:  data.TeamName,
		Ranking:   data.Ranking,
		Value:     data.RPI,
	}
}

func (e RPIEvent) String() string {
	return fmt.Sprintf("Timestamp: %s, TeamId: %d, TeamName: '%s', Ranking: %d, RPI: %f", e.Timestamp, e.TeamId, e.TeamName, e.Ranking, e.Value)
}
