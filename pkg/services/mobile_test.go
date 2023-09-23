package services_test

import (
	"github.com/jedi-knights/ecnl/pkg/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mobile", func() {
	var pService *services.MobileService

	BeforeEach(func() {
		pService = services.NewMobileService()
	})

	AfterEach(func() {
		pService = nil
	})

	Describe("GetEventTypes", func() {
		It("should return all event types", func() {
			// Act
			eventTypes, err := pService.GetEventTypes()

			// Assert
			Expect(eventTypes).NotTo(BeNil())
			Expect(len(eventTypes)).To(Equal(4))
			Expect(eventTypes[0].ToString()).To(Equal("Id: 1, Name: Tournament"))
			Expect(eventTypes[1].ToString()).To(Equal("Id: 2, Name: League"))
			Expect(eventTypes[2].ToString()).To(Equal("Id: 3, Name: Non TGS Event"))
			Expect(eventTypes[3].ToString()).To(Equal("Id: 4, Name: User Added Event"))

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
