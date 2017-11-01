package pk

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// GetPid returns the ID of the process that's exposing the given port.
func GetPid(port int, timeout time.Duration) (pid int, err error) {
	netstat := exec.Command("netstat", "-a", "-n", "-o")
	findstr := exec.Command("findstr", fmt.Sprintf("%d", port))
	findstr.Stdin, _ = netstat.StdoutPipe()

	buf := new(bytes.Buffer)
	findstr.Stdout = buf
	if err = findstr.Start(); err != nil {
		return 0, fmt.Errorf("error starting findstr: %v", err)
	}
	if err = netstat.Run(); err != nil {
		return 0, fmt.Errorf("error running netstat: %v", err)
	}
	if err = findstr.Wait(); err != nil {
		return 0, fmt.Errorf("error waiting on findstr: %v", err)
	}

	firstLine := strings.Split(buf.String(), "\n")[0]
	columns := strings.Split(firstLine, " ")

	id := columns[len(columns)-1]
	clean := strings.Trim(id, " \r\n")
	return strconv.Atoi(clean)
}

// KillPid terminates a process by a given process ID.
func KillPid(pid int, timeout time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "taskkill", "/F", "/PID", strconv.Itoa(pid))
	return cmd.Run()
}
