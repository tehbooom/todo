/*
Copyright © 2024 Alec Carpenter
*/
package group

import (
	"log"

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
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalf("provide group name")
			}
			name := args[0]
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
			err = t.RemoveGroup(name)
			if err != nil {
				log.Fatal(err)
			}
			err = t.WriteTasks(filename)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	return cmd
}
