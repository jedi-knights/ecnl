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

type AssociationService struct {
	HttpClient *http.Client
}

var currentOrganizations []models.Organization

func NewAssociationService() *AssociationService {
	pHttpClient := &http.Client{}
	return NewAssociationServiceWithClient(pHttpClient)
}

func NewAssociationServiceWithClient(pHttpClient *http.Client) *AssociationService {
	return &AssociationService{
		HttpClient: pHttpClient,
	}
}

func (s *AssociationService) GetUrl(path string) (string, error) {
	var err error
	var targetUrl string

	if targetUrl, err = url.JoinPath(pkg.TgsPrefix, "/api/Association", path); err != nil {
		return "", err
	}

	return targetUrl, nil
}

//func (s *AssociationService) GetAllCountries() ([]models.Country, error) {
//	var err error
//	var statusCode int
//	var data []byte
//	var targetUrl string
//	var pResponse *http.Response
//	var output struct {
//		Result    string           `json:"result"`
//		Countries []models.Country `json:"data"`
//	}
//
//	if targetUrl, err = s.GetUrl("get-all-countries"); err != nil {
//		return nil, err
//	}
//
//	if statusCode, err = json2.NewClient().Get(targetUrl, &output); err != nil {
//
//	}
//
//	if pResponse, err = s.HttpClient.Get(targetUrl); err != nil {
//		return nil, fmt.Errorf("Error getting countries: %v\n", err)
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			fmt.Println(err.Error())
//			panic("an error occured while attempting to close the body")
//		}
//	}(pResponse.Body)
//
//	if pResponse.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("Invalid status code %d: ", pResponse.StatusCode)
//	}
//
//	if data, err = io.ReadAll(pResponse.Body); err != nil {
//		return nil, fmt.Errorf("Error reading response body: %v\n", err)
//	}
//
//	if err = json.Unmarshal(data, &output); err != nil {
//		return nil, fmt.Errorf("Error unmarshalling response body: %v\n", err)
//	}
//
//	return output.Countries, nil
//}

func (s *AssociationService) GetAllStates() ([]models.State, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result string         `json:"result"`
		States []models.State `json:"data"`
	}

	if targetUrl, err = s.GetUrl("/get-all-states"); err != nil {
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

func (s *AssociationService) GetCurrentOrganizations(ecnlOnly bool) ([]models.Organization, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result        string                `json:"result"`
		Organizations []models.Organization `json:"data"`
	}

	if currentOrganizations != nil {
		return currentOrganizations, nil
	}

	if targetUrl, err = s.GetUrl("/get-current-orgs-list"); err != nil {
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

func (s *AssociationService) GetOrganizationById(id int) (*models.Organization, error) {
	var err error

	if currentOrganizations == nil {
		if currentOrganizations, err = s.GetCurrentOrganizations(false); err != nil {
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

func (s *AssociationService) GetOrganizationByName(name string) (*models.Organization, error) {
	var err error

	if currentOrganizations == nil {
		if _, err = s.GetCurrentOrganizations(false); err != nil {
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

func (s *AssociationService) GetCurrentOrganizationNames() ([]string, error) {
	var err error

	if currentOrganizations == nil {
		if _, err = s.GetCurrentOrganizations(false); err != nil {
			return nil, err
		}
	}

	orgNames := make([]string, 0)

	for _, org := range currentOrganizations {
		orgNames = append(orgNames, org.Name)
	}

	return orgNames, nil
}
