package dal

import (
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
)

type EventDAOer interface {
	Index() error
	GetAll() ([]models.Event, error)
	GetById(id int) (*models.Event, error)
	GetByName(name string) (*models.Event, error)
	Update(event models.Event) error
	Delete(event models.Event) error
	DeleteByName(name string) error
	DeleteById(id int) error
	Create(event models.Event) error
	Exists(event models.Event) (bool, error)
	ExistsByName(name string) (bool, error)
	ExistsById(id int) (bool, error)
	Sync(event models.Event) error
	SyncAll(events []models.Event) error
}

// EventDAO is the data access object for events.
type EventDAO struct {
	ctx context.Context
	col *mongo.Collection
}

// NewEventDAO creates a new event data access object.
func NewEventDAO(ctx context.Context, col *mongo.Collection) *EventDAO {
	return &EventDAO{ctx: ctx, col: col}
}

// Index indexes the collection.
func (dao *EventDAO) Index() error {
	var (
		err  error
		name string
	)

	eventsIndexModel := mongo.IndexModel{
		Keys: bson.D{{"id", 1}, {"orgid", 1}, {"name", 1}, {"orgname", 1}},
	}

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, eventsIndexModel); err != nil {
		return err
	}

	log.Printf("created index %s on events collection", name)

	return nil
}

// GetAll gets all events.
func (dao *EventDAO) GetAll() ([]models.Event, error) {
	cursor, err := dao.col.Find(dao.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var bevents []bson.M
	if err = cursor.All(dao.ctx, &bevents); err != nil {
		return nil, err
	}

	var events []models.Event

	for _, bevent := range bevents {
		var event models.Event

		bsonBytes, _ := bson.Marshal(bevent)
		if err = bson.Unmarshal(bsonBytes, &event); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// GetById gets the event by id.
func (dao *EventDAO) GetById(id int) (*models.Event, error) {
	var (
		err   error
		event models.Event
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"id": id}).Decode(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

// GetByName gets the event by name.
func (dao *EventDAO) GetByName(name string) (*models.Event, error) {
	var (
		err   error
		event models.Event
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"name": name}).Decode(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

// Update updates the event.
func (dao *EventDAO) Update(event models.Event) error {
	var (
		err          error
		updateResult *mongo.UpdateResult
	)

	if updateResult, err = dao.col.UpdateOne(dao.ctx, bson.M{"id": event.Id}, bson.M{"$set": event}); err != nil {
		return err
	}

	// Check to see if the update was successful.
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("update failed")
	}

	// The update was successful.

	return nil
}

// Delete deletes the event.
func (dao *EventDAO) Delete(event models.Event) error {
	return dao.DeleteById(event.Id)
}

// DeleteByName deletes the event by name.
func (dao *EventDAO) DeleteByName(name string) error {
	var (
		err          error
		deleteResult *mongo.DeleteResult
	)

	if deleteResult, err = dao.col.DeleteOne(dao.ctx, bson.M{"name": name}); err != nil {
		return err
	}

	// Check to see if the delete was successful.
	if deleteResult.DeletedCount != 1 {
		return fmt.Errorf("delete failed")
	}

	// The delete was successful.

	return nil
}

// DeleteById deletes the event by id.
func (dao *EventDAO) DeleteById(id int) error {
	var (
		err          error
		deleteResult *mongo.DeleteResult
	)

	if deleteResult, err = dao.col.DeleteOne(dao.ctx, bson.M{"id": id}); err != nil {
		return err
	}

	// Check to see if the delete was successful.
	if deleteResult.DeletedCount != 1 {
		return fmt.Errorf("delete failed")
	}

	// The delete was successful.

	return nil
}

// Create creates the event.
func (dao *EventDAO) Create(event models.Event) error {
	var err error

	if _, err = dao.col.InsertOne(dao.ctx, event); err != nil {
		return err
	}

	return nil
}

// Exists checks to see if the event exists.
func (dao *EventDAO) Exists(event models.Event) (bool, error) {
	return dao.ExistsById(event.Id)
}

// ExistsByName checks to see if the event exists by name.
func (dao *EventDAO) ExistsByName(name string) (bool, error) {
	var (
		err   error
		event models.Event
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"name": name}).Decode(&event); err != nil {
		return false, err
	}

	return true, nil
}

// ExistsById checks to see if the event exists by id.
func (dao *EventDAO) ExistsById(id int) (bool, error) {
	var (
		err   error
		event models.Event
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"id": id}).Decode(&event); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Sync syncs the event.
func (dao *EventDAO) Sync(event models.Event) error {
	var (
		err    error
		answer bool
	)

	if answer, err = dao.Exists(event); err != nil {
		return err
	}

	if !answer {
		// The event doesn't exist.
		log.Printf("inserting event '%s' ...", event.Name)
		if _, err = dao.col.InsertOne(dao.ctx, event); err != nil {
			return err
		}

		log.Printf("event insert successful: %v", event)
		return nil
	}

	// The event exists.
	log.Printf("updating event '%s' ...", event.Name)
	if err = dao.Update(event); err != nil {
		return err
	}

	log.Printf("event update successful: %v", event)
	return nil
}

// SyncAll syncs all events.
func (dao *EventDAO) SyncAll(events []models.Event) error {
	var err error

	for _, event := range events {
		if err = dao.Sync(event); err != nil {
			return err
		}
	}

	return nil
}
