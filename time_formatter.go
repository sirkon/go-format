package format

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	strftime "github.com/lestrrat/go-strftime"
)

var deltaWords = []string{"year", "month", "week", "day", "hour", "minute", "second"}
var pluralReduction = map[string]string{
	"years": "year", "months": "month", "weeks": "week", "days": "day", "hours": "hour",
	"minutes": "minute", "seconds": "second",
}
var deltaIndexes = map[string]int{}

func init() {
	for i, word := range deltaWords {
		deltaIndexes[word] = i
	}
}

func deltaIndex(singular string) int {
	index, ok := deltaIndexes[singular]
	if !ok {
		panic(fmt.Sprintf("Unsupported singular word %s", singular))
	}
	return index
}

func restMatcher(rest []string, singular string, handler func(int) time.Time) (res time.Time, ok bool, err error) {
	if len(rest) == 0 {
		res = handler(0)
		return
	}
	num, err := strconv.ParseInt(rest[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("Expected to see a number, got %s instead", rest[0])
		return
	}

	if len(rest) == 1 {
		index := deltaIndex(singular)
		err = fmt.Errorf("Expected to see one of %s, got nothing instead", strings.Join(deltaWords[index:], ", "))
		return
	}

	probableSingular, ok := pluralReduction[strings.ToLower(rest[1])]
	if !ok {
		probableSingular = strings.ToLower(rest[1])
	}
	index, ok := deltaIndexes[probableSingular]
	if !ok {
		index = deltaIndex(singular)
		err = fmt.Errorf("Expected to see one of %s, got %s instead", strings.Join(deltaWords[index:], ", "), rest[1])
		return
	}
	if index < deltaIndex(singular) {
		index = deltaIndex(singular)
		err = fmt.Errorf(
			"Wrong delta parts precedence: shorter interval (week is shorter than month or year) must follow longers ones"+
				", i.e. you cannot write `1 day 5 months` and must use `5 months 1 day` form instead. In this case we got %s"+
				" while we were expecting one of %s", rest[1], strings.Join(deltaWords[index:], ", "))
		return
	}
	if probableSingular == singular {
		return handler(int(num)), true, nil
	}
	return handler(0), false, nil
}

// timeFormatter formatter
type timeFormatter time.Time

// MapDelta returnes shifted time object compared to the original value
func (d timeFormatter) MapDelta(delta string) (datetime time.Time, err error) {
	datetime = time.Time(d)

	delta = strings.TrimSpace(delta)
	if len(delta) == 0 {
		return
	}

	var sign int
	if strings.HasPrefix(delta, "+") {
		sign = 1
		delta = delta[1:]
	} else if strings.HasPrefix(delta, "-") {
		sign = -1
		delta = delta[1:]
	} else {
		err = fmt.Errorf("Expected + or - sign at the start of clarfying expression, got %s instead", delta)
		return
	}

	delta = strings.TrimLeft(delta, "+ \t\n\r")
	rest := strings.Fields(delta)
	var ok bool

	datetime, ok, err = restMatcher(rest, "year", func(k int) time.Time { return datetime.AddDate(sign*k, 0, 0) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}

	datetime, ok, err = restMatcher(rest, "month", func(k int) time.Time { return datetime.AddDate(0, sign*k, 0) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}

	datetime, ok, err = restMatcher(rest, "week", func(k int) time.Time { return datetime.AddDate(0, 0, 7*sign*k) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}

	datetime, ok, err = restMatcher(rest, "day", func(k int) time.Time { return datetime.AddDate(0, 0, sign*k) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}

	datetime, ok, err = restMatcher(rest, "hour",
		func(k int) time.Time { return datetime.Add(time.Duration(sign*k) * time.Second * 3600) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}

	datetime, ok, err = restMatcher(rest, "minute",
		func(k int) time.Time { return datetime.Add(time.Duration(sign*k) * time.Second * 60) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}

	datetime, ok, err = restMatcher(rest, "second",
		func(k int) time.Time { return datetime.Add(time.Duration(sign*k) * time.Second) })
	if err != nil {
		return
	}
	if ok {
		rest = rest[2:]
	}
	if len(rest) == 0 {
		return
	}
	err = fmt.Errorf("Unparsed rest of the delta expression left \033[31m%s\033[0m", rest)
	return
}

// Clarify formatter implementation
func (d timeFormatter) Clarify(delta string) (Formatter, error) {
	datetime, err := d.MapDelta(delta)
	if err != nil {
		return nil, err
	}
	return timeFormatter(datetime), err
}

// Format formatter implementation
func (d timeFormatter) Format(format string) (string, error) {
	return strftime.Format(format, time.Time(d))
}
