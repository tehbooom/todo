/*
Copyright © 2024 Alec Carpenter
*/
package group

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// addCmd represents the add command
func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Creates a group",
		Long: `Creates a group that will hold similar tasks.

Run:
todo group add <group_name>`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("provide group name")
			}
			name := args[0]
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
			t.CreateGroup(name)
			err = t.WriteTasks(filename)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	return cmd
}
