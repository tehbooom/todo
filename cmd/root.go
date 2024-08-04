/*
Copyright © 2024 Alec Carpenter
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/cmd/group"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{

		Use:          "todo",
		Short:        "todo CLI app",
		Long:         `todo app thats simple and not unique allowing you to add, list and remove tasks`,
		SilenceUsage: true,
	}
	rootCmd.AddCommand(RmCmd())
	rootCmd.AddCommand(EditCmd())
	rootCmd.AddCommand(ListCmd())
	rootCmd.AddCommand(group.GroupCmd())
	rootCmd.AddCommand(AddCmd())
	return rootCmd
}

func Execute() error {
	rootCmd := NewRootCmd()
	return rootCmd.Execute()
}
