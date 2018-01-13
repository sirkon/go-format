package format

import (
	"fmt"
	"strings"
)

// intFormatter ...
type intFormatter struct {
	value interface{}
}

// Clarify formatter implementation
func (f intFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for string formatters")
	}
	return f, nil
}

// Format formatter implementation
func (f intFormatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}
