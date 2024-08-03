/*
Copyright © 2024 Alec Carpenter
*/

package group

import (
	"log"

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
			t.ListGroups()
		},
	}
	cmd.AddCommand(rmCmd())
	cmd.AddCommand(addCmd())
	return cmd
}
