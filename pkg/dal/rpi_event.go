package dal

import (
	"context"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type RPIEventDAOer interface {
	Index() error
	GetByTeamId(teamId int) ([]models.RPIEvent, error)
	GetByTeamName(teamName string) ([]models.RPIEvent, error)
	Create(rpiEvent models.RPIEvent) error
	DeleteByTeamId(teamId int) error
	DeleteByTeamName(teamName string) error
}

type RPIEventDAO struct {
	ctx context.Context
	col *mongo.Collection
}

func NewRPIEventDAO(ctx context.Context, col *mongo.Collection) *RPIEventDAO {
	return &RPIEventDAO{ctx: ctx, col: col}
}

// Index indexes the collection.
func (dao *RPIEventDAO) Index() error {
	var (
		name string
		err  error
	)

	// create an index on the team_id field
	if name, err = dao.col.Indexes().CreateOne(dao.ctx, mongo.IndexModel{
		Keys: bson.D{{"team_id", 1}},
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s on rpi_events collection", name)
	}

	// create an index on the team_name field
	if name, err = dao.col.Indexes().CreateOne(dao.ctx, mongo.IndexModel{
		Keys: bson.D{{"team_name", 1}},
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s on rpi_events collection", name)
	}

	return nil
}

// GetByTeamId gets a collection of RPI events by team id.
func (dao *RPIEventDAO) GetByTeamId(teamId int) ([]models.RPIEvent, error) {
	var (
		err     error
		cursor  *mongo.Cursor
		bEvents []bson.M
		event   models.RPIEvent
		events  []models.RPIEvent
	)

	if cursor, err = dao.col.Find(dao.ctx, bson.M{"team_id": teamId}); err != nil {
		return nil, err
	}
	if err = cursor.All(dao.ctx, &bEvents); err != nil {
		return nil, err
	}

	for _, bEvent := range bEvents {
		bsonBytes, _ := bson.Marshal(bEvent)
		if err = bson.Unmarshal(bsonBytes, &event); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// GetByTeamName gets a collection of RPI events by team name.
func (dao *RPIEventDAO) GetByTeamName(teamName string) ([]models.RPIEvent, error) {
	var (
		err     error
		cursor  *mongo.Cursor
		bEvents []bson.M
		event   models.RPIEvent
		events  []models.RPIEvent
	)

	if cursor, err = dao.col.Find(dao.ctx, bson.M{"team_name": teamName}); err != nil {
		return nil, err
	}
	if err = cursor.All(dao.ctx, &bEvents); err != nil {
		return nil, err
	}

	for _, bEvent := range bEvents {
		bsonBytes, _ := bson.Marshal(bEvent)
		if err = bson.Unmarshal(bsonBytes, &event); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// Create creates an RPI event.
func (dao *RPIEventDAO) Create(rpiEvent models.RPIEvent) error {
	var err error

	if _, err = dao.col.InsertOne(dao.ctx, rpiEvent); err != nil {
		return err
	}

	return nil
}

// DeleteByTeamId deletes a collection of RPI events by team id.
func (dao *RPIEventDAO) DeleteByTeamId(teamId int) error {
	var (
		err          error
		deleteResult *mongo.DeleteResult
	)

	if deleteResult, err = dao.col.DeleteMany(dao.ctx, bson.M{"team_id": teamId}); err != nil {
		return err
	}

	log.Printf("deleted %d rpi events for team id %d", deleteResult.DeletedCount, teamId)

	// Check to see if the delete was successful.
	if deleteResult.DeletedCount > 0 {
		return fmt.Errorf("delete failed")
	}

	// The delete was successful.
	log.Printf("rpi event delete successful: %v", deleteResult)

	return nil
}

// DeleteByTeamName deletes a collection of RPI event by team name.
func (dao *RPIEventDAO) DeleteByTeamName(teamName string) error {
	var (
		err          error
		deleteResult *mongo.DeleteResult
	)

	if deleteResult, err = dao.col.DeleteMany(dao.ctx, bson.M{"team_name": teamName}); err != nil {
		return err
	}

	log.Printf("deleted %d rpi events for team id %s", deleteResult.DeletedCount, teamName)

	// Check to see if the delete was successful.
	if deleteResult.DeletedCount > 0 {
		return fmt.Errorf("delete failed")
	}

	// The delete was successful.
	log.Printf("rpi event delete successful: %v", deleteResult)

	return nil
}
