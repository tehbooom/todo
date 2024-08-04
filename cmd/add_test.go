package cmd

import (
	"bytes"
	"testing"

	"github.com/tehbooom/todo/internal/helpers"
	"github.com/tehbooom/todo/internal/task"
)

func TestAddCmd(t *testing.T) {
	tempPath := t.TempDir() + "/add.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	cmd := AddCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"-d", tempPath, "task 3 desc"})
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}
	cmd.SetArgs([]string{"-d", tempPath, "task 4 desc", "-g", "grp1"})
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}
	cmd.SetArgs([]string{"-d", tempPath, "task 5 desc", "-g", "newgrp1"})
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}
	actualTask, err := task.ReadTasks(tempPath)
	if err != nil {
		t.Error(err)
	}
	var expectedTask = &task.Tasks{
		Task: []task.Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        1,
				Item:      "task 2 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
			{
				ID:        2,
				Item:      "task 3 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
			{
				ID:        3,
				Item:      "task 4 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        4,
				Item:      "task 5 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "newgrp1",
			},
		},
		Groups: []task.Group{
			{
				TaskIDs: []int{0, 3},
				Name:    "grp1",
			},
			{
				TaskIDs: []int{4},
				Name:    "newgrp1",
			},
		},
	}
	if !helpers.CompareTask(*expectedTask, actualTask) {
		t.Errorf("Task removed does not match: got \"%v\", expected \"%v\"", actualTask, *expectedTask)
	}
}
