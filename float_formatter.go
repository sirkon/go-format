package format

import (
	"fmt"
	"strings"
)

// f32Formatter ...
type f32Formatter struct {
	value float32
}

// Clarify formatter implementation
func (f f32Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for float32 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f f32Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"f", f.value), nil
}

// f64Formatter ...
type f64Formatter struct {
	value float64
}

// Clarify formatter implementation
func (f f64Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for float64 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f f64Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"f", f.value), nil
}
