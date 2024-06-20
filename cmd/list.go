/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long: `List all tasks located in the task file

Run:
todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		taskFile, err := task.TaskFilePath()
		if err != nil {
			log.Fatal(err)
		}
		err = task.ListTasks(taskFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
