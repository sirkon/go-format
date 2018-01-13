package format

import (
	"fmt"
	"strings"
)

// valueFormatter ...
type valueFormatter struct {
	value interface{}
}

// Clarify formatter implementation
func (f valueFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for value formatters")
	}
	return f, nil
}

// Format formatter implementation
func (f valueFormatter) Format(a string) (string, error) {
	return fmt.Sprintf("%#"+a+"v", f.value), nil
}
