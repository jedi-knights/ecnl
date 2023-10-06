package dal

import (
	"context"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// ClubDAOer is the interface for the club data access object.
type ClubDAOer interface {
	Index() error
	GetAll() ([]models.Club, error)
	GetById(id int) (*models.Club, error)
	GetByName(name string) (*models.Club, error)
	Update(club models.Club) error
	Delete(club models.Club) error
	DeleteById(id int) error
	DeleteByName(name string) error
	Create(club models.Club) error
	Exists(club models.Club) (bool, error)
	ExistsById(id int) (bool, error)
	ExistsByName(name string) (bool, error)
	Sync(club models.Club) error
	SyncAll(clubs []models.Club) error
}

// ClubDAO is the data access object for clubs.
type ClubDAO struct {
	ctx context.Context
	col *mongo.Collection
}

// NewClubDAO creates a new club data access object.
func NewClubDAO(ctx context.Context, col *mongo.Collection) *ClubDAO {
	return &ClubDAO{ctx: ctx, col: col}
}

// Index indexes the collection.
func (dao *ClubDAO) Index() error {
	var (
		err  error
		name string
	)

	clubsIndexModel := mongo.IndexModel{
		Keys: bson.D{{"orgid", 1}, {"clubid", 1}, {"name", 1}, {"statecode", 1}},
	}

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, clubsIndexModel); err != nil {
		return err
	}

	log.Printf("created index %s on clubs collection", name)

	return nil
}

// GetAll gets all the clubs.
func (dao *ClubDAO) GetAll() ([]models.Club, error) {
	cursor, err := dao.col.Find(dao.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var bclubs []bson.M
	if err = cursor.All(dao.ctx, &bclubs); err != nil {
		return nil, err
	}

	var clubs []models.Club

	for _, bclub := range bclubs {
		var club models.Club

		bsonBytes, _ := bson.Marshal(bclub)
		if err = bson.Unmarshal(bsonBytes, &club); err != nil {
			return nil, err
		}

		clubs = append(clubs, club)
	}

	return clubs, nil
}

// GetById gets the club by id.
func (dao *ClubDAO) GetById(id int) (*models.Club, error) {
	var (
		err   error
		bclub bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"clubid": id}).Decode(&bclub); err != nil {
		return nil, err
	}

	var club models.Club

	bsonBytes, _ := bson.Marshal(bclub)
	if err = bson.Unmarshal(bsonBytes, &club); err != nil {
		return nil, err
	}

	return &club, nil
}

// GetByName gets the club by name.
func (dao *ClubDAO) GetByName(name string) (*models.Club, error) {
	var (
		err   error
		bclub bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"name": name}).Decode(&bclub); err != nil {
		return nil, err
	}

	var club models.Club

	bsonBytes, _ := bson.Marshal(bclub)
	if err = bson.Unmarshal(bsonBytes, &club); err != nil {
		return nil, err
	}

	return &club, nil
}

// Update updates the club.
func (dao *ClubDAO) Update(club models.Club) error {
	var (
		err          error
		updateResult *mongo.UpdateResult
	)

	if updateResult, err = dao.col.UpdateOne(dao.ctx, bson.M{"clubid": club.ClubId}, bson.M{"$set": club}); err != nil {
		return err
	}

	// Check to see if the update was successful.
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("update failed")
	}

	// The update was successful.
	// log.Printf("club update successful: %v", updateResult)

	return nil
}

// Delete deletes the club.
func (dao *ClubDAO) Delete(club models.Club) error {
	return dao.DeleteById(club.ClubId)
}

// DeleteById deletes the club by id.
func (dao *ClubDAO) DeleteById(id int) error {
	var (
		err          error
		deleteResult *mongo.DeleteResult
	)

	if deleteResult, err = dao.col.DeleteOne(dao.ctx, bson.M{"clubid": id}); err != nil {
		return err
	}

	// Check to see if the delete was successful.
	if deleteResult.DeletedCount != 1 {
		return fmt.Errorf("delete failed")
	}

	// The delete was successful.
	log.Printf("club delete successful: %v", deleteResult)

	return nil
}

// DeleteByName deletes the club by name.
func (dao *ClubDAO) DeleteByName(name string) error {
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
	log.Printf("club delete successful: %v", deleteResult)

	return nil
}

// Create creates a club.
func (dao *ClubDAO) Create(club models.Club) error {
	var err error

	if _, err = dao.col.InsertOne(dao.ctx, club); err != nil {
		return err
	}

	return nil
}

// Exists checks to see if the club exists.
func (dao *ClubDAO) Exists(club models.Club) (bool, error) {
	return dao.ExistsById(club.ClubId)
}

// ExistsById checks to see if the club exists by id.
func (dao *ClubDAO) ExistsById(id int) (bool, error) {
	var (
		err  error
		club *models.Club
	)

	if club, err = dao.GetById(id); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	if club == nil {
		return false, nil
	}

	return true, nil
}

// ExistsByName checks to see if the club exists by name.
func (dao *ClubDAO) ExistsByName(name string) (bool, error) {
	var (
		err  error
		club *models.Club
	)

	if club, err = dao.GetByName(name); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	if club == nil {
		return false, nil
	}

	return true, nil
}

// Sync syncs the club.
func (dao *ClubDAO) Sync(club models.Club) error {
	var (
		err    error
		answer bool
	)

	if answer, err = dao.Exists(club); err != nil {
		return err
	}

	if !answer {
		// The club doesn't exist.
		// log.Printf("inserting club '%s' ...", club.Name)
		if _, err = dao.col.InsertOne(dao.ctx, club); err != nil {
			return err
		}

		log.Printf("club insert successful: %v", club)
		return nil
	}

	// The club exists.
	log.Printf("updating club '%s' ...", club.Name)
	if err = dao.Update(club); err != nil {
		return err
	}

	log.Printf("club update successful: %v", club)

	return nil
}

// SyncAll syncs all clubs.
func (dao *ClubDAO) SyncAll(clubs []models.Club) error {
	for _, club := range clubs {
		if err := dao.Sync(club); err != nil {
			return err
		}
	}

	return nil
}
