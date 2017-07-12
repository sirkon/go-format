package format

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	builder := NewContextBuilder()
	builder.AddFormatter("date", TimeFormatter(time.Date(2016, 9, 10, 11, 12, 13, 0, time.UTC)))
	builder.AddFormatter("path", StringFormatter("/path/to/logs"))
	builder.AddFormatter("num", IntFormatter(12))
	context, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	res, err := Format("${path}/bos-k011/bos_srv-k011a.fss.log.${date | %Y%m%d }.gz", context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "/path/to/logs/bos-k011/bos_srv-k011a.fss.log.20160910.gz", res) {
		return
	}

	res, err = Format("${path}/bos-k011/bos_srv-k011a.fss.log.${date + 1 day | %Y%m%d }.gz", context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "/path/to/logs/bos-k011/bos_srv-k011a.fss.log.20160911.gz", res) {
		return
	}

	res, err = Format("${path}/bos-k011/bos_srv-k011a.fss.log.${date  1 day | %Y%m%d }.gz", context)
	if !assert.NotNil(t, err) {
		return
	}

	res, err = Format("${path}/bos-k011/bos_srv-k011a.fss.log.${date - 1 day | %Y%m%d }.gz", context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "/path/to/logs/bos-k011/bos_srv-k011a.fss.log.20160909.gz", res) {
		return
	}

	res, err = Format("${path}/bos-k011/bos_srv-k011a.fss.log.${date - 2 months 1 day | %Y%m%d }.gz", context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "/path/to/logs/bos-k011/bos_srv-k011a.fss.log.20160709.gz", res) {
		return
	}

	res, err = Format("${path}/bos-k011/bos_srv-k011a.fss.log.${date - 12 hours | %Y%m%d }.gz", context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "/path/to/logs/bos-k011/bos_srv-k011a.fss.log.20160909.gz", res) {
		return
	}

	res, err = Format(`${path}/bos-k011/bos_srv-k011a.fss.log.${date - 12 hours | "%Y%m%d %H%M%S" }.gz`, context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "/path/to/logs/bos-k011/bos_srv-k011a.fss.log.20160909 231213.gz", res) {
		return
	}

	res, err = Format("${ num | 04 }", context)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, "0012", res) {
		return
	}

	res, err = Format("$path$num", context)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "/path/to/logs12", res)

	res, err = Format("$path abc", context)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "/path/to/logs abc", res)
}

func TestClarify(t *testing.T) {
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatal(err)
	}
	date := TimeFormatter(time.Date(1982, 10, 19, 18, 22, 33, 0, moscow))
	date2, err := date.MapDelta("+ 1 year")
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, time.Time(date).AddDate(1, 0, 0), date2) {
		return
	}
}
