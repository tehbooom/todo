/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long: `Adds a task and assigns it an ID

Run:
todo add <task_description>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("Must provide task to add")
		}
		taskDescription := strings.Join(args, " ")
		taskFile, err := task.TaskFilePath()
		if err != nil {
			log.Fatal(err)
		}
		newTask, err := task.AddTask(taskFile, taskDescription)
		if err != nil {
			log.Fatal(err)
		}
		taskWithID := fmt.Sprintf("Added task %d: %s", newTask.ID, newTask.Item)
		fmt.Println(taskWithID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
