package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: pk [options] PORT\n")
		flag.PrintDefaults()
	}

	timeout := flag.Duration("t", time.Second*10, "timeout in Go time.Duration format")
	flag.Parse()

	flagArgs := flag.Args()
	port, err := strconv.Atoi(flagArgs[0])
	if err != nil {
		log.Fatalf("'%s' is not a valid port number", flagArgs[0])
	}

	pid, err := getPid(port, *timeout)
	if err != nil {
		log.Fatal(err)
	}

	if err = killPid(pid, *timeout); err != nil {
		log.Fatal(err)
	}
}
