/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize zsync",
	Long: `Initialize zsync by running this command.
	This will ask you for your remote repository URL and will periodically sync your CLI commands between your local and remote machines.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		initZsync()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// initializes zsync	
func initZsync() {
	// setup a sqlite database

	// ask for remote URLs, pass them in comma separated format

	// check for all remote URLs are active

	// check for any remote sync config and set it up

	// ask and try to sync with remote
	
	// setup a cron job to sync periodically

	// print success message
	fmt.Println("zsync initialized successfully")
	fmt.Println("Use { zsync [command] } to sync your CLIs")
}
