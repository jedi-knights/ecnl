package dal

import (
	"context"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type TeamDAOer interface {
	Index() error
	GetAll() ([]models.Team, error)
	GetByName(name string) (*models.Team, error)
	GetById(id int) (*models.Team, error)
	Update(team models.Team) error
	Delete(team models.Team) error
	DeleteByName(name string) error
	DeleteById(id int) error
	Create(team models.Team) error
	Exists(team models.Team) (bool, error)
	ExistsByName(name string) (bool, error)
	ExistsById(id int) (bool, error)
	Sync(team models.Team) error
	SyncAll(teams []*models.Team) error
	AppendRPIRanking(timestamp time.Time, data models.RPIRankingData) error
}

// TeamDAO is the data access object for teams.
type TeamDAO struct {
	ctx context.Context
	col *mongo.Collection
}

// NewTeamDAO creates a new team data access object.
func NewTeamDAO(ctx context.Context, col *mongo.Collection) *TeamDAO {
	return &TeamDAO{ctx: ctx, col: col}
}

/*
{
  "result": "success",
  "data": {
    "teamList": [
      {
        "teamID": 58734,
        "teamName": "Alabama FC ECNL G09",
        "status": 2,
        "clubID": 18,
        "initialSeed": 1,
        "clubLogo": "https://s3.amazonaws.com/images.totalglobalsports.com/PlayerForm/_da463c0400e545cbbfb1f23ee81d0cde_afc.png",
        "firstName": "Thomas",
        "lastName": "Brower",
        "wdl": "32W 10D 16L",
        "flightRequested": null,
        "currentFlightID": 0,
        "currentFlight": "ECNL",
        "currentFlightIDString": "19789"
      },
...
*/

// Index indexes the collection.
func (dao *TeamDAO) Index() error {
	var (
		err  error
		name string
	)

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, mongo.IndexModel{
		Keys: bson.D{{"id", 1}},
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s on teams collection", name)
	}

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, mongo.IndexModel{
		Keys: bson.D{{"name", 1}},
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s on teams collection", name)
	}

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, mongo.IndexModel{
		Keys: bson.D{{"clubid", 1}},
	}); err != nil {
		return err
	} else {
		log.Printf("created index %s on teams collection", name)
	}

	return nil
}

// GetAll gets all teams.
func (dao *TeamDAO) GetAll() ([]models.Team, error) {
	var (
		cursor *mongo.Cursor
		err    error
	)

	if cursor, err = dao.col.Find(dao.ctx, bson.M{}); err != nil {
		return nil, err
	}

	var bteams []bson.M
	if err = cursor.All(dao.ctx, &bteams); err != nil {
		return nil, err
	}

	var teams []models.Team

	for _, bteam := range bteams {
		var team models.Team

		bsonBytes, _ := bson.Marshal(bteam)
		if err = bson.Unmarshal(bsonBytes, &team); err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	return teams, nil
}

// GetByName gets a team by name.
func (dao *TeamDAO) GetByName(name string) (*models.Team, error) {
	var (
		err   error
		bteam bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"name": name}).Decode(&bteam); err != nil {
		return nil, err
	}

	var team models.Team

	bsonBytes, _ := bson.Marshal(bteam)
	if err = bson.Unmarshal(bsonBytes, &team); err != nil {
		return nil, err
	}

	return &team, nil
}

// GetById gets a team by id.
func (dao *TeamDAO) GetById(id int) (*models.Team, error) {
	var (
		err   error
		bteam bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"id": id}).Decode(&bteam); err != nil {
		return nil, err
	}

	var team models.Team

	bsonBytes, _ := bson.Marshal(bteam)
	if err = bson.Unmarshal(bsonBytes, &team); err != nil {
		return nil, err
	}

	return &team, nil
}

// Update updates a team.
func (dao *TeamDAO) Update(team models.Team) error {
	var (
		err          error
		updateResult *mongo.UpdateResult
	)

	log.Println("updating team: ", team.String())

	if updateResult, err = dao.col.UpdateOne(dao.ctx, bson.M{"id": team.Id}, bson.M{"$set": team}); err != nil {
		return err
	}

	// Check to see if the update was successful.
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("update failed")
	}

	// The update was successful.
	log.Printf("team update successful: %v", updateResult)

	return nil
}

// Delete deletes a team.
func (dao *TeamDAO) Delete(team models.Team) error {
	return dao.DeleteById(team.Id)
}

// DeleteByName deletes a team by name.
func (dao *TeamDAO) DeleteByName(name string) error {
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
	log.Printf("team delete successful: %v", deleteResult)

	return nil
}

// DeleteById deletes a team by id.
func (dao *TeamDAO) DeleteById(id int) error {
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
	log.Printf("team delete successful: %v", deleteResult)

	return nil
}

// Create creates a team.
func (dao *TeamDAO) Create(team models.Team) error {
	var (
		err error
	)

	if _, err = dao.col.InsertOne(dao.ctx, team); err != nil {
		return err
	}

	return nil
}

// Exists checks if a team exists.
func (dao *TeamDAO) Exists(team models.Team) (bool, error) {
	var (
		err   error
		bTeam bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"id": team.Id}).Decode(&bTeam); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// ExistsByName checks if a team exists by name.
func (dao *TeamDAO) ExistsByName(name string) (bool, error) {
	var (
		err   error
		count int64
	)

	if count, err = dao.col.CountDocuments(dao.ctx, bson.M{"name": name}); err != nil {
		return false, err
	}

	return count >= 1, nil
}

// ExistsById checks if a team exists by id.
func (dao *TeamDAO) ExistsById(id int) (bool, error) {
	var (
		err   error
		count int64
	)

	if count, err = dao.col.CountDocuments(dao.ctx, bson.M{"id": id}); err != nil {
		return false, err
	}

	return count >= 1, nil
}

// Sync syncs a team.
func (dao *TeamDAO) Sync(team models.Team) error {
	var (
		err    error
		answer bool
	)

	if answer, err = dao.Exists(team); err != nil {
		return err
	}

	if !answer {
		// Create the team.
		if err = dao.Create(team); err != nil {
			return err
		}

		log.Printf("team '%s' created", team.Name)
	} else {
		// Update the team.
		if err = dao.Update(team); err != nil {
			return err
		}

		log.Printf("team '%s' updated", team.Name)
	}

	return nil
}

// SyncAll syncs all teams.
func (dao *TeamDAO) SyncAll(teams []*models.Team) error {
	for _, team := range teams {
		if err := dao.Sync(*team); err != nil {
			return err
		}
	}

	return nil
}

// AppendRPIRanking appends an RPI ranking to a team.
func (dao *TeamDAO) AppendRPIRanking(timestamp time.Time, data models.RPIRankingData) error {
	var (
		err          error
		updateResult *mongo.UpdateResult
	)

	log.Printf("appending RPI ranking - %s", data.String())

	rpiHistoryItem := models.NewRPIEvent(timestamp, data)

	filter := bson.M{"id": data.TeamId}
	update := bson.M{
		"$push": bson.M{
			"rpi_history": rpiHistoryItem,
		},
	}

	if updateResult, err = dao.col.UpdateOne(dao.ctx, filter, update); err != nil {
		return err
	}

	// Check to see if the update was successful.
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("update failed, matched %d - %s", updateResult.MatchedCount, data.String())
	}

	// The update was successful.
	log.Printf("team update successful: %v", updateResult)

	return nil
}
