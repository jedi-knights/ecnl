package services

import (
	"encoding/json"
	"github.com/jedi-knights/ecnl/pkg"
	"github.com/jedi-knights/ecnl/pkg/models"
	"net/http"
	"net/url"
)

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

func (s *MobileService) GetUrl(path string) (string, error) {
	var err error
	var targetUrl string

	if targetUrl, err = url.JoinPath(pkg.TgsPrefix, "/api/Mobile", path); err != nil {
		return "", err
	}

	return targetUrl, nil
}

func (s *MobileService) GetEventTypes() ([]models.EventType, error) {
	var err error
	var targetUrl string
	var pResponse *http.Response
	var output struct {
		Result     string             `json:"result"`
		EventTypes []models.EventType `json:"data"`
	}

	if targetUrl, err = s.GetUrl("get-event-types"); err != nil {
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
