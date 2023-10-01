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
			organizations, err := pService.GetCurrentOrganizations(true)

			// Assert
			Expect(organizations).NotTo(BeNil())
			Expect(len(organizations)).To(Equal(6))
			Expect(organizations[0].ToString()).To(Equal("Name: \"BOYS PRE-ECNL\", Id: 22, SeasonId: 56, SeasonGroupId: 8"))
			Expect(organizations[1].ToString()).To(Equal("Name: \"ECNL Boys\", Id: 12, SeasonId: 50, SeasonGroupId: 8"))
			Expect(organizations[2].ToString()).To(Equal("Name: \"ECNL Boys Regional League\", Id: 16, SeasonId: 52, SeasonGroupId: 8"))
			Expect(organizations[3].ToString()).To(Equal("Name: \"ECNL Girls\", Id: 9, SeasonId: 49, SeasonGroupId: 8"))
			Expect(organizations[4].ToString()).To(Equal("Name: \"ECNL Girls Regional League\", Id: 13, SeasonId: 51, SeasonGroupId: 8"))
			Expect(organizations[5].ToString()).To(Equal("Name: \"GIRLS PRE-ECNL\", Id: 21, SeasonId: 55, SeasonGroupId: 8"))

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
