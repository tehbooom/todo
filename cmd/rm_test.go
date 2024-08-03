package cmd

import (
	"bytes"
	"log"
	"slices"
	"testing"

	"github.com/tehbooom/todo/internal/task"
)

func initializeTask(t *testing.T) string {
	var bootstrapTask = &task.Tasks{
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
		},
	}
	tdJson := t.TempDir() + "/td.json"
	err := bootstrapTask.WriteTasks(tdJson)
	if err != nil {
		t.Error(err)
	}
	return tdJson
}

func TestRmCmd(t *testing.T) {
	path := initializeTask(t)
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
	actual := new(bytes.Buffer)
	NewRootCmd().SetOut(actual)
	NewRootCmd().SetErr(actual)
	log.SetOutput(actual)
	NewRootCmd().SetArgs([]string{"rm", "1", "-d", path})
	NewRootCmd().Execute()
	actualTask, err := task.ReadTasks(path)
	if err != nil {
		t.Error(err)
	}
	if !compareTask(expectedTask, &actualTask) {
		t.Errorf("Task removed does not match: got %v, expected %v", actualTask, expectedTask)
	}
}

func compareTask(expected, actual *task.Tasks) bool {
	for i, v := range expected.Task {
		if v.Group != actual.Task[i].Group {
			return false
		}
		if v.Item != actual.Task[i].Item {
			return false
		}
		if v.ID != actual.Task[i].ID {
			return false
		}
	}

	for j, group := range expected.Groups {
		if group.Name != actual.Groups[j].Name {
			return false
		}
		if !slices.Equal(group.TaskIDs, actual.Groups[j].TaskIDs) {
			return false
		}
	}

	return true
}
