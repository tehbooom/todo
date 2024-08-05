package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehbooom/todo/internal/task"
)

func TestCompareTask(t *testing.T) {
	var controlTask = &task.Tasks{
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
	tests := map[string]struct {
		input   *task.Tasks
		control *task.Tasks
		pass    bool
	}{
		"same-task-true": {
			input:   controlTask,
			control: controlTask,
			pass:    true,
		},
		"diff-task-group": {
			input: &task.Tasks{
				Task: []task.Task{
					{
						ID:        0,
						Item:      "task 1 desc",
						Timestamp: "2024-06-23T11:19:56-04:00",
						Group:     "diff",
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
			},
			control: controlTask,
			pass:    false,
		},
		"diff-task-Item": {
			input: &task.Tasks{
				Task: []task.Task{
					{
						ID:        0,
						Item:      "diff",
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
			},
			control: controlTask,
			pass:    false,
		},
		"diff-task-ID": {
			input: &task.Tasks{
				Task: []task.Task{
					{
						ID:        3,
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
			},
			control: controlTask,
			pass:    false,
		},
		"diff-group-name": {
			input: &task.Tasks{
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
						Name:    "diff",
					},
				},
			},
			control: controlTask,
			pass:    false,
		},
		"diff-group-IDs": {
			input: &task.Tasks{
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
						TaskIDs: []int{3},
						Name:    "grp1",
					},
				},
			},
			control: controlTask,
			pass:    false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result := CompareTask(*test.control, *test.input)
			if result != test.pass {
				t.Fatalf("Returned %t; expected %t", result, test.pass)
			}
		})
	}
}

func TestInitializeTask(t *testing.T) {
	tempPath := t.TempDir() + "/helper.json"
	err := InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	var expected = &task.Tasks{
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
	actual, err := task.ReadTasks(tempPath)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, *expected, actual)
}
