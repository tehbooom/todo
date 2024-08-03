/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// addCmd represents the add command
func AddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a task",
		Long: `Adds a task and assigns it an ID

Run:
todo add <task_description>`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				log.Fatalf("Must provide task to add")
			}
			path, _ := cmd.Flags().GetString("data-file")
			pathSet := cmd.Flags().Changed("data-file")
			group, _ := cmd.Flags().GetString("group")
			groupSet := cmd.Flags().Changed("group")
			taskDescription := strings.Join(args, " ")
			filename, err := task.FilePath(path, pathSet)
			if err != nil {
				return err
			}
			t, err := task.ReadTasks(filename)
			if err != nil {
				return err
			}
			err = t.AddTask(taskDescription, group, groupSet)
			if err != nil {
				return err
			}
			err = t.WriteTasks(filename)
			if err != nil {
				log.Fatal(err)
			}
			t.ListTasks()
			return nil
		},
	}
	cmd.Flags().StringP("group", "g", "", "Specify the group the task should be in")
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	return cmd
}
