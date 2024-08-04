package cmd

import (
	"bytes"
	"io"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehbooom/todo/internal/helpers"
)

func TestListCmd(t *testing.T) {
	tempPath := t.TempDir() + "/list.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}

	cmd := ListCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	cmd.SetArgs([]string{"-d", tempPath})
	err = cmd.Execute()
	if err != nil {
		t.Error(err)
	}
	output, err := io.ReadAll(actual)
	if err != nil {
		t.Error(err)
	}
	expectedOut := "" +
		"\x1b[100;32m┏\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┓\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;31m TIMESTAMP                 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m ID \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m TASK        \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m GROUP \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┣\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┫\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;33m 2024-06-23T11:19:56-04:00 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m  0 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m task 1 desc \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m grp1  \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[100;37m 2024-06-23T11:19:56-04:00 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[100;37m  1 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[100;37m task 2 desc \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[100;37m       \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┗\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┛\x1b[0m\n"

	assert.Equal(t, expectedOut, string(output))
}
