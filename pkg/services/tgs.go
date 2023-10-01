package services

import (
	"encoding/json"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg"
	"github.com/jedi-knights/ecnl/pkg/models"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var organizationClubMap map[int][]models.Club

type TGSServicer interface {
	Url() string
	Client() *http.Client
	Organizations(ecnlOnly bool) ([]models.Organization, error)
	OrganizationByName(name string) (*models.Organization, error)
	OrganizationById(id int) (*models.Organization, error)
}

type TGSService struct {
	HttpClient *http.Client
}

func NewTGSService() *TGSService {
	pHttpClient := &http.Client{}
	return NewTGSServiceWithClient(pHttpClient)
}

func NewTGSServiceWithClient(pHttpClient *http.Client) *TGSService {
	return &TGSService{
		HttpClient: pHttpClient,
	}
}

func (s *TGSService) Url() string {
	return pkg.TgsPrefix
}

func (s *TGSService) Client() *http.Client {
	return s.HttpClient
}

func (s *TGSService) States() ([]models.State, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result string         `json:"result"`
		States []models.State `json:"data"`
	}

	if targetUrl, err = url.JoinPath(s.Url(), "/api/Association", "/get-all-states"); err != nil {
		return nil, err
	}

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, fmt.Errorf("Error getting countries: %v\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			panic("an error occurred while attempting to close the body")
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d: ", pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("Error reading response body: %v\n", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("Error unmarshalling response body: %v\n", err)
	}

	return output.States, nil
}

func (s *TGSService) Countries() ([]models.Country, error) {
	var (
		err       error
		data      []byte
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result    string           `json:"result"`
			Countries []models.Country `json:"data"`
		}
	)

	if targetUrl, err = url.JoinPath(s.Url(), "/api/Association", "/get-all-countries"); err != nil {
		return nil, err
	}

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, fmt.Errorf("Error getting countries: %v\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			panic("an error occured while attempting to close the body")
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d: ", pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("Error reading response body: %v\n", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("Error unmarshalling response body: %v\n", err)
	}

	return output.Countries, nil
}

// Organizations returns the list of organizations for the current season.
// If ecnlOnly is true, then only ECNL organizations are returned.
// If ecnlOnly is false, then all organizations are returned.
// This function accesses the /api/Association/get-current-orgs-list endpoint.
// Each organization has a name and id, a season id, and a season group id.
func (s *TGSService) Organizations(ecnlOnly bool) ([]models.Organization, error) {
	var (
		err                  error
		data                 []byte
		targetUrl            string
		pResponse            *http.Response
		currentOrganizations []models.Organization
		output               struct {
			Result        string                `json:"result"`
			Organizations []models.Organization `json:"data"`
		}
	)

	if targetUrl, err = url.JoinPath(s.Url(), "/api/Association", "/get-current-orgs-list"); err != nil {
		return nil, err
	}

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, fmt.Errorf("Error getting current organizations: %v\n", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			panic("an error occurred while attempting to close the body")
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d: ", pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("Error reading response body: %v\n", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("Error unmarshalling response body: %v\n", err)
	}

	ecnlOrganizations := make([]models.Organization, 0)

	for _, org := range output.Organizations {
		if ecnlOnly {
			if strings.Contains(org.Name, "ECNL") {
				ecnlOrganizations = append(ecnlOrganizations, org)
			}
		} else {
			ecnlOrganizations = append(ecnlOrganizations, org)
		}
	}

	currentOrganizations = ecnlOrganizations

	return currentOrganizations, nil
}

func (s *TGSService) OrganizationByName(name string) (*models.Organization, error) {
	var (
		err                  error
		currentOrganizations []models.Organization
	)

	if currentOrganizations == nil {
		if _, err = s.Organizations(false); err != nil {
			return nil, err
		}
	}

	for _, org := range currentOrganizations {
		if org.Name == name {
			return &org, nil
		}
	}

	return nil, fmt.Errorf("organization name %s not found", name)
}

func (s *TGSService) OrganizationById(id int) (*models.Organization, error) {
	var (
		err                  error
		currentOrganizations []models.Organization
	)

	if currentOrganizations == nil {
		if currentOrganizations, err = s.Organizations(false); err != nil {
			return nil, err
		}
	}

	for _, org := range currentOrganizations {
		if org.Id == id {
			return &org, nil
		}
	}

	return nil, fmt.Errorf("organization id %d not found", id)
}

func (s *TGSService) ClubsByOrganization(org models.Organization) ([]models.Club, error) {
	return s.ClubsByOrganizationId(org.Id)
}

func (s *TGSService) ClubsByOrganizationId(orgId int) ([]models.Club, error) {
	var (
		err       error
		data      []byte
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result string        `json:"result"`
			Clubs  []models.Club `json:"data"`
		}
	)

	if organizationClubMap == nil {
		organizationClubMap = make(map[int][]models.Club)
	}

	if clubs, ok := organizationClubMap[orgId]; ok {
		return clubs, nil
	}

	suffix := fmt.Sprintf("/get-org-club-list/%d", orgId)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Association", suffix); err != nil {
		return nil, err
	}

	if pResponse, err = s.Client().Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting clubs for org %d: %v", orgId, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting clubs for org %d: invalid status code %d", orgId, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	organizationClubMap[orgId] = output.Clubs

	return output.Clubs, nil
}

func (s *TGSService) ClubsByOrganizationName(orgName string) ([]models.Club, error) {
	var (
		err error
		org *models.Organization
	)

	if org, err = s.OrganizationByName(orgName); err != nil {
		return nil, err
	}

	return s.ClubsByOrganization(*org)
}

func (s *TGSService) OrganizationForClubName(clubName string) (*models.Organization, error) {
	var (
		err  error
		orgs []models.Organization
	)

	if orgs, err = s.Organizations(false); err != nil {
		return nil, err
	}

	for _, org := range orgs {
		var clubs []models.Club

		if clubs, err = s.ClubsByOrganization(org); err != nil {
			return nil, err
		}

		for _, club := range clubs {
			if club.Name == clubName {
				return &org, nil
			}
		}
	}

	return nil, fmt.Errorf("organization for club '%s' not found", clubName)
}

func (s *TGSService) ClubByName(clubName string) (*models.Club, error) {
	var (
		err  error
		orgs []models.Organization
	)

	if orgs, err = s.Organizations(false); err != nil {
		return nil, err
	}

	for _, org := range orgs {
		var clubs []models.Club

		if clubs, err = s.ClubsByOrganization(org); err != nil {
			return nil, err
		}

		for _, club := range clubs {
			if club.Name == clubName {
				return &club, nil
			}
		}
	}

	return nil, fmt.Errorf("club '%s' not found", clubName)
}

func (s *TGSService) EventIdsByOrgName(orgName string) ([]int, error) {
	var (
		err      error
		clubs    []models.Club
		eventIds []int
	)

	if clubs, err = s.ClubsByOrganizationName(orgName); err != nil {
		return nil, err
	}

	for _, club := range clubs {
		eventIds = append(eventIds, club.EventId)
	}

	return eventIds, nil
}

func (s *TGSService) EventById(eventId int) (*models.Event, error) {
	var (
		err       error
		data      []byte
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result string       `json:"result"`
			Event  models.Event `json:"data"`
		}
	)

	suffix := fmt.Sprintf("/get-org-event-by-eventID/%d", eventId)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Event", suffix); err != nil {
		return nil, err
	}

	if pResponse, err = s.Client().Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting event %d: %v", eventId, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting event %d: invalid status code %d", eventId, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%s", err, string(data))
	}

	return &output.Event, nil
}

func (s *TGSService) EventsByOrgName(orgName string) ([]models.Event, error) {
	var err error
	var eventIds []int
	var events []models.Event

	if eventIds, err = s.EventIdsByOrgName(orgName); err != nil {
		return nil, err
	}

	for _, eventId := range eventIds {
		var pEvent *models.Event

		if pEvent, err = s.EventById(eventId); err != nil {
			return nil, err
		}

		events = append(events, *pEvent)
	}

	return events, nil
}

func (s *TGSService) EventTypes() ([]models.EventType, error) {
	var (
		err       error
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result     string             `json:"result"`
			EventTypes []models.EventType `json:"data"`
		}
	)

	if targetUrl, err = url.JoinPath(s.Url(), "/api/Mobile", "get-event-types"); err != nil {
		return nil, err
	}

	fmt.Printf("GET %s\n", targetUrl)

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, err
	}
	defer pResponse.Body.Close()

	if err = json.NewDecoder(pResponse.Body).Decode(&output); err != nil {
		return nil, err
	}

	return output.EventTypes, nil
}

func (s *TGSService) EventTypeByName(name string) (*models.EventType, error) {
	var (
		err        error
		eventTypes []models.EventType
	)

	if eventTypes, err = s.EventTypes(); err != nil {
		return nil, err
	}

	for _, eventType := range eventTypes {
		if eventType.Name == name {
			return &eventType, nil
		}
	}

	return nil, fmt.Errorf("event type name %s not found", name)
}

func (s *TGSService) EventTypeById(eventTypeId int) (*models.EventType, error) {
	var (
		err        error
		eventTypes []models.EventType
	)

	if eventTypes, err = s.EventTypes(); err != nil {
		return nil, err
	}

	for _, eventType := range eventTypes {
		if eventType.Id == eventTypeId {
			return &eventType, nil
		}
	}

	return nil, fmt.Errorf("event type eventTypeId %d not found", eventTypeId)
}
