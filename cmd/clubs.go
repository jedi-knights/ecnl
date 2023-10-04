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
)

// clubsCmd represents the clubs command
var clubsCmd = &cobra.Command{
	Use:   "clubs",
	Short: "Display all clubs",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err   error
			clubs []models.Club
		)

		service := services.NewTGSService()

		if clubs, err = service.ClubsByOrganizationId(orgId); err != nil {
			panic(err)
		}

		fmt.Printf("There are a total of %d clubs.\n", len(clubs))

		for _, club := range clubs {
			fmt.Printf("\t%s\n", club.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(clubsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clubsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clubsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	clubsCmd.Flags().IntVarP(&orgId, "orgId", "o", 0, "The organization id")
	_ = clubsCmd.MarkFlagRequired("orgId")
}
