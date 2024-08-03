package cmd

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCmd(t *testing.T) {
	path := initializeTask(t)
	actual := new(bytes.Buffer)
	NewRootCmd().SetOut(actual)
	NewRootCmd().SetErr(actual)
	log.SetOutput(actual)
	NewRootCmd().SetArgs([]string{"list", "-d", path})
	NewRootCmd().Execute()
	expectedOut := "" +
		"\x1b[100;32m┏\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┳\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┓\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;31m TIMESTAMP                 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m ID \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m TASK        \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;31m GROUP \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┣\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┫\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;33m 2024-06-23T11:19:56-04:00 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m  0 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m task 1 desc \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m grp1  \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┣\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m╋\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┫\x1b[0m\n" +
		"\x1b[100;32m┃\x1b[0m\x1b[40;33m 2024-06-23T11:19:56-04:00 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m  1 \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m task 2 desc \x1b[0m\x1b[100;32m┃\x1b[0m\x1b[40;33m       \x1b[0m\x1b[100;32m┃\x1b[0m\n" +
		"\x1b[100;32m┗\x1b[0m\x1b[100;32m━━━━━━━━━━━━━━━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━━━━━━━\x1b[0m\x1b[100;32m┻\x1b[0m\x1b[100;32m━━━━━━━\x1b[0m\x1b[100;32m┛\x1b[0m\n"
	assert.Equal(t, actual.String(), expectedOut)
}
