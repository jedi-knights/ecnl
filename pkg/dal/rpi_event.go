package dal

import (
	"context"
	"github.com/jedi-knights/ecnl/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RPIEventDAOer interface {
	Index() error
	GetById(id int) (*models.RPIEvent, error)
	GetByTeamId(teamId int) ([]models.RPIEvent, error)
	GetByTeamName(teamName string) ([]models.RPIEvent, error)
	Create(rpiEvent models.RPIEvent) error
	Exists(rpiEvent models.RPIEvent) (bool, error)
	ExistsById(id primitive.ObjectID) (bool, error)
	DeleteByTeamId(teamId int) error
	DeleteByTeamName(teamName string) error
	Update(id int, event models.RPIEvent) error
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
	panic("implement me")
}

// GetById gets an RPI event by id.
func (dao *RPIEventDAO) GetById(id int) (*models.RPIEvent, error) {
	panic("implement me")
}

// GetByTeamId gets an RPI event by team id.
func (dao *RPIEventDAO) GetByTeamId(teamId int) ([]models.RPIEvent, error) {
	panic("implement me")
}

// GetByTeamName gets an RPI event by team name.
func (dao *RPIEventDAO) GetByTeamName(teamName string) ([]models.RPIEvent, error) {
	panic("implement me")
}

// Create creates an RPI event.
func (dao *RPIEventDAO) Create(rpiEvent models.RPIEvent) error {
	panic("implement me")
}

// Exists checks if an RPI event exists.
func (dao *RPIEventDAO) Exists(rpiEvent models.RPIEvent) (bool, error) {
	panic("implement me")
}

// DeleteByTeamId deletes an RPI event by team id.
func (dao *RPIEventDAO) DeleteByTeamId(teamId int) error {
	panic("implement me")
}

// DeleteByTeamName deletes an RPI event by team name.
func (dao *RPIEventDAO) DeleteByTeamName(teamName string) error {
	panic("implement me")
}

// Update updates an RPI event.
func (dao *RPIEventDAO) Update(id int, event models.RPIEvent) error {
	panic("implement me")
}
