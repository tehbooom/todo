/*
Copyright © 2024 Alec Carpenter
*/

package group

import (
	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

// groupCmd represents the group command
func GroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Lists groups",
		Long: `Lists all groups that exist

to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
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
			t.ListGroups(cmd.OutOrStdout())
			return nil
		},
	}
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	cmd.AddCommand(rmCmd())
	cmd.AddCommand(addCmd())
	return cmd
}
