package services_test

import (
	"github.com/jedi-knights/ecnl/pkg/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Association", func() {
	var pService *services.AssociationService

	BeforeEach(func() {
		pService = services.NewAssociationService()
	})

	AfterEach(func() {
		pService = nil
	})

	Describe("GetAllCountries", func() {
		It("should return all countries", func() {
			// Act
			countries, err := pService.GetAllCountries()

			// Assert
			Expect(countries).NotTo(BeNil())
			Expect(len(countries)).To(Equal(249))

			Expect(err).NotTo(HaveOccurred())
			Expect(countries[0].ToString()).To(Equal("Id: 1, Name: United States"))
			Expect(countries[len(countries)-1].ToString()).To(Equal("Id: 249, Name: Zimbabwe"))
		})
	})

	Describe("GetAllStates", func() {
		It("should return all states", func() {
			// Act
			states, err := pService.GetAllStates()

			// Assert
			Expect(states).NotTo(BeNil())
			Expect(len(states)).To(Equal(51))

			Expect(err).NotTo(HaveOccurred())

			Expect(states[0].ToString()).To(Equal("Id: 4, Name: Alabama"))
			Expect(states[len(states)-1].ToString()).To(Equal("Id: 52, Name: Wyoming"))
		})
	})
})
