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

func (s *AssociationService) GetAllCountries() ([]models.Country, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result    string           `json:"result"`
		Countries []models.Country `json:"data"`
	}

	if targetUrl, err = s.GetUrl("get-all-countries"); err != nil {
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

func (s *AssociationService) GetCurrentOrganizations() ([]models.Organization, error) {
	var err error
	var data []byte
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result        string                `json:"result"`
		Organizations []models.Organization `json:"data"`
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
		if strings.Contains(org.Name, "ECNL") {
			ecnlOrganizations = append(ecnlOrganizations, org)
		}
	}

	return ecnlOrganizations, nil
}
