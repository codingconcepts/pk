package pk

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// GetPid returns the ID of the process that's exposing the given port.
func GetPid(port int, timeout time.Duration) (pid int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "lsof", "-t", "-i", fmt.Sprintf(":%d", port))
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("error fetching ports: %v", err)
	}

	clean := strings.Trim(buf.String(), " \n")
	return strconv.Atoi(clean)
}

// KillPid terminates a process by a given process ID.
func KillPid(pid int, timeout time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "kill", strconv.Itoa(pid))
	return cmd.Run()
}
