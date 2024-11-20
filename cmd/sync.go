/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync your CLI with between your local and remote",
	Long:  "Sync your CLI with between your local and remote",
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Sync the local and remote commands

	// Also show a lodash spinner while syncing

	// Show a success message after syncing
}
