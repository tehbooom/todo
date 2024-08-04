package cmd

import (
	"bytes"
	"testing"

	"github.com/tehbooom/todo/internal/helpers"
	"github.com/tehbooom/todo/internal/task"
)

func TestRmCmd(t *testing.T) {
	tempPath := t.TempDir() + "/rm.json"
	err := helpers.InitializeTask(tempPath)
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
		},
		Groups: []task.Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	cmd := RmCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"1", "-d", tempPath})
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	actualTask, err := task.ReadTasks(tempPath)
	if err != nil {
		t.Error(err)
	}
	if !helpers.CompareTask(*expectedTask, actualTask) {
		t.Errorf("Task removed does not match: got \"%v\", expected \"%v\"", actualTask, *expectedTask)
	}
}