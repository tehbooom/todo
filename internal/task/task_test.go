package task

import (
	"encoding/json"
	"os"
	"path"
	"slices"
	"testing"
)

func TestReadTasks(t *testing.T) {
	tmpDir := t.TempDir()
	testTmpFile := path.Join(tmpDir, "test-td.json")
	_, err := ReadTasks(testTmpFile)
	if err != nil {
		t.Error(err)
	}

	var newTask = &Tasks{
		Task: []Task{
			{
				ID:        1,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{1},
				Name:    "grp1",
			},
		},
	}
	jsonTask, err := json.Marshal(newTask)
	if err != nil {
		t.Error(err)
	}

	err = os.WriteFile(testTmpFile, jsonTask, 0600)
	if err != nil {
		t.Error(err)
	}

	readTask, err := ReadTasks(testTmpFile)
	if err != nil {
		t.Error(err)
	}

	if !compareTask(newTask, &readTask) {
		t.Errorf("Task list in file does not equal read task")
	}
}

func TestWriteTasks(t *testing.T) {
	tmpDir := t.TempDir()
	testTmpFile := path.Join(tmpDir, "test-td.json")
	task, err := ReadTasks(testTmpFile)
	if err != nil {
		t.Error(err)
	}
	err = task.WriteTasks(testTmpFile)
	if err != nil {
		t.Error(err)
	}
}

func TestAddTask(t *testing.T) {
	tmpDir := t.TempDir()
	testTmpFile := path.Join(tmpDir, "test-td.json")
	task, err := ReadTasks(testTmpFile)
	if err != nil {
		t.Error(err)
	}
	var expectedTask = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	err = task.AddTask(expectedTask.Task[0].Item, expectedTask.Task[0].Group, true)
	if err != nil {
		t.Error(err)
	}
	if !compareTask(expectedTask, &task) {
		t.Errorf("Task added does not match: got %v, expected %v", task, expectedTask)
	}

	var addedExpectedTask = &Tasks{
		Task: []Task{
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
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}

	err = task.AddTask(addedExpectedTask.Task[1].Item, addedExpectedTask.Task[1].Group, false)
	if err != nil {
		t.Error(err)
	}
	if !compareTask(addedExpectedTask, &task) {
		t.Errorf("Task added does not match: got %v, expected %v", task, expectedTask)
	}
}

func TestRemoveTask(t *testing.T) {
	var beforeTasks = &Tasks{
		Task: []Task{
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
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	err := beforeTasks.RemoveTask(2)
	if err == nil {
		t.Errorf("Task does not exist and should return error")
	}
	err = beforeTasks.RemoveTask(1)
	if err != nil {
		t.Error(err)
	}
	var afterTasks = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	if !compareTask(afterTasks, beforeTasks) {
		t.Errorf("Task removed does not match: got %v, expected %v", beforeTasks, afterTasks)
	}
}

func TestEditTask(t *testing.T) {
	var beforeTasks = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}

	// Test editing item in task
	err := beforeTasks.EditTask("task 1 desc after", "", 0, false)
	if err != nil {
		t.Error(err)
	}
	var messageTasks = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc after",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	if !compareTask(messageTasks, beforeTasks) {
		t.Errorf("Task edited does not match: got %v, expected %v", beforeTasks, messageTasks)
	}

	// Test changing group task belongs int
	err = messageTasks.EditTask("", "newgrp", 0, true)
	if err != nil {
		t.Error(err)
	}
	var groupEditTasks = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc after",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "newgrp",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	if !compareTask(groupEditTasks, messageTasks) {
		t.Errorf("Task group edited does not match: got %v, expected %v", messageTasks, groupEditTasks)
	}

	// Test removed task from a group
	err = groupEditTasks.EditTask("", "", 0, true)
	if err != nil {
		t.Error(err)
	}
	var groupRemovedTasks = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc after",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0},
				Name:    "grp1",
			},
		},
	}
	if !compareTask(groupRemovedTasks, groupEditTasks) {
		t.Errorf("Task group removed does not match: got %v, expected %v", groupEditTasks, groupRemovedTasks)
	}
}

func TestFilePath(t *testing.T) {
	pathProvided, err := FilePath("test", true)
	if err != nil {
		t.Error(err)
	}
	if pathProvided != "test" {
		t.Errorf("Path does not equal provided path: got %s, expected %s", pathProvided, "test")
	}

	pathNotProvided, err := FilePath("", false)
	if err != nil {
		t.Error(err)
	}
	testHome, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	if pathNotProvided != testHome+"/.td.json" {
		t.Errorf("Path generated does not equal %s: got %s", testHome+"/.td.json", pathNotProvided)
	}
}

func compareTask(expected, actual *Tasks) bool {
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
