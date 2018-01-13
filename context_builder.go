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
	return c.AddFormatter(name, stringFormatter(value))
}

// AddNumber adds integer formatter
func (c *ContextBuilder) AddInteger(name string, value interface{}) *ContextBuilder {
	switch value.(type) {
	case int8:
	case int16:
	case int32:
	case int64:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	default:
		panic(fmt.Errorf("can only consume int or uint types into integer formatter, got %T", value))
	}
	return c.AddFormatter(name, intFormatter{
		value: value,
	})
}

// AddFloat adds floating point number formatter
func (c *ContextBuilder) AddFloat(name string, value interface{}) *ContextBuilder {
	switch value.(type) {
	case float32:
	case float64:
	default:
		panic(fmt.Errorf("can only consume float32 or float32 into float formatter, got %T", value))
	}
	return c.AddFormatter(name, floatFormatter{
		value: value,
	})
}

// AddTime adds time formatter
func (c *ContextBuilder) AddTime(name string, datetime time.Time) *ContextBuilder {
	return c.AddFormatter(name, timeFormatter(datetime))
}

// AddValue adds value formatter
func (c *ContextBuilder) AddValue(name string, value interface{}) *ContextBuilder {
	return c.AddFormatter(name, valueFormatter{
		value: value,
	})
}

// Build retrieves context object from the builder
func (c *ContextBuilder) Build() (Context, error) {
	if c.err != nil {
		return nil, c.err
	}
	return contextFromBuilder(c.formatters), nil
}
