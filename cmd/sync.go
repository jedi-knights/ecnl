/*
Copyright © 2023 Omar Crosby <omar.crosby@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"github.com/jedi-knights/ecnl/pkg/dal"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/jedi-knights/ecnl/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs data from the ECNL backend to the local database",
	Long: `Heavy calculations like RPI generation require a lot of back and forth
with the ECNL backend.  This command will sync the local database with the ECNL backend
so that the RPI computation can occur more rapidly.
`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err           error
			client        *mongo.Client
			clubs         []models.Club
			events        []models.Event
			organizations []models.Organization
		)

		ctx, _ := context.WithTimeout(context.Background(), 1*time.Hour)

		// get the client
		client = dal.MustGetClient(ctx)

		// get the database
		database := client.Database("ecnl")

		// create collections
		orgsCollection := database.Collection("organizations")
		clubsCollection := database.Collection("clubs")
		eventsCollection := database.Collection("events")
		matchesCollection := database.Collection("matches")
		teamsCollection := database.Collection("teams")

		// create data access objects
		orgDAO := dal.NewOrganizationDAO(ctx, orgsCollection)
		clubDAO := dal.NewClubDAO(ctx, clubsCollection)
		eventDAO := dal.NewEventDAO(ctx, eventsCollection)
		matchEventDAO := dal.NewMatchEventDAO(ctx, matchesCollection)
		teamDAO := dal.NewTeamDAO(ctx, teamsCollection)

		// Index Reference
		// https://www.mongodb.com/docs/drivers/go/current/fundamentals/indexes/

		// index the collections
		if err = orgDAO.Index(); err != nil {
			log.Fatal(err)
		}
		if err = clubDAO.Index(); err != nil {
			log.Fatal(err)
		}
		if err = teamDAO.Index(); err != nil {
			log.Fatal(err)
		}
		if err = eventDAO.Index(); err != nil {
			log.Fatal(err)
		}
		if err = teamDAO.Index(); err != nil {
			log.Fatal(err)
		}

		// create a new total global sports service
		svc := services.NewTGSService()

		// sync the organizations
		if organizations, err = svc.Organizations(false); err != nil {
			log.Fatal(err)
		}
		if err = orgDAO.SyncAll(organizations); err != nil {
			log.Fatal(err)
		}

		// sync the clubs, events, and matches
		for _, org := range organizations {
			if clubs, err = svc.ClubsByOrganization(org); err != nil {
				log.Fatal(err)
			}
			if err = clubDAO.SyncAll(clubs); err != nil {
				log.Fatal(err)
			}

			if events, err = svc.EventsByOrganization(org); err != nil {
				log.Fatal(err)
			}
			if err = eventDAO.SyncAll(events); err != nil {
				log.Fatal(err)
			}

			log.Printf("Done syncing events for organization '%s'", org.Name)

			for _, event := range events {
				var teams []*models.Team

				if teams, err = svc.TeamsByEvent(event); err != nil {
					log.Fatal(err)
				}

				if err = teamDAO.SyncAll(teams); err != nil {
					log.Fatal(err)
				}
			}

			log.Printf("Done syncing teams for organization '%s'", org.Name)

			for _, club := range clubs {
				var event *models.Event
				var data []models.MatchEvent

				// Some clubs aren't associated with an event.
				// I think this means they are no longer part of the ECNL, so in this case
				// it's safe to continue on because they don't have any match data anyway for
				// the current season.
				if club.EventId == 0 {
					continue
				}

				if event, err = svc.EventById(club.EventId); err != nil {
					log.Fatal(err)
				}

				log.Printf("Syncing match results for club '%s' and event '%s' ...", club.Name, event.Name)
				if data, err = svc.MatchEventsByClubNameAndEventName(club.Name, event.Name); err != nil {
					log.Fatal(err)
				}

				if err = matchEventDAO.SyncAll(data); err != nil {
					log.Fatal(err)
				}
			}
		}

		// 3. The teams
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
