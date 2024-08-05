/*
Copyright © 2024 Alec Carpenter
*/
package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Tasks struct {
	Task   []Task  `json:"tasks"`
	Groups []Group `json:"groups"`
}

type Task struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Timestamp string `json:"timestamp"`
	Group     string `json:"group"`
}

// Reads the tasks and groups from file and returns *Tasks
func ReadTasks(path string) (Tasks, error) {
	var tasks Tasks
	if _, err := os.Stat(path); err != nil {
		emptyList, err := json.Marshal(&tasks)
		if err != nil {
			return tasks, err
		}
		err = os.WriteFile(path, emptyList, 0600)
		if err != nil {
			return tasks, err
		}
	}

	tasksFile, err := os.Open(path)
	if err != nil {
		return tasks, err
	}
	defer tasksFile.Close()

	tasksByte, _ := io.ReadAll(tasksFile)
	err = json.Unmarshal(tasksByte, &tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

// Writes tasks to the file
func (t *Tasks) WriteTasks(path string) error {
	jsonTasks, err := json.Marshal(&t)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonTasks, 0600)
	if err != nil {
		return err
	}
	return nil
}

// Adds a single task
func (t *Tasks) AddTask(message, group string, groupSet bool) error {
	var newTask Task
	var newTaskID int
	if len(t.Task) != 0 {
		newTaskID = len(t.Task)
	} else {
		newTaskID = 0
	}
	if !groupSet {
		group = ""
	} else if !t.checkGroupCreated(group) {
		t.CreateGroup(group)
	}
	timestamp := time.Now()
	newTask = Task{ID: newTaskID, Item: message, Group: group, Timestamp: timestamp.Format(time.RFC3339)}
	t.Task = append(t.Task, newTask)
	if groupSet {
		err := t.addTaskToGroup(group, newTaskID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Edit a task
func (t *Tasks) EditTask(message, group string, id int, groupSet bool) error {
	existingTask := t.Task[id]
	if message == "" {
		message = existingTask.Item
	}
	timestamp := time.Now()
	if !groupSet {
		group = existingTask.Group
	} else if group == "" && existingTask.Group != "" {
		t.removeTaskFromGroup(id)
	}
	t.Task[id] = Task{ID: id, Item: message, Group: group, Timestamp: timestamp.Format(time.RFC3339)}
	return nil
}

func (t *Tasks) RemoveTask(id int) error {
	if id > len(t.Task)-1 {
		return fmt.Errorf("ID provided does not exist, must be 0-%d", len(t.Task)-1)
	}
	t.Task[id] = t.Task[len(t.Task)-1]
	t.Task = t.Task[:len(t.Task)-1]
	for i := range t.Task {
		t.Task[i].ID = i
	}
	return nil
}

func (t *Tasks) ListTasks(w io.Writer, group string, groupSet bool) {
	if groupSet {
		drawTasksGroup(w, t, group)
	} else {
		drawTasks(w, t)
	}
}

func FilePath(path string, pathSet bool) (string, error) {
	if pathSet {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := home + "/.td.json"
	return filePath, nil
}
