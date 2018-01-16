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

// AddInt adds int formatter
func (c *ContextBuilder) AddInt(name string, value int) *ContextBuilder {
	return c.AddFormatter(name, intFormatter{
		value: value,
	})
}

// AddUint adds uint formatter
func (c *ContextBuilder) AddUint(name string, value uint) *ContextBuilder {
	return c.AddFormatter(name, uintFormatter{
		value: value,
	})
}

// AddInt8 adds int8 formatter
func (c *ContextBuilder) AddInt8(name string, value int8) *ContextBuilder {
	return c.AddFormatter(name, i8Formatter{
		value: value,
	})
}

// AddInt16 adds int16 formatter
func (c *ContextBuilder) AddInt16(name string, value int16) *ContextBuilder {
	return c.AddFormatter(name, i16Formatter{
		value: value,
	})
}

// AddInt32 adds int32 formatter
func (c *ContextBuilder) AddInt32(name string, value int32) *ContextBuilder {
	return c.AddFormatter(name, i32Formatter{
		value: value,
	})
}

// AddInt64 adds int64 formatter
func (c *ContextBuilder) AddInt64(name string, value int64) *ContextBuilder {
	return c.AddFormatter(name, i64Formatter{
		value: value,
	})
}

// AddUint8 adds uint8 formatter
func (c *ContextBuilder) AddUint8(name string, value uint8) *ContextBuilder {
	return c.AddFormatter(name, u8Formatter{
		value: value,
	})
}

// AddUint16 adds uint16 formatter
func (c *ContextBuilder) AddUint16(name string, value uint16) *ContextBuilder {
	return c.AddFormatter(name, u16Formatter{
		value: value,
	})
}

// AddUint32 adds uint32 formatter
func (c *ContextBuilder) AddUint32(name string, value uint32) *ContextBuilder {
	return c.AddFormatter(name, u32Formatter{
		value: value,
	})
}

// AddUint64 adds uint64 formatter
func (c *ContextBuilder) AddUint64(name string, value uint64) *ContextBuilder {
	return c.AddFormatter(name, u64Formatter{
		value: value,
	})
}

// AddFloat adds floating point number formatter
func (c *ContextBuilder) AddFloat32(name string, value float32) *ContextBuilder {
	return c.AddFormatter(name, f32Formatter{
		value: value,
	})
}

// AddFloat adds floating point number formatter
func (c *ContextBuilder) AddFloat64(name string, value float64) *ContextBuilder {
	return c.AddFormatter(name, f64Formatter{
		value: value,
	})
}

// AddTime adds time formatter
func (c *ContextBuilder) AddTime(name string, datetime time.Time) *ContextBuilder {
	return c.AddFormatter(name, timeFormatter(datetime))
}

// AddValue adds formatter to format it as %...v
func (c *ContextBuilder) AddValue(name string, value interface{}) *ContextBuilder {
	return c.AddFormatter(name, valueFormatter{
		value: value,
	})
}

// Add adds formatter with type guessing
func (c *ContextBuilder) Add(name string, value interface{}) *ContextBuilder {
	switch v := value.(type) {
	case int:
		c.AddInt(name, v)
	case uint:
		c.AddUint(name, v)
	case int8:
		c.AddInt8(name, v)
	case int16:
		c.AddInt16(name, v)
	case int32:
		c.AddInt32(name, v)
	case int64:
		c.AddInt64(name, v)
	case uint8:
		c.AddUint8(name, v)
	case uint16:
		c.AddUint16(name, v)
	case uint32:
		c.AddUint32(name, v)
	case uint64:
		c.AddUint64(name, v)
	case float32:
		c.AddFloat32(name, v)
	case float64:
		c.AddFloat64(name, v)
	case string:
		c.AddString(name, v)
	case time.Time:
		c.AddTime(name, v)
	case Formatter:
		c.AddFormatter(name, v)
	case fmt.Stringer:
		c.AddString(name, v.String())
	default:
		c.AddValue(name, v)
	}
	return c
}

// Build retrieves context object from the builder
func (c *ContextBuilder) Build() (Context, error) {
	if c.err != nil {
		return nil, c.err
	}
	return contextFromBuilder(c.formatters), nil
}
