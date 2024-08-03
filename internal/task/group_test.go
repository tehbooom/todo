package task

import "testing"

func TestAddTaskToGroup(t *testing.T) {
	var newTask = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []Group{},
	}

	// Test adding task to a group that does not exist
	err := newTask.addTaskToGroup("grp1", 0)
	if err != nil {
		t.Error(err)
	}
	var groupCreated = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
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
	if !compareTask(groupCreated, newTask) {
		t.Errorf("Task group created does not match: got %v, expected %v", newTask, groupCreated)
	}

	// Test adding task to a group that does exist
	var taskAdd = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        1,
				Item:      "task 1 desc",
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
	err = taskAdd.addTaskToGroup("grp1", 1)
	if err != nil {
		t.Error(err)
	}
	var taskAdded = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        1,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0, 1},
				Name:    "grp1",
			},
		},
	}
	if !compareTask(taskAdded, taskAdd) {
		t.Errorf("Task added to existing group does not match: got %v, expected %v", taskAdd, taskAdded)
	}
}

func TestRemoveTaskFromGroup(t *testing.T) {
	var taskRemove = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        1,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{0, 1},
				Name:    "grp1",
			},
		},
	}
	err := taskRemove.removeTaskFromGroup(1)
	if err != nil {
		t.Error(err)
	}
	var taskRemoved = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        1,
				Item:      "task 1 desc",
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
	if !compareTask(taskRemoved, taskRemove) {
		t.Errorf("Task group created does not match: got %v, expected %v", taskRemove, taskRemoved)
	}
}

func TestRemoveGroup(t *testing.T) {
	var groupRemove = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 0 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "grp1",
			},
			{
				ID:        1,
				Item:      "task 1 desc",
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

	err := groupRemove.RemoveGroup("grp1")
	if err != nil {
		t.Error(err)
	}
	var groupRemoved = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []Group{},
	}
	if !compareTask(groupRemoved, groupRemove) {
		t.Errorf("Task group created does not match: got %v, expected %v", groupRemove, groupRemoved)
	}
}

func TestCreategroup(t *testing.T) {
	var noGroup = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []Group{},
	}
	err := noGroup.CreateGroup("grp1")
	if err != nil {
		t.Error(err)
	}
	var groupAdded = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{},
				Name:    "grp1",
			},
		},
	}
	if !compareTask(groupAdded, noGroup) {
		t.Errorf("Task group created does not match: got %v, expected %v", noGroup, groupAdded)
	}
}

func TestCheckGroupCreated(t *testing.T) {
	var group = &Tasks{
		Task: []Task{
			{
				ID:        0,
				Item:      "task 1 desc",
				Timestamp: "2024-06-23T11:19:56-04:00",
				Group:     "",
			},
		},
		Groups: []Group{
			{
				TaskIDs: []int{},
				Name:    "grp1",
			},
		},
	}
	if !group.checkGroupCreated("grp1") {
		t.Errorf("Group exists and should return true")
	}
	if group.checkGroupCreated("none") {
		t.Errorf("Group does not exist and should return false")
	}
}
