package main

import (
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/jedi-knights/ecnl/pkg/services"
)

func displayCountries(service *services.AssociationService) {
	var err error
	var countries []models.Country

	if countries, err = service.GetAllCountries(); err != nil {
		panic(err)
	}

	for _, country := range countries {
		println(country.ToString())
	}
}

func displayStates(service *services.AssociationService) {
	var err error
	var states []models.State

	if states, err = service.GetAllStates(); err != nil {
		panic(err)
	}

	for _, state := range states {
		println(state.ToString())
	}
}

func displayOrganizations(ecnlOnly bool) {
	var err error
	var organizations []models.Organization

	associationService := services.NewAssociationService()

	if organizations, err = associationService.GetCurrentOrganizations(ecnlOnly); err != nil {
		panic(err)
	}

	if ecnlOnly {
		fmt.Printf("ECNL Organizations [%d]:\n", len(organizations))
	} else {
		fmt.Printf("Organizations [%d]:\n", len(organizations))
	}

	for offset, organization := range organizations {
		fmt.Printf("\t%d: %s\n", offset, organization.ToString())
	}
}

func displayClubsForOrganization(orgName string, esvc *services.EventService, asvc *services.AssociationService) ([]models.Club, error) {
	var err error
	var pOrganization *models.Organization
	var clubs []models.Club

	if pOrganization, err = asvc.GetOrganizationByName(orgName); err != nil {
		return nil, err
	}

	fmt.Printf("Getting clubs for organization [%s] ...\n", pOrganization.Name)
	if clubs, err = esvc.GetClubsByOrganization(pOrganization); err != nil {
		return nil, err
	}

	fmt.Printf("Clubs [%d]:\n", len(clubs))
	for offset, club := range clubs {
		fmt.Printf("\t%d: %s\n", offset, club.ToString())
	}

	return clubs, nil
}

func displayEvents(orgName string, esvc *services.EventService) ([]models.Event, error) {
	var err error
	var eventIds []int
	var events []models.Event

	if eventIds, err = esvc.GetEventIdsByOrgName(orgName); err != nil {
		return nil, err
	}

	fmt.Printf("\nEvents [%d]:\n", len(eventIds))
	for offset, eventId := range eventIds {
		var pEvent *models.Event

		if pEvent, err = esvc.GetEventById(eventId); err != nil {
			return nil, err
		}

		fmt.Printf("\t%d: %s\n", offset, pEvent.ToString())
	}

	return events, nil
}

func probe(orgName string, esvc *services.EventService, asvc *services.AssociationService) error {
	var err error
	var pOrganization *models.Organization
	var clubs []models.Club
	var events []models.Event

	fmt.Printf("Probing Organization '%s'\n", orgName)

	fmt.Printf("Getting organization [%s] ...\n", orgName)
	if pOrganization, err = asvc.GetOrganizationByName(orgName); err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Organization: %s\n", pOrganization.ToString())
	fmt.Println()

	fmt.Printf("Getting events for organization [%s] ...\n", pOrganization.Name)
	if events, err = esvc.GetEventsByOrgName(orgName); err != nil {
		return err
	}

	fmt.Println()

	fmt.Printf("Events for organization [%s] [%d]:\n", orgName, len(events))
	for _, event := range events {
		fmt.Printf("\t%s\n", event.ToString())
	}

	fmt.Println()

	fmt.Printf("Getting clubs for organization [%s] ...\n", pOrganization.Name)
	if clubs, err = esvc.GetClubsByOrganization(pOrganization); err != nil {
		return err
	}

	fmt.Printf("Clubs [%d]:\n", len(clubs))
	for _, event := range events {
		fmt.Printf("\t%s\n", event.ToString())

		for _, club := range clubs {
			if club.EventId != event.Id {
				continue
			}

			fmt.Printf("\t\t%s\n", club.ToString())
		}

		fmt.Println()
	}

	return nil
}

func main() {
	var err error

	associationService := services.NewAssociationService()
	eventService := services.NewEventService(associationService)

	displayOrganizations(false)

	if err = probe("ECNL Girls", eventService, associationService); err != nil {
		panic(err)
	}

	var pPremier, pPlatinum *models.Club

	if pPremier, err = eventService.GetClub("ECNL Girls", "Concorde Fire Premier"); err != nil {
		panic(err)
	}

	if pPlatinum, err = eventService.GetClub("ECNL Girls", "Concorde Fire Platinum"); err != nil {
		panic(err)
	}

	fmt.Print("\n")
	fmt.Printf("Premier: %s\n", pPremier.ToString())
	fmt.Printf("Platinum: %s\n", pPlatinum.ToString())
}
