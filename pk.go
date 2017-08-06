package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected port argument")
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("'%s' is not a valid port number", os.Args[1])
	}

	pid, err := getPid(port)
	if err != nil {
		log.Fatal(err)
	}

	if err = killPid(pid); err != nil {
		log.Fatal(err)
	}
}
