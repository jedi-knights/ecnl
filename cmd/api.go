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
	_ "github.com/jedi-knights/ecnl/docs"
	v1routes "github.com/jedi-knights/ecnl/pkg/routes/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A Restful API for providing ECNL data",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set some sane default values in case they are not set in the config file.
		viper.SetDefault("env", "development")
		viper.SetDefault("mongo.uri", "mongodb://localhost:27017/ecnl")

		env := viper.GetString("env")

		e := echo.New()

		switch env {
		case "development":
			log.Printf("Development environment")

			e.Logger.SetLevel(log.DEBUG)
		case "production":
			log.Printf("Production environment")

			e.HideBanner = true
			e.Logger.SetLevel(log.INFO)

			// Setup SSL/TLS
			// Redirect all HTTP requests to HTTPS
			e.Pre(middleware.HTTPSRedirect())
		default:
			// Do nothing for now on an unknown environment.
			log.Fatalf("Unknown environment: %s expected development|production", env)
		}

		e.Logger.Info("Starting server")

		// e.Use(middleware.CORS())

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				echo.GET,
				echo.PUT,
				echo.POST,
				echo.DELETE,
				echo.PATCH,
				echo.HEAD,
				echo.OPTIONS,
				echo.CONNECT,
				echo.TRACE,
			},
		}))

		e.Pre(middleware.RemoveTrailingSlash())
		e.Pre(middleware.Logger())

		v1 := e.Group("/api/v1")

		v1.GET("/health", v1routes.HandleHealthCheck)
		v1.GET("/version", v1routes.HandleVersion)
		v1.GET("/rpi/:division", v1routes.HandleGetRPIRankings)

		e.GET("/swagger/*", echoSwagger.WrapHandler)

		// Redirect root path to /swagger/index.html
		e.GET("/", func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})

		// Redirect /swagger to /swagger/index.html
		e.GET("/swagger", func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})

		//if env == "production" {
		//	certFile := viper.GetString("ssl.cert")
		//	keyFile := viper.GetString("ssl.key")
		//
		//	e.Logger.Fatal(e.StartTLS(":8080", certFile, keyFile))
		//} else {
		//	e.Logger.Fatal(e.Start(":8080"))
		//}
		e.Logger.Fatal(e.Start(":8080"))
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
