package services

import (
	"encoding/json"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/jedi-knights/rpi/pkg/match"
	"github.com/jedi-knights/rpi/pkg/schedule"
	"io"
	"log"
	"net/http"
	"net/url"
	"slices"
	"sort"
	"strings"
	"time"
)

var organizationClubMap map[int][]models.Club

type GlobalServicer interface {
	Url() string
	Client() *http.Client
	Organizations(ecnlOnly bool) ([]models.Organization, error)
	OrganizationByName(name string) (*models.Organization, error)
	OrganizationById(id int) (*models.Organization, error)
}

type GlobalService struct {
	HttpClient *http.Client
}

func NewTGSService() *GlobalService {
	pHttpClient := &http.Client{}

	return NewTGSServiceWithClient(pHttpClient)
}

func NewTGSServiceWithClient(pHttpClient *http.Client) *GlobalService {
	return &GlobalService{
		HttpClient: pHttpClient,
	}
}

func (s *GlobalService) Url() string {
	return pkg.TgsPrefix
}

func (s *GlobalService) Client() *http.Client {
	return s.HttpClient
}

func (s *GlobalService) States() ([]models.State, error) {
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
	defer pResponse.Body.Close()

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

func (s *GlobalService) Countries() ([]models.Country, error) {
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

	defer pResponse.Body.Close()

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

var organizationCache []models.Organization

// Organizations returns the list of organizations for the current season.
// If ecnlOnly is true, then only ECNL organizations are returned.
// If ecnlOnly is false, then all organizations are returned.
// This function accesses the /api/Association/get-current-orgs-list endpoint.
// Each organization has a name and id, a season id, and a season group id.
func (s *GlobalService) Organizations(ecnlOnly bool) ([]models.Organization, error) {
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

	if organizationCache != nil {
		return organizationCache, nil
	}

	if targetUrl, err = url.JoinPath(s.Url(), "/api/Association", "/get-current-orgs-list"); err != nil {
		return nil, err
	}

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, fmt.Errorf("Error getting current organizations: %v\n", err)
	}

	defer pResponse.Body.Close()

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

	organizationCache = currentOrganizations

	return currentOrganizations, nil
}

func (s *GlobalService) OrganizationByName(name string) (*models.Organization, error) {
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
		if org.Name == name {
			return &org, nil
		}
	}

	return nil, fmt.Errorf("organization name %svc not found", name)
}

func (s *GlobalService) OrganizationById(id int) (*models.Organization, error) {
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

func (s *GlobalService) ClubsByOrganization(org models.Organization) ([]models.Club, error) {
	return s.ClubsByOrganizationId(org.Id)
}

func (s *GlobalService) ClubsByOrganizationId(orgId int) ([]models.Club, error) {
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

	suffix := fmt.Sprintf("get-org-club-list-by-orgID-improved/%d", orgId)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Event", suffix); err != nil {
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

func (s *GlobalService) ClubsByOrgName(orgName string) ([]models.Club, error) {
	var (
		err error
		org *models.Organization
	)

	if org, err = s.OrganizationByName(orgName); err != nil {
		return nil, err
	}

	return s.ClubsByOrganization(*org)
}

func (s *GlobalService) OrganizationForClubName(clubName string) (*models.Organization, error) {
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

	return nil, fmt.Errorf("organization for club '%svc' not found", clubName)
}

var clubsByNameCache map[string]*models.Club

func (s *GlobalService) ClubByName(clubName string) (*models.Club, error) {
	var (
		err  error
		orgs []models.Organization
	)

	if clubsByNameCache != nil {
		if club, ok := clubsByNameCache[clubName]; ok {
			return club, nil
		}
	}

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
				if clubsByNameCache == nil {
					clubsByNameCache = make(map[string]*models.Club)

					clubsByNameCache[clubName] = &club
				}

				return &club, nil
			}
		}
	}

	return nil, fmt.Errorf("club '%svc' not found", clubName)
}

func (s *GlobalService) EventIdsByOrgId(orgId int) ([]int, error) {
	var (
		err      error
		clubs    []models.Club
		eventIds []int
	)

	if clubs, err = s.ClubsByOrganizationId(orgId); err != nil {
		return nil, err
	}

	for _, club := range clubs {
		if slices.Contains(eventIds, club.EventId) {
			continue
		}

		eventIds = append(eventIds, club.EventId)
	}

	return eventIds, nil
}

func (s *GlobalService) EventIdsByOrgName(orgName string) ([]int, error) {
	var (
		err      error
		clubs    []models.Club
		eventIds []int
	)

	if clubs, err = s.ClubsByOrgName(orgName); err != nil {
		return nil, err
	}

	for _, club := range clubs {
		if slices.Contains(eventIds, club.EventId) {
			continue
		}

		eventIds = append(eventIds, club.EventId)
	}

	sort.Slice(eventIds, func(i, j int) bool {
		return eventIds[i] < eventIds[j]
	})

	return eventIds, nil
}

var eventCache map[int]*models.Event

func (s *GlobalService) EventById(eventId int) (*models.Event, error) {
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

	if eventCache != nil {
		if event, ok := eventCache[eventId]; ok {
			return event, nil
		}
	}

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
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%svc", err, string(data))
	}

	if eventCache == nil {
		eventCache = make(map[int]*models.Event)
	}

	eventCache[eventId] = &output.Event

	return &output.Event, nil
}

func (s *GlobalService) EventByName(eventName string) (*models.Event, error) {
	var (
		err    error
		events []models.Event
	)

	if events, err = s.Events(); err != nil {
		return nil, err
	}

	for _, event := range events {
		if event.Name == eventName {
			return &event, nil
		}
	}

	return nil, fmt.Errorf("event name %s not found", eventName)
}

func (s *GlobalService) Events() ([]models.Event, error) {
	var err error
	var orgs []models.Organization
	var events []models.Event

	if orgs, err = s.Organizations(false); err != nil {
		return nil, err
	}

	for _, org := range orgs {
		var orgEvents []models.Event

		if orgEvents, err = s.EventsByOrgId(org.Id); err != nil {
			return nil, err
		}

		events = append(events, orgEvents...)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Name < events[j].Name
	})

	return events, nil
}

func (s *GlobalService) EventsByOrganization(org models.Organization) ([]models.Event, error) {
	return s.EventsByOrgId(org.Id)
}

var orgEventsCache map[int][]models.Event

func (s *GlobalService) EventsByOrgId(orgId int) ([]models.Event, error) {
	var err error
	var eventIds []int
	var events []models.Event

	if orgEventsCache != nil {
		val, ok := orgEventsCache[orgId]
		if ok {
			return val, nil
		}
	}

	if eventIds, err = s.EventIdsByOrgId(orgId); err != nil {
		return nil, err
	}

	for _, eventId := range eventIds {
		var pEvent *models.Event

		if pEvent, err = s.EventById(eventId); err != nil {
			return nil, err
		}

		if pEvent.Name == "" {
			continue
		}

		events = append(events, *pEvent)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Name < events[j].Name
	})

	if orgEventsCache == nil {
		orgEventsCache = make(map[int][]models.Event)
	}

	orgEventsCache[orgId] = events

	return events, nil
}

func (s *GlobalService) EventsByOrgName(orgName string) ([]models.Event, error) {
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

		if pEvent.Name == "" {
			continue
		}

		events = append(events, *pEvent)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Name < events[j].Name
	})

	return events, nil
}

func (s *GlobalService) EventTypes() ([]models.EventType, error) {
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

	fmt.Printf("GET %svc\n", targetUrl)

	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if err = json.NewDecoder(pResponse.Body).Decode(&output); err != nil {
		return nil, err
	}

	return output.EventTypes, nil
}

func (s *GlobalService) EventTypeByName(name string) (*models.EventType, error) {
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

	return nil, fmt.Errorf("event type name %svc not found", name)
}

func (s *GlobalService) EventTypeById(eventTypeId int) (*models.EventType, error) {
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

/*
/api/Event/get-org-events-standings-by-id/{orgID}/{orgSeasonGroupID}/{eventID}/{divisionID}
/api/Event/get-org-events-team-List-by-id/{orgID}/{orgSeasonGroupID}/{eventID}/{divisionID}
*/

func (s *GlobalService) OrganizationDivisionsByOrgId(orgId, eventId int) ([]models.OrganiationDivision, error) {
	var (
		err       error
		data      []byte
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result string `json:"result"`
			Data   struct {
				OrgId        int                          `json:"orgID"`
				OrgSeasonId  int                          `json:"orgSeasonID"`
				DivisionList []models.OrganiationDivision `json:"divisionList"`
			} `json:"data"`
		}
	)

	suffix := fmt.Sprintf("/get-org-division-list/%d/%d", orgId, eventId)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Event", suffix); err != nil {
		return nil, err
	}

	fmt.Printf("GET %svc\n", targetUrl)

	if pResponse, err = s.Client().Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting organization divisoins for orgId %d eventId %d: %v", orgId, eventId, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting organization divisoins for orgId %d eventId %d: invalid status code %d", orgId, eventId, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%svc", err, string(data))
	}

	return output.Data.DivisionList, nil
}

// MatchResults returns the match results by club name and event name. (e.g. "Concorde Fire Premier" and "ECNL Girls")
// Keep in mind the results are across all age groups so they still need to be filtered.
func (s *GlobalService) MatchEventsByClubNameAndEventName(clubName string, eventName string) ([]models.MatchEvent, error) {
	var (
		err       error
		data      []byte
		targetUrl string
		pResponse *http.Response
		event     *models.Event
		club      *models.Club
		output    struct {
			Result string `json:"result"`
			Data   struct {
				ReportedScore   int                 `json:"reportedScore"`
				UnReportedScore int                 `json:"unReportedScore"`
				MatchEvents     []models.MatchEvent `json:"eventPastScheduleList"`
			} `json:"data"`
		}
	)

	if club, err = s.ClubByName(clubName); err != nil {
		return nil, err
	}

	if event, err = s.EventByName(eventName); err != nil {
		return nil, err
	}

	suffix := fmt.Sprintf("/get-score-reporting-schedule-list/%d/%d", club.ClubId, event.Id)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Club", suffix); err != nil {
		return nil, err
	}

	// fmt.Printf("GET %s\n", targetUrl)

	if pResponse, err = s.Client().Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting schedule for %s club %s: %v", eventName, clubName, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting schedule for %s club %s: invalid status code %d", eventName, clubName, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%svc", err, string(data))
	}

	return output.Data.MatchEvents, nil
}

// ClubsByEvent returns all of the clubs for the given event.
func (s *GlobalService) ClubsByEvent(event models.Event) ([]models.Club, error) {
	var (
		err   error
		clubs []models.Club
		orgs  []models.Organization
	)

	if orgs, err = s.Organizations(false); err != nil {
		return nil, err
	}

	for _, org := range orgs {
		if clubs, err = s.ClubsByOrganization(org); err != nil {
			return nil, err
		}

		for _, club := range clubs {
			if club.EventId != event.Id {
				continue
			}

			clubs = append(clubs, club)
		}
	}

	return clubs, nil
}

func (s *GlobalService) DivisionsByEvent(event models.Event) ([]models.Division, error) {
	var (
		err       error
		data      []byte
		gender    string
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result    string            `json:"result"`
			Divisions []models.Division `json:"data"`
		}
	)

	if strings.Contains(event.Name, "Boys") {
		gender = "M"
	} else if strings.Contains(event.Name, "Girls") {
		gender = "F"
	} else {
		log.Printf("can't determine gender for event '%s' assuming it's male", event.String())

		gender = "M"
	}

	// /api/Club/get-event-divisions-by-event-and-gender/{eventID}/{gender}
	suffix := fmt.Sprintf("/get-event-divisions-by-event-and-gender/%d/%s", event.Id, gender)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Club", suffix); err != nil {
		return nil, err
	}

	if pResponse, err = s.Client().Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting divisions for event %s: %v", event.Name, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting divisions for event %s: invalid status code %d", event.Name, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%svc", err, string(data))
	}

	return output.Divisions, nil
}

func (s *GlobalService) DivisionsByEventId(id int) ([]models.Division, error) {
	var (
		err       error
		event     *models.Event
		divisions []models.Division
	)

	if event, err = s.EventById(id); err != nil {
		return nil, err
	}

	if divisions, err = s.DivisionsByEvent(*event); err != nil {
		return nil, err
	}

	return divisions, nil
}

func (s *GlobalService) DivisionsByEventName(name string) ([]models.Division, error) {
	var (
		err       error
		event     *models.Event
		divisions []models.Division
	)

	if event, err = s.EventByName(name); err != nil {
		return nil, err
	}

	if divisions, err = s.DivisionsByEvent(*event); err != nil {
		return nil, err
	}

	return divisions, nil
}

// How can I get the age groups?

// MatchResultsBayAgeGroup returns all of the match results filtered by age group.
// Note: ageGroup takes the form "G2009" for example. (it maps to the Division property)
func (s *GlobalService) MatchEventsByAgeGroup(clubName, eventName, ageGroup string) ([]models.MatchEvent, error) {
	var (
		err                  error
		matchResults         []models.MatchEvent
		filteredMatchResults []models.MatchEvent
	)

	if matchResults, err = s.MatchEventsByClubNameAndEventName(clubName, eventName); err != nil {
		return nil, err
	}

	for _, matchResult := range matchResults {
		if matchResult.Division == ageGroup {
			filteredMatchResults = append(filteredMatchResults, matchResult)
		}
	}

	return filteredMatchResults, nil
}

// I need something that returns all the match data so that I can save it in a collection.

func MatchEventsByClub(club models.Club) []models.MatchEvent {
	return nil
}

// RPISchedule returns all of the match results for the given organization name, event name, and age group.
// Example: MatchResults("ECNL Girls", "G2009
// Warning: This is going to make a lot of API calls!
func (s *GlobalService) RPISchedule(orgName, ageGrouop string) (*schedule.Schedule, []string, error) {
	var (
		err         error
		events      []models.Event
		clubs       []models.Club
		rpiSchedule *schedule.Schedule
		teamNames   []string
	)

	if events, err = s.EventsByOrgName(orgName); err != nil {
		return nil, nil, err
	}

	if clubs, err = s.ClubsByOrgName(orgName); err != nil {
		return nil, nil, err
	}

	rpiSchedule = schedule.NewSchedule()

	for _, event := range events {
		fmt.Println(event.String())
		for _, club := range clubs {
			fmt.Println("\t" + club.String())
			if club.EventId != event.Id {
				continue
			}

			var matchResults []models.MatchEvent

			if matchResults, err = s.MatchEventsByAgeGroup(club.Name, event.Name, ageGrouop); err != nil {
				return nil, nil, err
			}

			for _, matchResult := range matchResults {
				var parsedTime time.Time

				if parsedTime, err = time.Parse("2006-01-02T15:04:05", matchResult.GameDate); err != nil {
					// If the date parsing goes wonky then just use the current time.
					parsedTime = time.Now()
				}

				currentMatch := match.NewMatch()
				currentMatch.Date = parsedTime
				currentMatch.Home.Name = matchResult.HomeTeamName
				currentMatch.Home.Score = matchResult.HomeTeamScore
				currentMatch.Away.Name = matchResult.AwayTeamName
				currentMatch.Away.Score = matchResult.AwayTeamScore

				fmt.Printf("\t\t%s\n", currentMatch.ToString())

				rpiSchedule.AddMatch(currentMatch)

				if !slices.Contains(teamNames, currentMatch.Home.Name) {
					teamNames = append(teamNames, currentMatch.Home.Name)
				}

				if !slices.Contains(teamNames, currentMatch.Away.Name) {
					teamNames = append(teamNames, currentMatch.Away.Name)
				}
			}
		}
	}

	// sort the returned team names
	sort.Slice(teamNames, func(i, j int) bool {
		return teamNames[i] < teamNames[j]
	})

	return rpiSchedule, teamNames, nil
}

// TeamsByEventId returns all of the teams for the given event.
func (s *GlobalService) TeamsByEventId(eventId int) ([]*models.Team, error) {
	var (
		err           error
		divisions     []models.Division
		teams         []*models.Team
		divisionTeams []*models.Team
	)

	if divisions, err = s.DivisionsByEventId(eventId); err != nil {
		return nil, err
	}

	for _, division := range divisions {
		if divisionTeams, err = s.TeamsByEventIdAndDivisionId(eventId, division.Id); err != nil {
			return nil, err
		}

		teams = append(teams, divisionTeams...)
	}

	return teams, nil
}

// TeamsByEvent returns all of the teams for the given event.
func (s *GlobalService) TeamsByEvent(event models.Event) ([]*models.Team, error) {
	var (
		err           error
		divisions     []models.Division
		teams         []*models.Team
		divisionTeams []*models.Team
	)

	if divisions, err = s.DivisionsByEvent(event); err != nil {
		return nil, err
	}

	for _, division := range divisions {
		if divisionTeams, err = s.TeamsByEventIdAndDivisionId(event.Id, division.Id); err != nil {
			return nil, err
		}

		for _, team := range divisionTeams {
			team.AgeGroup = division.Name
		}

		teams = append(teams, divisionTeams...)
	}

	return teams, nil
}

// TeamsByEventName returns all of the teams for the given event name.
func (s *GlobalService) TeamsByEventName(eventName string) ([]*models.Team, error) {
	var (
		err           error
		divisions     []models.Division
		teams         []*models.Team
		divisionTeams []*models.Team
		event         *models.Event
	)

	if event, err = s.EventByName(eventName); err != nil {
		return nil, err
	}

	if divisions, err = s.DivisionsByEventName(eventName); err != nil {
		return nil, err
	}

	for _, division := range divisions {
		if divisionTeams, err = s.TeamsByEventIdAndDivisionId(event.Id, division.Id); err != nil {
			return nil, err
		}

		teams = append(teams, divisionTeams...)
	}

	return teams, nil
}

// We can also get the teams in the conference for a given age group once we know the eventID and divisionID

// /api/Event/get-event-division-teams/{eventID}/{divisionID}

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

// TeamsByEventIdAndDivisionId returns all of the teams for the given event and division.
func (s *GlobalService) TeamsByEventIdAndDivisionId(eventId, divisionId int) ([]*models.Team, error) {
	var (
		err       error
		data      []byte
		targetUrl string
		pResponse *http.Response
		output    struct {
			Result string `json:"result"`
			Data   struct {
				TeamList []models.Team `json:"teamList"`
			} `json:"data"`
		}
	)

	suffix := fmt.Sprintf("/get-event-division-teams/%d/%d", eventId, divisionId)
	if targetUrl, err = url.JoinPath(s.Url(), "/api/Event", suffix); err != nil {
		return nil, err
	}

	if pResponse, err = s.Client().Get(targetUrl); err != nil {
		return nil, fmt.Errorf("error getting teams for event %d division %d: %v", eventId, divisionId, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(pResponse.Body)

	if pResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting teams for event %d division %d: invalid status code %d", eventId, divisionId, pResponse.StatusCode)
	}

	if data, err = io.ReadAll(pResponse.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v\n%svc", err, string(data))
	}

	var teams []*models.Team

	for _, team := range output.Data.TeamList {
		teams = append(teams, &team)
	}

	return teams, nil
}

// TeamsByEventAndDivision returns all of the teams for the given event and division.
func (s *GlobalService) TeamsByEventAndDivision(event models.Event, division models.Division) ([]*models.Team, error) {
	return s.TeamsByEventIdAndDivisionId(event.Id, division.Id)
}
