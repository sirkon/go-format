package format

import (
	"fmt"
	"strings"
)

// StringFormatter formatting primitive
type StringFormatter string

// Clarify formatter implementation
func (f StringFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for string formatters")
	}
	return f, nil
}

// Format formatter implementation
func (f StringFormatter) Format(a string) (string, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return string(f), fmt.Errorf("No clarification available for string formatters")
	}
	return string(f), nil
}
