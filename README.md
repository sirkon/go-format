# go-format
Simple string formatting tool with date arithmetics and format (strftime) support

### Retrieval
```bash
go get github.com/sirkon/go-format
```

### Example

```go
package main

import (
	"fmt"
	"time"

	format "github.com/sirkon/go-format"
)

func main() {
	bctx := format.NewContextBuilder()
	bctx.AddString("name", "explosion")
	bctx.AddInteger("count", 15)
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

```go
package main

import (
 	"fmt"
 	"time"
 
 	format "github.com/sirkon/go-format"
 )

func main() {
	res := format.Formatf("wake up at ${|%H:%M}, the call will be repeated ${} times", time.Date(2006, 1, 2, 3, 4, 5, 0, time.UTC), 5)
	fmt.Println(res)
}
```

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
