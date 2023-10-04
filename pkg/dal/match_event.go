package dal

import (
	"context"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MatchEventDAOer interface {
	Index() error
	GetAll() ([]models.MatchEvent, error)
	GetById(id int) (*models.MatchEvent, error)
	GetByDivision(division string) ([]models.MatchEvent, error)
	GetByHomeTeamName(teamName string) ([]models.MatchEvent, error)
	GetByAwayTeamName(teamName string) ([]models.MatchEvent, error)
	GetByTeamName(teamName string) ([]models.MatchEvent, error)
	GetByHomeTeamId(teamId int) ([]models.MatchEvent, error)
	GetByAwayTeamId(teamId int) ([]models.MatchEvent, error)
	GetByTeamId(teamId int) ([]models.MatchEvent, error)
	Update(matchEvent models.MatchEvent) error
	Delete(matchEvent models.MatchEvent) error
	DeleteById(id int) error
	Create(matchEvent models.MatchEvent) error
	Exists(matchEvent models.MatchEvent) (bool, error)
	ExistsById(id int) (bool, error)
	Sync(matchEvent models.MatchEvent) error
	SyncAll(matchEvents []models.MatchEvent) error
}

// MatchEventDAO is the data access object for match events.
type MatchEventDAO struct {
	ctx context.Context
	col *mongo.Collection
}

// NewMatchEventDAO creates a new match event data access object.
func NewMatchEventDAO(ctx context.Context, col *mongo.Collection) *MatchEventDAO {
	return &MatchEventDAO{ctx: ctx, col: col}
}

// Index indexes the collection.
func (dao *MatchEventDAO) Index() error {
	var (
		err  error
		name string
	)

	matchEventsIndexModel := mongo.IndexModel{
		Keys: bson.D{{"matchid", 1}, {"hometeamname", 1}, {"awayteamname", 1}, {"division", 1}},
	}

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, matchEventsIndexModel); err != nil {
		return err
	}

	log.Printf("created index %s on match events collection", name)

	return nil
}

// GetAll gets all match events.
func (dao *MatchEventDAO) GetAll() ([]models.MatchEvent, error) {
	cursor, err := dao.col.Find(dao.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M
	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonbytes, _ := bson.Marshal(bMatchEvent)

		if err = bson.Unmarshal(bsonbytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetById gets a match event by id.
func (dao *MatchEventDAO) GetById(id int) (*models.MatchEvent, error) {
	var (
		err         error
		matchEvents bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"matchid": id}).Decode(&matchEvents); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("match event not found")
		}

		return nil, err
	}

	var matchEvent models.MatchEvent

	bsonBytes, _ := bson.Marshal(matchEvents)
	if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
		return nil, err
	}

	return &matchEvent, nil
}

// GetByDivision gets match events by division.
func (dao *MatchEventDAO) GetByDivision(division string) ([]models.MatchEvent, error) {
	var (
		err error
	)

	cursor, err := dao.col.Find(dao.ctx, bson.M{"division": division})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M

	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonBytes, _ := bson.Marshal(bMatchEvent)
		if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetByHomeTeamName gets match events by home team name.
func (dao *MatchEventDAO) GetByHomeTeamName(teamName string) ([]models.MatchEvent, error) {
	var (
		err error
	)

	cursor, err := dao.col.Find(dao.ctx, bson.M{"hometeamname": teamName})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M

	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonBytes, _ := bson.Marshal(bMatchEvent)
		if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetByAwayTeamName gets match events by away team name.
func (dao *MatchEventDAO) GetByAwayTeamName(teamName string) ([]models.MatchEvent, error) {
	var (
		err error
	)

	cursor, err := dao.col.Find(dao.ctx, bson.M{"awayteamname": teamName})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M

	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonBytes, _ := bson.Marshal(bMatchEvent)
		if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetByTeamName gets match events by team name.
// This one is a little trickey because we have to check both home and away team names.
func (dao *MatchEventDAO) GetByTeamName(teamName string) ([]models.MatchEvent, error) {
	var (
		err        error
		homeEvents []models.MatchEvent
		awayEvents []models.MatchEvent
	)

	if homeEvents, err = dao.GetByHomeTeamName(teamName); err != nil {
		return nil, err
	}

	if awayEvents, err = dao.GetByAwayTeamName(teamName); err != nil {
		return nil, err
	}

	events := append(homeEvents, awayEvents...)

	return events, nil
}

// GetByHomeTeamId gets match events by home team id.
func (dao *MatchEventDAO) GetByHomeTeamId(teamId int) ([]models.MatchEvent, error) {
	var (
		err error
	)

	cursor, err := dao.col.Find(dao.ctx, bson.M{"hometeamid": teamId})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M

	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonBytes, _ := bson.Marshal(bMatchEvent)
		if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetByAwayTeamId gets match events by away team id.
func (dao *MatchEventDAO) GetByAwayTeamId(teamId int) ([]models.MatchEvent, error) {
	var (
		err error
	)

	cursor, err := dao.col.Find(dao.ctx, bson.M{"awayteamid": teamId})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M

	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonBytes, _ := bson.Marshal(bMatchEvent)
		if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetECNLByAgeGroup gets ECNL match events by age group.
func (dao *MatchEventDAO) GetECNLByAgeGroup(ageGroup string) ([]models.MatchEvent, error) {
	var (
		err error
	)

	cursor, err := dao.col.Find(dao.ctx, bson.M{"flight": "ECNL", "division": ageGroup})
	if err != nil {
		return nil, err
	}

	var bMatchEvents []bson.M

	if err = cursor.All(dao.ctx, &bMatchEvents); err != nil {
		return nil, err
	}

	var matchEvents []models.MatchEvent

	for _, bMatchEvent := range bMatchEvents {
		var matchEvent models.MatchEvent

		bsonBytes, _ := bson.Marshal(bMatchEvent)
		if err = bson.Unmarshal(bsonBytes, &matchEvent); err != nil {
			return nil, err
		}

		matchEvents = append(matchEvents, matchEvent)
	}

	return matchEvents, nil
}

// GetByTeamId gets match events by team id.
func (dao *MatchEventDAO) GetByTeamId(teamId int) ([]models.MatchEvent, error) {
	var (
		err        error
		homeEvents []models.MatchEvent
		awayEvents []models.MatchEvent
	)

	if homeEvents, err = dao.GetByHomeTeamId(teamId); err != nil {
		return nil, err
	}

	if awayEvents, err = dao.GetByAwayTeamId(teamId); err != nil {
		return nil, err
	}

	events := append(homeEvents, awayEvents...)

	return events, nil
}

// Update updates a match event.
func (dao *MatchEventDAO) Update(matchEvent models.MatchEvent) error {
	var (
		err          error
		updateResult *mongo.UpdateResult
	)

	log.Println("updating match event: ", matchEvent.String())

	if updateResult, err = dao.col.UpdateOne(dao.ctx, bson.M{"matchid": matchEvent.MatchId}, bson.M{"$set": matchEvent}); err != nil {
		return err
	}

	// Check to see if the update was successful.
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("update failed")
	}

	// The update was successful.
	log.Printf("match event update successful: %v", updateResult)

	return nil
}

// Delete deletes a match event.
func (dao *MatchEventDAO) Delete(matchEvent models.MatchEvent) error {
	return dao.DeleteById(matchEvent.MatchId)
}

// DeleteById deletes a match event by id.
func (dao *MatchEventDAO) DeleteById(id int) error {
	var (
		err          error
		deleteResult *mongo.DeleteResult
	)

	if deleteResult, err = dao.col.DeleteOne(dao.ctx, bson.M{"matchid": id}); err != nil {
		return err
	}

	// Check to see if the delete was successful.
	if deleteResult.DeletedCount != 1 {
		return fmt.Errorf("delete failed")
	}

	// The delete was successful.
	log.Printf("match event delete successful: %v", deleteResult)

	return nil
}

// Create creates a match event.
func (dao *MatchEventDAO) Create(matchEvent models.MatchEvent) error {
	var (
		err error
	)

	if _, err = dao.col.InsertOne(dao.ctx, matchEvent); err != nil {
		return err
	}

	return nil
}

// Exists checks if a match event exists.
func (dao *MatchEventDAO) Exists(matchEvent models.MatchEvent) (bool, error) {
	var (
		err         error
		bMatchEvent bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"hometeamname": matchEvent.HomeTeamName, "awayteamname": matchEvent.AwayTeamName, "gamedate": matchEvent.GameDate}).Decode(&bMatchEvent); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// ExistsById checks if a match event exists by id.
func (dao *MatchEventDAO) ExistsById(id int) (bool, error) {
	var (
		err         error
		matchEvents bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"matchid": id}).Decode(&matchEvents); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Sync syncs a match event.
func (dao *MatchEventDAO) Sync(matchEvent models.MatchEvent) error {
	var (
		err    error
		answer bool
	)

	// log.Printf("syncing match event %v", matchEvent.String())

	if answer, err = dao.Exists(matchEvent); err != nil {
		return err
	}

	if !answer {
		// The match event doesn't exist.
		log.Printf("inserting match event '%d' ...", matchEvent.MatchId)
		if _, err = dao.col.InsertOne(dao.ctx, matchEvent); err != nil {
			return err
		}

		// log.Printf("match event insert successful: %v", matchEvent)
		return nil
	}

	// The match event exists.
	// The code below has been commented out for performance.
	// It is unlikely that a match is going to change so we don't need to update it if we already have it.
	// log.Printf("skipping match event '%s'", matchEvent.String())

	//log.Printf("updating match event '%s' ...", matchEvent.MatchId)
	//if err = dao.Update(matchEvent); err != nil {
	//	return err
	//}
	//
	//log.Printf("match event update successful: %v", matchEvent)

	return nil
}

// SyncAll syncs all match events.
func (dao *MatchEventDAO) SyncAll(matchEvents []models.MatchEvent) error {
	for _, matchEvent := range matchEvents {
		if err := dao.Sync(matchEvent); err != nil {
			return err
		}
	}

	return nil
}
