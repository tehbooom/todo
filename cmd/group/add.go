/*
Copyright © 2024 Alec Carpenter
*/
package group

import (
	"log"

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
			err = t.CreateGroup(name)
			if err != nil {
				log.Fatal(err)
			}
			err = t.WriteTasks(filename)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	return cmd
}
