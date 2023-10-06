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
	"fmt"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/jedi-knights/ecnl/pkg/services"
	"github.com/spf13/cobra"
	"os"
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err    error
			events []models.Event
		)

		service := services.NewTGSService()

		if orgId > 0 {
			if orgName != "" {
				fmt.Println("Must specify either orgId or orgName, not both")
				os.Exit(2)
			}

			fmt.Printf("Getting events for orgId %d ...\n", orgId)
			if events, err = service.EventsByOrgId(orgId); err != nil {
				panic(err)
			}
		} else if orgName != "" {
			fmt.Printf("Getting events for orgName '%s' ...\n", orgName)
			if events, err = service.EventsByOrgName(orgName); err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("Getting events for all organizations ...\n")
			if events, err = service.Events(); err != nil {
				panic(err)
			}
		}

		fmt.Printf("There are a total of %d events.\n", len(events))
		for _, event := range events {
			fmt.Printf("\t%s\n", event.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eventsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	eventsCmd.Flags().IntVarP(&orgId, "id", "i", 0, "Organization ID")
	_ = eventsCmd.MarkFlagRequired("orgId")

	eventsCmd.Flags().StringVarP(&orgName, "name", "n", "", "Organization Name")
	_ = eventsCmd.MarkFlagRequired("orgName")
}
