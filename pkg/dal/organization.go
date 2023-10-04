package dal

import (
	"context"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type OrganizationDAOer interface {
	Index() error
	GetAll() ([]models.Organization, error)
	GetById(id int) (*models.Organization, error)
	GetByName(name string) (*models.Organization, error)
	Update(organization models.Organization) error
	Delete(organization models.Organization) error
	DeleteByName(name string) error
	DeleteById(id int) error
	Create(organization models.Organization) error
	Exists(organization models.Organization) (bool, error)
	ExistsByName(name string) (bool, error)
	ExistsById(id int) (bool, error)
	Sync(organization models.Organization) error
	SyncAll(organizations []models.Organization) error
}

// OrganizationDAO is the data access object for organizations.
type OrganizationDAO struct {
	ctx context.Context
	col *mongo.Collection
}

// NewOrganizationDAO creates a new organization data access object.
func NewOrganizationDAO(ctx context.Context, col *mongo.Collection) *OrganizationDAO {
	return &OrganizationDAO{ctx: ctx, col: col}
}

// Index indexes the collection.
func (dao *OrganizationDAO) Index() error {
	var (
		name string
		err  error
	)

	organizationsIndexModel := mongo.IndexModel{
		Keys: bson.D{{"id", 1}, {"name", 1}},
	}

	if name, err = dao.col.Indexes().CreateOne(dao.ctx, organizationsIndexModel); err != nil {
		return err
	}

	log.Printf("created index %s on organizations collection", name)

	return nil
}

// GetAll gets all organizations.
func (dao *OrganizationDAO) GetAll() ([]models.Organization, error) {
	var (
		cursor *mongo.Cursor
		err    error
	)

	if cursor, err = dao.col.Find(dao.ctx, bson.M{}); err != nil {
		return nil, err
	}

	var borganizations []bson.M
	if err = cursor.All(dao.ctx, &borganizations); err != nil {
		return nil, err
	}

	var organizations []models.Organization

	for _, borganization := range borganizations {
		var organization models.Organization

		bsonBytes, _ := bson.Marshal(borganization)
		if err = bson.Unmarshal(bsonBytes, &organization); err != nil {
			return nil, err
		}

		organizations = append(organizations, organization)
	}

	return organizations, nil
}

// GetById gets the organization by id.
func (dao *OrganizationDAO) GetById(id int) (*models.Organization, error) {
	var (
		err           error
		borganization bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"id": id}).Decode(&borganization); err != nil {
		return nil, err
	}

	var organization models.Organization

	bsonBytes, _ := bson.Marshal(borganization)
	if err = bson.Unmarshal(bsonBytes, &organization); err != nil {
		return nil, err
	}

	return &organization, nil
}

// GetByName gets the organization by name.
func (dao *OrganizationDAO) GetByName(name string) (*models.Organization, error) {
	var (
		err           error
		borganization bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"name": name}).Decode(&borganization); err != nil {
		return nil, err
	}

	var organization models.Organization

	bsonBytes, _ := bson.Marshal(borganization)
	if err = bson.Unmarshal(bsonBytes, &organization); err != nil {
		return nil, err
	}

	return &organization, nil
}

// Update updates the organization.
func (dao *OrganizationDAO) Update(organization models.Organization) error {
	var (
		err          error
		updateResult *mongo.UpdateResult
	)

	if updateResult, err = dao.col.UpdateOne(dao.ctx, bson.M{"id": organization.Id}, bson.M{"$set": organization}); err != nil {
		return err
	}

	// Check to see if the update was successful.
	if updateResult.MatchedCount != 1 {
		return fmt.Errorf("update failed")
	}

	// The update was successful.
	log.Printf("organization update successful: %v", updateResult)

	return nil
}

// Delete deletes the organization.
func (dao *OrganizationDAO) Delete(organization models.Organization) error {
	return dao.DeleteById(organization.Id)
}

// DeleteByName deletes the organization by name.
func (dao *OrganizationDAO) DeleteByName(name string) error {
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
	log.Printf("organization delete successful: %v", deleteResult)

	return nil
}

// DeleteById deletes the organization by id.
func (dao *OrganizationDAO) DeleteById(id int) error {
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
	log.Printf("organization delete successful: %v", deleteResult)

	return nil
}

// Create creates the organization.
func (dao *OrganizationDAO) Create(organization models.Organization) error {
	var err error

	if _, err = dao.col.InsertOne(dao.ctx, organization); err != nil {
		return err
	}

	return nil
}

// Exists checks to see if the organization exists.
func (dao *OrganizationDAO) Exists(organization models.Organization) (bool, error) {
	return dao.ExistsById(organization.Id)
}

// ExistsByName checks to see if the organization exists by name.
func (dao *OrganizationDAO) ExistsByName(name string) (bool, error) {
	var (
		err         error
		existingOrg bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"name": name}).Decode(&existingOrg); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// ExistsById checks to see if the organization exists by id.
func (dao *OrganizationDAO) ExistsById(id int) (bool, error) {
	var (
		err         error
		existingOrg bson.M
	)

	if err = dao.col.FindOne(dao.ctx, bson.M{"id": id}).Decode(&existingOrg); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Sync syncs the organization.
func (dao *OrganizationDAO) Sync(organization models.Organization) error {
	var (
		err    error
		answer bool
	)

	if answer, err = dao.Exists(organization); err != nil {
		return err
	}

	if !answer {
		// the organziaton does not exist, so create it
		if err = dao.Create(organization); err != nil {
			return err
		}

		log.Printf("organization created: %v", organization)
		return nil
	}

	// the organization exists, so update it
	if err = dao.Update(organization); err != nil {
		return err
	}

	log.Printf("organization updated: %v", organization)

	return nil
}

// SyncAll syncs all organizations in the slice.
func (dao *OrganizationDAO) SyncAll(organizations []models.Organization) error {
	for _, organization := range organizations {
		if err := dao.Sync(organization); err != nil {
			return err
		}
	}

	return nil
}
