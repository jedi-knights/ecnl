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
	"github.com/jedi-knights/ecnl/pkg/controllers"
	"github.com/jedi-knights/ecnl/pkg/models"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rpiCmd represents the rpi command
var rpiCmd = &cobra.Command{
	Use:   "rpi",
	Short: "Displays RPI junk",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err  error
			ctrl *controllers.RPI
			data []models.RPIRankingData
		)

		ctrl = controllers.NewRPI()

		if data, err = ctrl.GenerateRankings(ageGroup); err != nil {
			log.Printf("Error generating rankings: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("RPI Rankings for %s %s\n", orgName, ageGroup)
		for _, d := range data {
			fmt.Printf("#%d: '%s' (%f)\n", d.Ranking, d.TeamName, d.RPI)
		}
	},
}

func init() {
	rootCmd.AddCommand(rpiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rpiCmd.PersistentFlags().StringVarP(&ageGroup, "age", "a", "", "Age group (e.g. G2009)")
	rpiCmd.MarkPersistentFlagRequired("age")
}
