/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
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
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
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
			t.ListTasks(cmd.OutOrStdout())
			return nil
		},
	}
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	return cmd
}
