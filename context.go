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
		err = fmt.Errorf("Unknown formatter `%s`, only these are available:%s\n\n", name, strings.Join(formatters, "\n"))
	}
	return
}
