/*
Copyright © 2024 Alec Carpenter
*/
package main

import (
	"os"

	"github.com/tehbooom/todo/cmd"
)

func main() {
	err := cmd.Execute(os.Stdout)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
