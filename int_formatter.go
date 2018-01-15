package format

import (
	"fmt"
	"strings"
)

// intFormatter formats int
type intFormatter struct {
	value int
}

// Clarify ...
func (f intFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for int8 formatter")
	}
	return f, nil
}

// Format ...
func (f intFormatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// uintFormatter formats uint
type uintFormatter struct {
	value uint
}

// Clarify ...
func (f uintFormatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for uint8 formatter")
	}
	return f, nil
}

// Format ...
func (f uintFormatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// i8Formatter ...
type i8Formatter struct {
	value int8
}

// Clarify formatter implementation
func (f i8Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for int8 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f i8Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// i16Formatter ...
type i16Formatter struct {
	value int16
}

// Clarify formatter implementation
func (f i16Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for int16 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f i16Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// i32Formatter ...
type i32Formatter struct {
	value int32
}

// Clarify formatter implementation
func (f i32Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for int32 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f i32Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// i64Formatter ...
type i64Formatter struct {
	value int64
}

// Clarify formatter implementation
func (f i64Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for int64 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f i64Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// u8Formatter ...
type u8Formatter struct {
	value uint8
}

// Clarify formatter implementation
func (f u8Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for uint8 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f u8Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// u16Formatter ...
type u16Formatter struct {
	value uint16
}

// Clarify formatter implementation
func (f u16Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for uint16 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f u16Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// u32Formatter ...
type u32Formatter struct {
	value uint32
}

// Clarify formatter implementation
func (f u32Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for uint32 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f u32Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}

// u64Formatter ...
type u64Formatter struct {
	value uint64
}

// Clarify formatter implementation
func (f u64Formatter) Clarify(a string) (Formatter, error) {
	if len(strings.TrimSpace(a)) != 0 {
		return f, fmt.Errorf("No clarification available for uint64 formatter")
	}
	return f, nil
}

// Format formatter implementation
func (f u64Formatter) Format(a string) (string, error) {
	return fmt.Sprintf("%"+a+"d", f.value), nil
}
