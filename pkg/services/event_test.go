package services_test

import (
	"github.com/jedi-knights/ecnl/pkg/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	var pService *services.EventService

	BeforeEach(func() {
		pService = services.NewEventService()
	})

	AfterEach(func() {
		pService = nil
	})

	Describe("GetClubsByOrganizationId", func() {
		It("should return all clubs by organization id - ECNL Girls", func() {
			// Act
			clubs, err := pService.GetClubsByOrganizationId(9)

			// Assert
			Expect(clubs).NotTo(BeNil())
			Expect(len(clubs)).To(Equal(128))

			Expect(err).NotTo(HaveOccurred())
			Expect(clubs[0].Name).To(Equal("Alabama FC"))
			Expect(clubs[len(clubs)-1].Name).To(Equal("World Class FC"))
		})
	})
})
