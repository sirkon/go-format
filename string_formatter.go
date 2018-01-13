package format

import (
	"fmt"
	"strings"
)

// stringFormatter formatting primitive
type stringFormatter string

// Clarify formatter implementation
func (f stringFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for string formatters")
	}
	return f, nil
}

// Format formatter implementation
func (f stringFormatter) Format(a string) (string, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return string(f), fmt.Errorf("No clarification available for string formatters")
	}
	return string(f), nil
}
