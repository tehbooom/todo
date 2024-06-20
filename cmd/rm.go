/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task by ID number",
	Long: `Can only remove tasks by ID. To get a list of tasks and their ID run 'todo list'

To remove a task run:

todo rm <task_ID>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			log.Fatalf("Please provide a single number")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		taskFile, err := task.TaskFilePath()
		if err != nil {
			log.Fatal(err)
		}

		err = task.RemoveTask(taskFile, id)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
