package dal

import "github.com/jedi-knights/ecnl/pkg/models"

type NameMappingDAOer interface {
	GetAll() ([]models.NameMapping, error)
}

type NameMappingDAO struct {
	url string
}

// GetAll gets all name mappings from the specified URL.
func NewNameMappingDAO(url string) *NameMappingDAO {
	return &NameMappingDAO{url: url}
}

// GetAll gets all name mappings from the specified URL.
func (dao *NameMappingDAO) GetAll() ([]models.NameMapping, error) {
	var (
		err  error
		data []models.NameMapping
	)

	if data, err = models.GetNameMappings(dao.url); err != nil {
		return nil, err
	}

	return data, nil
}
