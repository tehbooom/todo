package cmd

import (
	"bytes"
	"testing"

	"github.com/tehbooom/todo/internal/helpers"
	"github.com/tehbooom/todo/internal/task"
)

func TestEditCmd(t *testing.T) {
	tempPath := t.TempDir() + "/edit.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	cmd := EditCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"1", "-d", tempPath, "task 2 edited"})
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}
	cmd.SetArgs([]string{"0", "-d", tempPath, "task 1 edited", "-g", "newgrp1"})
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
				Item:      "task 1 edited",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "newgrp1",
			},
			{
				ID:        1,
				Item:      "task 2 edited",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []task.Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	if !helpers.CompareTask(*expectedTask, actualTask) {
		t.Errorf("Task removed does not match: got \"%v\", expected \"%v\"", actualTask, *expectedTask)
	}
}
