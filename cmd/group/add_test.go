package group

import (
	"bytes"
	"testing"

	"github.com/tehbooom/todo/internal/helpers"
	"github.com/tehbooom/todo/internal/task"
)

func TestAddCmd(t *testing.T) {
	tempPath := t.TempDir() + "/add_group.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	cmd := addCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"-d", tempPath, "newgrp1"})
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
		},
		Groups: []task.Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
			{
				TaskIDs: []int{},
				Name:    "newgrp1",
			},
		},
	}
	if !helpers.CompareTask(*expectedTask, actualTask) {
		t.Errorf("Task removed does not match: got \"%v\", expected \"%v\"", actualTask, *expectedTask)
	}
}

func TestAddErrCmd(t *testing.T) {
	tempPath := t.TempDir() + "/add_err_group.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	cmd := addCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"-d", tempPath})
	err = cmd.Execute()
	if err == nil {
		t.Errorf("Did not provide group name and command returned no error")
	}
}
