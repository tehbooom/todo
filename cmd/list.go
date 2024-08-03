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
func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all tasks",
		Long: `List all tasks located in the task file

Run:
todo list`,
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("data-file")
			pathSet := cmd.Flags().Changed("data-file")
			filename, err := task.FilePath(path, pathSet)
			if err != nil {
				log.Fatal(err)
			}
			t, err := task.ReadTasks(filename)
			if err != nil {
				log.Fatal(err)
			}
			t.ListTasks()
		},
	}
	return cmd
}
