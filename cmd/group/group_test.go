package group

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehbooom/todo/internal/helpers"
)

func TestGroupCmd(t *testing.T) {
	tempPath := t.TempDir() + "/ls_group.json"
	err := helpers.InitializeTask(tempPath)
	if err != nil {
		t.Error(err)
	}
	cmd := GroupCmd()
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
		"\x1b[100;32m┏\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━━━━━━\x1b[0m\x1b[100;32m┓\x1b[0m\n\x1b[100;32m┃\x1b[0m\x1b[40;31m GROUP \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m # OF TASKS \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┣\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━━━━━━\x1b[0m\x1b[100;32m┫\x1b[0m\n\x1b[100;32m┃\x1b[0m\x1b[40;33m grp1  \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m          1 \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┗\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━━━━━━\x1b[0m\x1b[100;32m┛\x1b[0m\n"
	assert.Equal(t, expectedOut, string(output))
}
