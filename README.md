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

### Example

##### Positional parameters, you can use `${1}`, `${2}` to address specific parameter (by number)
```go
t := time.Date(2006, 1, 2, 3, 4, 5, 0, time.UTC)
res := format.Formatp("wake up at ${|%H:%M}, the call will be repeated ${} times", t, 5)
fmt.Println(res)
```

##### Named parameters via map[string]interface{}
```go
res := format.Formatm("${name} $count ${weight|1.1}", format.Values{
	"name": "name",
	"count": 12,
	"weight": 0.79,
})
fmt.Println(res)
```

##### Harder usage
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

Also, there's a simplified function with positional parameters



Date arithmetics allows following delta values:

* year/years
* month/months
* week/weeks
* day/days
* hour/hours
* minute/minutes
* second/seconds

Where *year* and *years* (and other couples) are equivalent (*5 years* is nicer than *5 year*, just like *1 year* is prefered over *1 years*).
There's a limitation though, longer period must precedes shorter one, i.e. the following expression is valid

```
2 year + 15 weeks + 1 second
```

while this one is invalid

```
1 week + 5 months
```
