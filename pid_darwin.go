// +build darwin

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"bufio"

	"strings"

	"github.com/pkg/errors"
)

var (
	spaceCollapsePattern = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

func getProcs() (r io.Reader, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "lsof", "-Pn", "-i4", "-i6")

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "error fetching ports")
	}

	return buf, nil
}

func getPid(r io.Reader, port int) (pid int, err error) {
	portStr := fmt.Sprintf(":%d", port)

	lines := bufio.NewScanner(r)
	for lines.Scan() {
		line := collapseSpaces(lines.Text())
		parts := strings.Split(line, " ")

		if strings.HasSuffix(parts[8], portStr) {
			return strconv.Atoi(parts[1])
		}
	}

	return 0, fmt.Errorf("no pid found for port %d", port)
}

func killPid(pid int) (err error) {
	cmd := exec.Command("kill", strconv.Itoa(pid))
	return cmd.Run()
}

func collapseSpaces(input string) (output string) {
	return strings.Trim(spaceCollapsePattern.ReplaceAllString(input, " "), " ")
}
