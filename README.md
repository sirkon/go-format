# go-format
[![Build Status](https://travis-ci.org/sirkon/go-format.svg?branch=master)](https://travis-ci.org/sirkon/go-format)

Simple string formatting tool with date arithmetics and format (strftime) support

### Installation
```bash
go get github.com/sirkon/go-format
```
or
```bash
dep ensure -add github.com/sirkon/go-format
```

### Examples

##### Positional parameters, you can use `${1}`, `${2}` to address specific parameter (by number)
```go
res := format.Formatp("$ ${} $1 ${0}", 1, 2)
// res = "1 2 2 1"
```


##### Named parameters via map[string]interface{}
```go
res := format.Formatm("${name} $count ${weight|1.2}", format.Values{
	"name": "name",
	"count": 12,
	"weight": 0.79,
})
// res = "name 12 0.79"
``` 

##### Named parameters via type guesses
```go
var s struct {
	A string
	Field int
}
s.A = "str"
s.Field = 12
res := format.Formatg("$A $Field", s)
```
Substructures are supported. Also, any kind of map with keys being one of
1) string
2) any kind of integer, signed or unsigned
3) `fmt.Stringer`
is supported as well.

```go
type t struct {
	A     string
	Field int
}
var s = t{
	A:     "str",
	Field: 12,
}
var d struct {
	F     t
	Entry float64
}
d.F = s
d.Entry = 0.5
res := format.Formatg("$F.A $F.Field $Entry", d)
// res = "str 12 0.500000"
```

```go
v := map[int]string{
	1: "bc",
	12: "bd"
}
res := format.Formatg("$1-$12")
// res = "bc-bc"
```

##### Date arithmetics
```go
t := time.Date(2018, 1, 18, 22, 57, 37, 12, time.UTC)
res := format.Formatm("${ date + 1 day | %Y-%m-%d %H:%M:%S }", format.Values{
	"date": t,
})
// res = "2018-01-19 22:57:37"
```
Date arithmetics allows following expression values:

* year/years
* month/months
* week/weeks
* day/days
* hour/hours
* minute/minutes
* second/seconds

Where *year* and *years* (and other couples) are equivalent (*5 years* is nicer than *5 year*, just like *1 year* is prefered over *1 years*).
There's a limitation though, longer period must precedes shorter one, i.e. the following expression is valid.

Examples of valid time delta expressions:
```
+ 1 year + 5 weeks + 3 days + 12 seconds
+ 25 years + 3 months
- 17 days - 16 hours
+ 2 year + 15 weeks + 1 second
```

invalid time delta expression
```
+ 1 week + 5 months
```

month must precedes a week


##### Low level usage, how it is doing in the background
```go
package main


import (
	"fmt"
	"time"
	
	"github.com/sirkon/go-format"
)


func main() {
	bctx := format.NewContextBuilder()
	bctx.AddString("name", "explosion")
	bctx.AddInt("count", 15)
	bctx.AddTime("time", time.Date(2006, 1, 2, 3, 4, 5, 0, time.UTC))
	ctx, err := bctx.Build()
	if err != nil {
		panic(err)
	}

	res, err := format.Format(
		"${name} will be registered by $count independent sources in ${ time + 1 day | %Y-%m-%d } at ${ time | %H:%M:%S }",
		ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

It is probably to be used for most typical usecases.  


 
 

