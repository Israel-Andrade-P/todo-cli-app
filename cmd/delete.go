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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task",
	Long:  `Deletes a task from your list by ID.\nEx: delete <task ID>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Provide an task ID to delete.")
			return
		}
		message, err := todo.Delete(args[0])
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Println("Please add some todos before doing this command.")
				return
			}
			log.Fatalf("error ERR: %v", err)
		}
		fmt.Println(message)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
