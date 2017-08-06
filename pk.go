package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	port := flag.Int("p", 0, "the port number to kill")
	flag.Parse()

	if *port == 0 {
		usageAndExit()
	}

	r, err := getProcs()
	if err != nil {
		log.Fatal(err)
	}

	pid, err := getPid(r, *port)
	if err != nil {
		log.Fatal(err)
	}

	if err = killPid(pid); err != nil {
		log.Fatal(err)
	}
}

func usageAndExit() {
	flag.Usage()
	os.Exit(1)
}
