/*
Copyright © 2021 The Haul Authors

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// assetCmd represents the asset command
var assetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Manipulate assets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asset called")
	},
}

func init() {
	rootCmd.AddCommand(assetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// assetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// assetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
