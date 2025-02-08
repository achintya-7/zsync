/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/achintya-7/zsync/db"
	"github.com/achintya-7/zsync/utils"
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
	// initialize the store
	db.NewStore("zsync.db", "db/migrations")

	// ask for remote URLs, pass them in comma separated format

	// check for all remote URLs are active

	// check for any remote sync config and set it up

	// ask and try to sync with remote

	// setup a cron job to sync periodically

	// start the server to listen for zsh commands and sync requests

	// ask from the user to fill the .zshrc file
	var response string
	fmt.Print("Do you want to proceed with filling the .zshrc file? (y/n): ")
	fmt.Scanln(&response)
	if response != "y" {
		fmt.Println("Aborting the .zshrc file setup.")
		return
	}

	// fill the .zshr file with the required data
	msg, err := utils.CheckAndFillZshrc()
	fmt.Println(msg)
	if err != nil {
		return
	}

	// print success message
	fmt.Println("zsync initialized successfully")
}
