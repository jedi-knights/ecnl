/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/controllers"
	"github.com/jedi-knights/ecnl/pkg/dal"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// rpigenCmd represents the rpigen command
var rpigenCmd = &cobra.Command{
	Use:   "rpigen",
	Short: "A command to generate and store RPI rankings for an age grouop.",
	Long: `In order to break out the process of generating RPI data into the database

I have opted to create the rpigen command with the explicit purpose of generating and
storing RPI data on command.  This way the API only becomes a consumer and doesn't
end up polluting the database with record after record when all it needs to do is
read.

This means from time to time this command will have to be run in order to get updated
data into the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err  error
			ctrl *controllers.RPI
			data []models.RPIRankingData
			age  string
		)

		ctrl = controllers.NewRPI()

		if age, err = cmd.Flags().GetString("age"); err != nil {
			log.Fatalf("Unable to retrieve the age parameter: %v\n", err)
		}

		if data, err = ctrl.GenerateRankings(age); err != nil {
			log.Fatalf("Error generating rankings: %s\n", err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)

		// get the client
		client := dal.MustGetClient(ctx)

		// get the database
		database := client.Database("ecnl")

		rpiEventsCollection := database.Collection("rpi_events")
		rpiEventDAO := dal.NewRPIEventDAO(ctx, rpiEventsCollection)

		currentTime := time.Now()

		fmt.Printf("RPI Rankings for %s %s\n", orgName, ageGroup)
		for _, d := range data {
			// Attempt to append the RPI ranking
			var event = models.RPIEvent{
				Timestamp: currentTime,
				TeamId:    d.TeamId,
				TeamName:  d.TeamName,
				Ranking:   d.Ranking,
				Value:     d.RPI,
			}

			if err = rpiEventDAO.Create(event); err != nil {
				log.Println(err)
			} else {
				formattedTime := currentTime.Format("January 2, 2006 3:04 PM MST")
				fmt.Println("Saved " + formattedTime + " " + d.String())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rpigenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpigenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpigenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rpigenCmd.PersistentFlags().StringP("age", "a", "", "Age group (e.g. G2009)")
	_ = rpigenCmd.MarkPersistentFlagRequired("age")
}
