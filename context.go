package format

import (
	"fmt"
	"sort"
	"strings"
)

// Context ...
type Context interface {
	GetFormatter(name string) (Formatter, error)
}

type contextFromBuilder map[string]Formatter

func (ctx contextFromBuilder) GetFormatter(name string) (res Formatter, err error) {
	res, ok := ctx[name]
	if !ok {
		formatters := []string{}
		for key := range ctx {
			formatters = append(formatters, key)
		}
		sort.Sort(sort.StringSlice(formatters))
		for i, key := range formatters {
			formatters[i] = "\t" + key
		}
		err = fmt.Errorf(
			"unknown formatter `%s`, only these are available:\n%s\n",
			name,
			strings.Join(formatters, "\n"),
		)
	}
	return
}

// contextFunc implementation of context using given func as a source of information
type contextFunc func(string) string

func (f contextFunc) GetFormatter(name string) (Formatter, error) {
	value := f(name)
	return stringFormatter(value), nil
}
