package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/codingconcepts/pk"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: pk [options] PORT\n")
		flag.PrintDefaults()
	}

	timeout := flag.Duration("t", time.Second*10, "timeout in Go time.Duration format")
	debug := flag.Bool("d", false, "performs a dry-run")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	flagArgs := flag.Args()
	port, err := strconv.Atoi(flagArgs[0])
	if err != nil {
		log.Fatalf("'%s' is not a valid port number", flagArgs[0])
	}

	pid, err := pk.GetPid(port, *timeout)
	if err != nil {
		log.Fatal(err)
	}

	if *debug {
		fmt.Println(pid)
		return
	}

	if err = pk.KillPid(pid, *timeout); err != nil {
		log.Fatal(err)
	}
}
