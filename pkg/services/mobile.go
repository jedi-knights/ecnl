package services

import (
	"encoding/json"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg"
	"github.com/jedi-knights/ecnl/pkg/models"
	"net/http"
	"net/url"
)

var eventTypes []models.EventType

type MobileService struct {
	HttpClient *http.Client
}

func NewMobileService() *MobileService {
	pHttpClient := &http.Client{}
	return NewMobileServiceWithClient(pHttpClient)
}

func NewMobileServiceWithClient(httpClient *http.Client) *MobileService {
	return &MobileService{HttpClient: httpClient}
}

// getUrl returns the target URL for the given path
func (s *MobileService) getUrl(path string) (string, error) {
	var err error
	var targetUrl string

	if targetUrl, err = url.JoinPath(pkg.TgsPrefix, "/api/Mobile", path); err != nil {
		return "", err
	}

	return targetUrl, nil
}

// GetEventTypes returns all event types
func (s *MobileService) GetEventTypes() ([]models.EventType, error) {
	var err error
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result     string             `json:"result"`
		EventTypes []models.EventType `json:"data"`
	}

	if eventTypes != nil {
		return eventTypes, nil
	}

	if targetUrl, err = s.getUrl("get-event-types"); err != nil {
		return nil, err
	}

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, err
	}
	defer pResponse.Body.Close()

	if err = json.NewDecoder(pResponse.Body).Decode(&output); err != nil {
		return nil, err
	}

	return output.EventTypes, nil
}

// GetEventTypeByName returns the event type with the given name
func (s *MobileService) GetEventTypeByName(name string) (*models.EventType, error) {
	var err error

	if eventTypes == nil {
		if eventTypes, err = s.GetEventTypes(); err != nil {
			return nil, err
		}
	}

	for _, eventType := range eventTypes {
		if eventType.Name == name {
			return &eventType, nil
		}
	}

	return nil, fmt.Errorf("event type name %s not found", name)
}

func (s *MobileService) GetEventTypeById(id int) (*models.EventType, error) {
	if eventTypes == nil {
		if _, err := s.GetEventTypes(); err != nil {
			return nil, err
		}
	}

	for _, eventType := range eventTypes {
		if eventType.Id == id {
			return &eventType, nil
		}
	}

	return nil, fmt.Errorf("event type id %d not found", id)
}
