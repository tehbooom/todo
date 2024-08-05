/*
Copyright © 2024 Alec Carpenter
*/
package group

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// rmCmd represents the rm command
func rmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Removes a task from a group",
		Long: `Will delete a group and all tasks that are in this group. You cannot
keep the task if the group is deleted. To do so first remove the task from the group then delete the group.


To remove a task run from a group:

todo edit <task_id> --group ""`,
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
			err = t.RemoveGroup(name)
			if err != nil {
				return err
			}
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
