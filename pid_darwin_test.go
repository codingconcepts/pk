// +build darwin

package main

import (
	"os"
	"testing"
)

func TestCollapseSpaces(t *testing.T) {
	input := "  this  is a   test "
	output := collapseSpaces(input)

	Equals(t, "this is a test", output)
}

func TestGetPIDWhenProcessListening(t *testing.T) {
	dump, err := os.Open("test/darwin_dump.txt")
	ErrorNil(t, err)
	defer dump.Close()

	pid, err := getPid(dump, 27017)
	ErrorNil(t, err)
	Equals(t, 34673, pid)
}

func TestGetPIDWhenProcessNotListening(t *testing.T) {
	dump, err := os.Open("test/darwin_dump.txt")
	ErrorNil(t, err)
	defer dump.Close()

	pid, err := getPid(dump, 60130)
	ErrorNil(t, err)
	Equals(t, 10893, pid)
}
