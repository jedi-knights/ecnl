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

	Describe("GetCurrentOrganizations", func() {
		It("should return all current organizations", func() {
			// Act
			organizations, err := pService.GetCurrentOrganizations(false)

			// Assert
			Expect(organizations).NotTo(BeNil())
			Expect(len(organizations)).To(Equal(6))
			Expect(organizations[0].ToString()).To(Equal("Id: 22, Name: BOYS PRE-ECNL, SeasonId: 56, SeasonGroupId: 8"))
			Expect(organizations[1].ToString()).To(Equal("Id: 12, Name: ECNL Boys, SeasonId: 50, SeasonGroupId: 8"))
			Expect(organizations[2].ToString()).To(Equal("Id: 16, Name: ECNL Boys Regional League, SeasonId: 52, SeasonGroupId: 8"))
			Expect(organizations[3].ToString()).To(Equal("Id: 9, Name: ECNL Girls, SeasonId: 49, SeasonGroupId: 8"))
			Expect(organizations[4].ToString()).To(Equal("Id: 13, Name: ECNL Girls Regional League, SeasonId: 51, SeasonGroupId: 8"))
			Expect(organizations[5].ToString()).To(Equal("Id: 21, Name: GIRLS PRE-ECNL, SeasonId: 55, SeasonGroupId: 8"))

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
