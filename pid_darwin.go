// +build darwin

package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func getPid(port int) (pid int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "lsof", "-t", "-i", fmt.Sprintf(":%d", port))

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return 0, errors.Wrap(err, "error fetching ports")
	}

	clean := strings.Trim(buf.String(), " \n")
	return strconv.Atoi(clean)
}

func killPid(pid int) (err error) {
	cmd := exec.Command("kill", strconv.Itoa(pid))
	return cmd.Run()
}
