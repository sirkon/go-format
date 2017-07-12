package format

import (
	"fmt"
	"strings"
)

// IntFormatter ...
type IntFormatter int

// Clarify formatter implementation
func (f IntFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for string formatters")
	}
	return f, nil
}

// Format formatter implementation
func (f IntFormatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", int(f)), nil
}
