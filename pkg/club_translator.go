package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"net/http"
)

type ClubTranslator interface {
	Translate(clubName string) string
}

type ClubTranslate struct {
	translations *models.ClubTranslations
}

func NewClubTranslate(translations *models.ClubTranslations) *ClubTranslate {
	return &ClubTranslate{translations: translations}
}

// NewClubTranslateFromUrl creates a new club translator from the specified URL.
func NewClubTranslateFromUrl(targetUrl string, ctx context.Context, client *http.Client) (*ClubTranslate, error) {
	var (
		err          error
		translations *models.ClubTranslations
	)

	if translations, err = readMappings(targetUrl, ctx, client); err != nil {
		return nil, err
	}

	return &ClubTranslate{translations: translations}, nil
}

// Translate translates the club name to the new name if a match is found.
func (ct *ClubTranslate) Translate(clubName string) (string, error) {
	for _, mapping := range ct.translations.Data {
		if mapping.From == clubName {
			return mapping.To, nil
		}
	}

	return clubName, nil
}

// readMappings reads the mappings from the specified URL.
func readMappings(targetURL string, ctx context.Context, client *http.Client) (*models.ClubTranslations, error) {
	var (
		err error
		req *http.Request
		res *http.Response
	)

	// Create an HTTP GET request
	if req, err = http.NewRequestWithContext(ctx, "GET", targetURL, nil); err != nil {
		return nil, err
	}

	// Execute the request
	if res, err = client.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP GET request failed with status: %s", res.Status)
	}

	// Decode the JSON response
	var data models.ClubTranslations
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
