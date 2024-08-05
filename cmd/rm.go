/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
	"fmt"
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
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("please provide a single number")
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			path, err := cmd.Flags().GetString("data-file")
			if err != nil {
				return err
			}
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
				return err
			}
			t.ListTasks(cmd.OutOrStdout(), "", false)
			return nil
		},
	}
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	return cmd
}
