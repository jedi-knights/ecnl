package main

import (
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

func main() {
	service := services.NewAssociationService()

	displayCountries(service)
	displayStates(service)
}
