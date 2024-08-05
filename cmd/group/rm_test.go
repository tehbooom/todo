package group

import (
	"bytes"
	"testing"

	"github.com/tehbooom/todo/internal/helpers"
	"github.com/tehbooom/todo/internal/task"
)

func TestRmCmd(t *testing.T) {
	tempPath := t.TempDir() + "/rm_group.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	cmd := rmCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"-d", tempPath, "grp1"})
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
				Item:      "task 2 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []task.Group{},
	}
	if !helpers.CompareTask(*expectedTask, actualTask) {
		t.Errorf("Task removed does not match: got \"%v\", expected \"%v\"", actualTask, *expectedTask)
	}
}
