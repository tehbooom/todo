package cmd

import (
	"bytes"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	cmd := NewRootCmd()
	actual := new(bytes.Buffer)
	cmd.SetOut(actual)
	cmd.SetErr(actual)
	err := cmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestExecute(t *testing.T) {
	actual := new(bytes.Buffer)
	err := Execute(actual)
	if err != nil {
		t.Error(err)
	}
}
