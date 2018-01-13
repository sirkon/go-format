package format

import (
	"fmt"
	"strconv"
	"time"
)

// Format function
func Format(format string, context Context) (string, error) {
	splitter := NewSplitter(format, context)
	res := ""
	for splitter.Split() {
		res += splitter.Text()
	}
	return res, splitter.Err()
}

// Formatf is a Python-like formatting
func Formatf(format string, a ...interface{}) string {
	bctx := NewContextBuilder()
	for i, value := range a {
		key := strconv.Itoa(i)
		switch v := value.(type) {
		case int8:
			bctx.AddInteger(key, v)
		case int16:
			bctx.AddInteger(key, v)
		case int32:
			bctx.AddInteger(key, v)
		case int64:
			bctx.AddInteger(key, v)
		case uint8:
			bctx.AddInteger(key, v)
		case uint16:
			bctx.AddInteger(key, v)
		case uint32:
			bctx.AddInteger(key, v)
		case uint64:
			bctx.AddInteger(key, v)
		case float32:
			bctx.AddFloat(key, v)
		case float64:
			bctx.AddFloat(key, v)
		case string:
			bctx.AddString(key, v)
		case time.Time:
			bctx.AddTime(key, v)
		case Formatter:
			bctx.AddFormatter(key, v)
		case fmt.Stringer:
			bctx.AddString(key, v.String())
		default:
			bctx.AddValue(key, value)
		}
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
