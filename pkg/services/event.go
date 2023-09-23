package services

import (
	"encoding/json"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg"
	"github.com/jedi-knights/ecnl/pkg/models"
	"io"
	"net/http"
	"net/url"
	"slices"
	"sort"
)

var organizationClubMap map[int][]models.Club

type EventService struct {
	pHttpClient         *http.Client
	pAssociationService *AssociationService
}

func NewEventService(pAssociationService *AssociationService) *EventService {
	pHttpClient := &http.Client{}
	return NewEventServiceWithClient(pAssociationService, pHttpClient)
}

func NewEventServiceWithClient(pAssociationService *AssociationService, pHttpClient *http.Client) *EventService {
	return &EventService{
		pHttpClient:         pHttpClient,
		pAssociationService: pAssociationService,
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

func (s *EventService) GetClubsByOrganizationName(name string) ([]models.Club, error) {
	var err error
	var pOrganization *models.Organization

	if pOrganization, err = s.pAssociationService.GetOrganizationByName(name); err != nil {
		return nil, err
	}

	return s.GetClubsByOrganization(pOrganization)
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

	if organizationClubMap == nil {
		organizationClubMap = make(map[int][]models.Club)
	}

	if clubs, ok := organizationClubMap[organizationId]; ok {
		return clubs, nil
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

	organizationClubMap[organizationId] = output.Clubs

	return output.Clubs, nil
}

func (s *EventService) GetClub(orgName string, clubName string) (*models.Club, error) {
	var err error
	var clubs []models.Club

	if clubs, err = s.GetClubsByOrganizationName(orgName); err != nil {
		return nil, err
	}

	for _, club := range clubs {
		if club.Name == clubName {
			return &club, nil
		}
	}

	return nil, fmt.Errorf("club %s not found", clubName)
}

func (s *EventService) GetEventsByOrgName(orgName string) ([]models.Event, error) {
	var err error
	var eventIds []int
	var events []models.Event

	if eventIds, err = s.GetEventIdsByOrgName(orgName); err != nil {
		return nil, err
	}

	for _, eventId := range eventIds {
		var pEvent *models.Event

		if pEvent, err = s.GetEventById(eventId); err != nil {
			return nil, err
		}

		events = append(events, *pEvent)
	}

	return events, nil
}

func (s *EventService) GetEventIdsByOrgName(orgName string) ([]int, error) {
	var err error
	var clubs []models.Club
	var eventIds []int

	if clubs, err = s.GetClubsByOrganizationName(orgName); err != nil {
		return nil, err
	}

	for _, club := range clubs {
		if club.EventId == 0 {
			// What the heck does it mean for a club to have an event id of 0?
			// I found that all the other clubs have non-zero event ids, so I'm
			// a bit confused by this case and am just skipping it for now.

			fmt.Printf("Warning: club '%s' has event id of 0\n", club.Name)
			continue
		}

		if !slices.Contains(eventIds, club.EventId) {
			eventIds = append(eventIds, club.EventId)
		}
	}

	// sort the event ids in ascending order for consistency
	sort.Slice(eventIds, func(i, j int) bool {
		return eventIds[i] < eventIds[j]
	})

	return eventIds, nil
}

func (s *EventService) GetEventById(id int) (*models.Event, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result string       `json:"result"`
		Event  models.Event `json:"data"`
	}

	suffix := fmt.Sprintf("/get-org-event-by-eventID/%d", id)
	if targetUrl, err = s.GetUrl(suffix); err != nil {
		return nil, err
	}

	if pResponse, err = s.pHttpClient.Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting event %d: %v", id, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting event %d: invalid status code %d", id, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%s", err, string(data))
	}

	return &output.Event, nil
}
