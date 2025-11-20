/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"log"

	"github.com/Israel-Andrade-P/todo-cli-app.git/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks",
	Long:  `It lists all your current tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := todo.ListAll()
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Println("Please add some todos before doing this command.")
				return
			}
			log.Fatalf("error ERR: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
