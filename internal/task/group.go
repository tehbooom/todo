/*
Copyright © 2024 Alec Carpenter
*/
package task

import (
	"os"
	"slices"
	"sort"
)

type Group struct {
	Name    string `json:"name"`
	TaskIDs []int  `json:"tasks"`
}

func (t *Tasks) addTaskToGroup(name string, task int) error {
	groupExist := t.checkGroupCreated(name)

	if !groupExist {
		err := t.CreateGroup(name)
		if err != nil {
			return err
		}
	}

	t.Task[task].Group = name

	for i, groupItem := range t.Groups {
		if name == groupItem.Name {
			if slices.Contains(t.Groups[i].TaskIDs, task) {
				return nil
			}
			t.Groups[i].TaskIDs = append(t.Groups[i].TaskIDs, task)
			sort.Ints(t.Groups[i].TaskIDs)
		}
	}

	return nil
}

func (t *Tasks) removeTaskFromGroup(id int) error {
	groupName := t.Task[id].Group

	for i, groupItem := range t.Groups {
		if groupName == groupItem.Name {
			for j, taskID := range t.Groups[i].TaskIDs {
				if taskID == id {
					t.Groups[i].TaskIDs = append(t.Groups[i].TaskIDs[:j], t.Groups[i].TaskIDs[j+1:]...)
				}
			}
		}
	}

	t.Task[id].Group = ""

	return nil
}

func (t *Tasks) RemoveGroup(name string) error {
	var taskIDs []int
	var groupIndex int
	for i, group := range t.Groups {
		if name == group.Name {
			taskIDs = group.TaskIDs
			groupIndex = i
		}
	}

	// Remove tasks from task list
	for i, id := range taskIDs {
		positionalID := id - i
		err := t.RemoveTask(positionalID)
		if err != nil {
			return err
		}
	}

	// Remove the group
	t.Groups[groupIndex] = t.Groups[len(t.Groups)-1]
	t.Groups = t.Groups[:len(t.Groups)-1]
	return nil
}

func (t *Tasks) CreateGroup(name string) error {
	for _, group := range t.Groups {
		if name == group.Name {
			return nil
		}
	}

	var group Group
	group.Name = name
	group.TaskIDs = []int{}
	t.Groups = append(t.Groups, group)
	return nil
}

func (t *Tasks) ListGroups() {
	drawListGroups(os.Stdout, t)
}

func (t *Tasks) checkGroupCreated(name string) bool {
	for _, group := range t.Groups {
		if name == group.Name {
			return true
		}
	}
	return false
}
