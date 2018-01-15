package format

import (
	"strconv"
)

// Values is easier than map[string]interface{}
type Values = map[string]interface{}

// Format function
func Format(format string, context Context) (string, error) {
	splitter := NewSplitter(format, context)
	res := ""
	for splitter.Split() {
		res += splitter.Text()
	}
	return res, splitter.Err()
}

// Formatp is a formatting with positional arguments
func Formatp(format string, a ...interface{}) string {
	bctx := NewContextBuilder()
	for i, value := range a {
		key := strconv.Itoa(i)
		bctx.Add(key, value)
	}
	ctx, err := bctx.Build()
	if err != nil {
		panic(err)
	}
	res, err := Format(format, ctx)
	if err != nil {
		panic(err)
	}
	return res
}

// Formatm is a formatting with keys from given map
func Formatm(format string, data Values) string {
	bctx := NewContextBuilder()
	for key, value := range data {
		bctx.Add(key, value)
	}
	ctx, err := bctx.Build()
	if err != nil {
		panic(err)
	}
	res, err := Format(format, ctx)
	if err != nil {
		panic(err)
	}
	return res
}
