package format

import (
	"fmt"
	"time"
)

// ContextBuilder formatting context builder
type ContextBuilder struct {
	formatters map[string]Formatter
	err        error
}

// NewContextBuilder ...
func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{
		formatters: map[string]Formatter{},
	}
}

// AddFormatter ...
func (c *ContextBuilder) AddFormatter(name string, formatter Formatter) *ContextBuilder {
	if c.err != nil {
		return c
	}
	_, ok := c.formatters[name]
	if ok {
		c.err = fmt.Errorf("Attemtp to redefine formatter %s", name)
		return c
	}
	c.formatters[name] = formatter
	return c
}

// AddString adds string formatter
func (c *ContextBuilder) AddString(name, value string) *ContextBuilder {
	return c.AddFormatter(name, StringFormatter(value))
}

// AddInt adds int formatter
func (c *ContextBuilder) AddInt(name string, value int) *ContextBuilder {
	return c.AddFormatter(name, IntFormatter(value))
}

// AddTime adds time formatter
func (c *ContextBuilder) AddTime(name string, datetime time.Time) *ContextBuilder {
	return c.AddFormatter(name, TimeFormatter(datetime))
}

// BuildContext retrieves context object from the builder
func (c *ContextBuilder) BuildContext() (Context, error) {
	if c.err != nil {
		return nil, c.err
	}
	return contextFromBuilder(c.formatters), nil
}
