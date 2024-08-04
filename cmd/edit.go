/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tehbooom/todo/internal/task"
)

func EditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit existing task",
		Long: `Edit an existing task by ID number. You can edit the group it is in and the description.
Editing a task does update the timestamp

To edit a task run:
todo edit <task_id> <task_description>

Optionally you can specify a group or leav a group as empty will remove it from a group.
todo edit <task_id> --group example_group <task_description>
todo edit <task_id> --group`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("must provide task ID to edit")
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			group, _ := cmd.Flags().GetString("group")
			groupSet := cmd.Flags().Changed("group")
			path, _ := cmd.Flags().GetString("data-file")
			pathSet := cmd.Flags().Changed("data-file")
			taskDescription := strings.Join(args[1:], " ")
			filename, err := task.FilePath(path, pathSet)
			if err != nil {
				return err
			}
			t, err := task.ReadTasks(filename)
			if err != nil {
				return err
			}
			err = t.EditTask(taskDescription, group, id, groupSet)
			if err != nil {
				return err
			}
			err = t.WriteTasks(filename)
			if err != nil {
				return err
			}
			t.ListTasks(cmd.OutOrStdout())
			return nil
		},
	}
	cmd.Flags().StringP("group", "g", "", "Specify the group the task should be in")
	cmd.Flags().StringP("data-file", "d", "~/.td.json", "Path to file storing tasks")
	return cmd
}
