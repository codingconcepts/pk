package pk

import (
	"fmt"
	"strconv"
	"strings"
)

func stringToPids(s string) ([]int, error) {
	parts := strings.Split(s, "\n")

	var intParts []int
	for _, part := range parts {
		p, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("parsing pid: %w", err)
		}

		intParts = append(intParts, p)
	}

	return intParts, nil
}
