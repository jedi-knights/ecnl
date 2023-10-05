package pkg_test

import (
    "context"
    "github.com/jedi-knights/ecnl/pkg"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "net/http"
    "time"
)

const targetUrl = "https://raw.githubusercontent.com/ocrosby/soccer-data/main/org/club_translations.json"

var _ = Describe("ClubTranslate", Ordered, func() {
    var (
        ctx   context.Context
        cli   *http.Client
        trans *pkg.ClubTranslate
    )

    BeforeAll(func() {
        var (
            err error
        )

        cli = &http.Client{}
        ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
        trans, err = pkg.NewClubTranslateFromUrl(targetUrl, ctx, cli)

        Expect(err).NotTo(HaveOccurred())
        Expect(trans).NotTo(BeNil())
    })

    AfterAll(func() {
        trans = nil
        ctx = nil
        cli = nil
    })

    Describe("Translate", func() {
        It("should trans 'Albion Hurricanes FC (TX)' to 'Albion Hurricanes FC'", func() {
            // Arrange
            clubName := "Albion Hurricanes FC (TX)"

            // Act
            translatedClubName, err := trans.Translate(clubName)

            // Assert
            Expect(err).NotTo(HaveOccurred())
            Expect(translatedClubName).To(Equal("Albion Hurricanes FC"))
        })

        It("should translate 'So Cal Blues' to 'So Cal Blues SC'", func() {
            // Arrange
            clubName := "So Cal Blues"

            // Act
            translatedClubName, err := trans.Translate(clubName)

            // Assert
            Expect(err).NotTo(HaveOccurred())
            Expect(translatedClubName).To(Equal("So Cal Blues SC"))
        })

        It("should not modify 'Luke FC'", func() {
            // Arrange
            clubName := "Luke FC"

            // Act
            translatedClubName, err := trans.Translate(clubName)

            // Assert
            Expect(err).NotTo(HaveOccurred())
            Expect(translatedClubName).To(Equal("Luke FC"))
        })
    })
})
