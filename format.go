package format

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

func resolveValue(prefix []string, bctx *ContextBuilder, value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Ptr:
		resolveValue(prefix, bctx, value.Elem())
	case reflect.Struct:
		iterateStruct(prefix, bctx, value)
		return true
	case reflect.Map:
		valueType := value.Type().Key()
		switch valueType.Kind() {
		case reflect.String:
			iterateStringMap(prefix, bctx, value)
			return true
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			iterateIntMap(prefix, bctx, value)
			return true
		default:
			stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
			if valueType.Implements(stringerType) {
				iterateStringerMap(prefix, bctx, value)
				return true
			}
		}
	}
	return false
}

func iterateStringerMap(prefix []string, builder *ContextBuilder, value reflect.Value) {
	keys := value.MapKeys()
	for _, keyData := range keys {
		keyValue := value.MapIndex(keyData)
		key := append(prefix, keyData.Interface().(fmt.Stringer).String())
		builder.Add(join(key), keyValue.Interface())
		resolveValue(key, builder, keyValue)
	}
}

func iterateIntMap(prefix []string, builder *ContextBuilder, value reflect.Value) {
	keys := value.MapKeys()
	for _, keyData := range keys {
		keyValue := value.MapIndex(keyData)
		key := append(prefix, fmt.Sprintf("%d", keyData.Interface()))
		builder.Add(join(key), keyValue.Interface())
		resolveValue(key, builder, keyValue)
	}
}

func iterateStringMap(prefix []string, builder *ContextBuilder, value reflect.Value) {
	keys := value.MapKeys()
	for _, keyData := range keys {
		keyValue := value.MapIndex(keyData)
		key := append(prefix, keyData.String())
		builder.Add(join(key), keyValue.Interface())
		resolveValue(key, builder, keyValue)
	}
}

func iterateStruct(prefix []string, builder *ContextBuilder, value reflect.Value) {
	for i := 0; i < value.NumField(); i++ {
		fieldMeta := value.Type().Field(i)
		fieldValue := value.Field(i)
		if fieldMeta.Anonymous {
			resolveValue(prefix, builder, fieldValue)
		} else {
			key := append(prefix, fieldMeta.Name)
			builder.Add(join(key), fieldValue.Interface())
			resolveValue(key, builder, fieldValue)
		}
	}
}

func join(src []string) string {
	return strings.Join(src, ".")
}

// Formatg is a formatting with type guessing
func Formatg(format string, data interface{}) string {
	bctx := NewContextBuilder()
	value := reflect.ValueOf(data)
	if !resolveValue(nil, bctx, value) {
		panic(fmt.Errorf("struct or map[string | integer | Stringer]X expected, got %T", data))
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

// Formatf is a formatting where values are taken form given func(string) string function
func Formatf(format string, data func(string) string) string {
	ctx := contextFunc(data)
	res, err := Format(format, ctx)
	if err != nil {
		panic(err)
	}
	return res
}
