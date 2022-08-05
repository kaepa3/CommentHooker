/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed resource/version.txt
var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "dispay now chooker version",
	Long:  `dispay now chooker version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version:%s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
