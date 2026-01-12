/*
Copyright © 2021 The Haul Authors

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/vlbeaudoin/haul/data"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the haul server",
	Run: func(cmd *cobra.Command, args []string) {
		data.OpenDatabase()
		data.MigrateDatabase()
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	declareFlagsForServer()
}

func declareFlagsForServer() {
	serverCmd.Flags().IntP("port", "p", 8080, "The network port to bind the server to (config: \"server.port\")")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))

	serverCmd.Flags().StringP("motd", "m", "oh hisse", "Message of the day as exposed by the webserver (config: \"server.motd\")")
	viper.BindPFlag("server.motd", serverCmd.Flags().Lookup("motd"))
}

func runServer() {
	log.Print("[I] Starting webserver")

	port := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	motd := viper.GetString("server.motd")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, motd)
	})

	if err := e.Start(port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
