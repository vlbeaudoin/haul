/*
Copyright © 2021 The Haul Authors

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port int
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the haul server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		fmt.Println(viper.GetInt("server.port"))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntP("port", "p", 8080, "The network port to bind the server to")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
}
