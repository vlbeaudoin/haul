/*
Copyright © 2021 The Haul Authors

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// assemblyCmd represents the assembly command
var assemblyCmd = &cobra.Command{
	Use:   "assembly",
	Short: "Manipulate assemblies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("assembly called")
	},
}

func init() {
	rootCmd.AddCommand(assemblyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// assemblyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// assemblyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
