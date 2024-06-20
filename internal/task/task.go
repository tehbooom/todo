/*
Copyright © 2024 Alec Carpenter
*/
package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Tasks struct {
	Task []Task `json:"tasks"`
}

type Task struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

func readTasks(path string) (Tasks, error) {
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

func writeTasks(path string, item Task) error {
	existingTasks, err := readTasks(path)
	if err != nil {
		return err
	}
	existingTasks.Task = append(existingTasks.Task, item)

	jsonTasks, err := json.Marshal(existingTasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonTasks, 0600)
	if err != nil {
		return err
	}

	return nil
}

func AddTask(path, message string) (Task, error) {
	var newTask Task
	var newTaskID int

	existingsTasks, err := readTasks(path)
	if err != nil {
		return newTask, err
	}
	if len(existingsTasks.Task) != 0 {
		lastTask := existingsTasks.Task[len(existingsTasks.Task)-1]

		newTaskID = lastTask.ID + 1
	} else {
		newTaskID = 0
	}
	newTask = Task{ID: newTaskID, Item: message}
	err = writeTasks(path, newTask)
	if err != nil {
		return newTask, err
	}
	return newTask, nil
}

func RemoveTask(path string, id int) error {
	taskList, err := readTasks(path)
	if err != nil {
		return err
	}
	if id > len(taskList.Task)-1 {
		fmt.Println("ID number provided does not exist")
		err = ListTasks(path)
		if err != nil {
			return err
		}
		return fmt.Errorf("please provide a valid ID number")
	}
	taskList.Task[id] = taskList.Task[len(taskList.Task)-1]
	taskList.Task = taskList.Task[:len(taskList.Task)-1]
	for i := range taskList.Task {
		taskList.Task[i].ID = i
	}
	newTaskList, err := json.Marshal(taskList)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, newTaskList, 0600)
	if err != nil {
		return err
	}

	return nil
}

func ListTasks(path string) error {
	taskList, err := readTasks(path)
	if err != nil {
		return err
	}
	var taskItemList []string
	for _, task := range taskList.Task {
		taskWithID := fmt.Sprintf("%d: %s", task.ID, task.Item)
		taskItemList = append(taskItemList, taskWithID)
	}
	allTasks := strings.Join(taskItemList, "\n")
	fmt.Println(allTasks)
	return nil
}

func TaskFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := home + "/.td.json"
	return filePath, nil
}
