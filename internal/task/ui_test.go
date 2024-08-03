/*
Copyright © 2024 Alec Carpenter
*/
package task

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawTable(t *testing.T) {
	var task = &Tasks{
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
	var buf bytes.Buffer

	drawTasks(&buf, task)

	output := buf.String()
	expectedOut := "" +
		"\x1b[100;32m┏\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┓\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;31m TIMESTAMP                 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m ID \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m TASK        \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m GROUP \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┣\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┫\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;33m 2024-06-23T11:19:56-04:00 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m  1 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m task 1 desc \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m grp1  \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┗\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┛\x1b[0m\n"

	assert.Equal(t, expectedOut, output)
}

func TestDrawGroupTable(t *testing.T) {
	var task = &Tasks{
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
	var buf bytes.Buffer

	drawTasksGroup(&buf, task, task.Task[0].Group)

	output := buf.String()

	expectedOut := "" +
		"\x1b[100;32m┏\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┓\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;31m TIMESTAMP                 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m ID \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m TASK        \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┣\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┫\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;33m 2024-06-23T11:19:56-04:00 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m  1 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m task 1 desc \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┗\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┛\x1b[0m\n"

	assert.Equal(t, expectedOut, output)
}
