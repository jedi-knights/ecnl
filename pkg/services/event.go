package services

import (
	"encoding/json"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg"
	"github.com/jedi-knights/ecnl/pkg/models"
	"io"
	"net/http"
	"net/url"
)

type EventService struct {
	pHttpClient *http.Client
}

func NewEventService() *EventService {
	pHttpClient := &http.Client{}
	return NewEventServiceWithClient(pHttpClient)
}

func NewEventServiceWithClient(pHttpClient *http.Client) *EventService {
	return &EventService{
		pHttpClient: pHttpClient,
	}
}

func (s *EventService) GetUrl(path string) (string, error) {
	var err error
	var targetUrl string

	if targetUrl, err = url.JoinPath(pkg.TgsPrefix, "/api/Event", path); err != nil {
		return "", err
	}

	return targetUrl, nil
}

func (s *EventService) GetClubsByOrganization(pOrganization *models.Organization) ([]models.Club, error) {
	if pOrganization == nil {
		return nil, fmt.Errorf("Organization is nil")
	}

	return s.GetClubsByOrganizationId(pOrganization.Id)
}

func (s *EventService) GetClubsByOrganizationId(organizationId int) ([]models.Club, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result string        `json:"result"`
		Clubs  []models.Club `json:"data"`
	}

	suffix := fmt.Sprintf("/get-org-club-list/%d", organizationId)
	if targetUrl, err = s.GetUrl(suffix); err != nil {
		return nil, err
	}

	if pResponse, err = s.pHttpClient.Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting clubs for org %d: %v", organizationId, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting clubs for org %d: invalid status code %d", organizationId, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return output.Clubs, nil
}
