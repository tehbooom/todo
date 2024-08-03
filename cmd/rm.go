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
func RmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a task by ID number",
		Long: `Can only remove tasks by ID. To get a list of tasks and their ID run 'todo list'

To remove a task run:

todo rm <task_ID>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				log.Fatalf("Please provide a single number")
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			path, _ := cmd.Flags().GetString("data-file")
			pathSet := cmd.Flags().Changed("data-file")
			filename, err := task.FilePath(path, pathSet)
			if err != nil {
				return err
			}
			t, err := task.ReadTasks(filename)
			if err != nil {
				return err
			}
			err = t.RemoveTask(id)
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
	return cmd
}
