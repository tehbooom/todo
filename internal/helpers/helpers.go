package helpers

import (
	"slices"

	"github.com/tehbooom/todo/internal/task"
)

func CompareTask(expected, actual task.Tasks) bool {
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

func InitializeTask(path string) error {
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
	err := bootstrapTask.WriteTasks(path)
	if err != nil {
		return err
	}
	return nil
}
