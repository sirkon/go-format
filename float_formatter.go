package format

import (
	"fmt"
	"strings"
)

// floatFormatter ...
type floatFormatter struct {
	value interface{}
}

// Clarify formatter implementation
func (f floatFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for string formatters")
	}
	return f, nil
}

// Format formatter implementation
func (f floatFormatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"f", f.value), nil
}
