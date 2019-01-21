package format

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	builder := NewContextBuilder()
	builder.AddFormatter("date", timeFormatter(time.Date(2016, 9, 10, 11, 12, 13, 0, time.UTC)))
	builder.AddFormatter("path", stringFormatter("/path/to/logs"))
	builder.AddFormatter("num", i8Formatter{
		int8(12),
	})
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

	res = Formatp("${+1 day|%Y-%m-%d}", time.Date(2018, 1, 15, 0, 0, 0, 0, time.UTC))
	require.Equal(t, "2018-01-16", res)
}

func TestClarify(t *testing.T) {
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatal(err)
	}
	date := timeFormatter(time.Date(1982, 10, 19, 18, 22, 33, 0, moscow))
	date2, err := date.MapDelta("+ 1 year")
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, time.Time(date).AddDate(1, 0, 0), date2) {
		return
	}
}

func TestFormatp(t *testing.T) {
	require.Equal(t, "a 2 4.5 2 2018", Formatp("${} ${} ${|1.1} ${1} ${|%Y}", "a", 2, 4.5, time.Date(2018, 10, 19, 18, 0, 5, 0, time.UTC)))
	require.Equal(t, "a a", Formatp("$ $0", "a"))
}

func TestFormatm(t *testing.T) {
	require.Equal(t, "a 2 4.5", Formatm("${key} ${num} ${float|1.1}", map[string]interface{}{
		"key":   "a",
		"num":   2,
		"float": 4.5,
	}))
}

func TestFormatg(t *testing.T) {
	require.Equal(t, "a 2 4.5", Formatg("${key} ${num} ${float|1.1}", map[string]interface{}{
		"key":   "a",
		"num":   2,
		"float": 4.5,
	}))

	type t1 struct {
		A string
		B int
	}
	s1 := t1{
		A: "a1",
		B: 2,
	}
	require.Equal(t, "a1→2", Formatg("${A}→$B", s1))

	var s2 struct {
		t1
		C float64
		D t1
	}
	s2.t1 = s1
	s2.C = 0.5
	s2.D = s1
	require.Equal(t, "a1 2 0.5 2 a1", Formatg("$A $B ${C|1.1} ${D.B} ${D.A}", s2))
}

func TestReadmeFormatp(t *testing.T) {
	res := Formatp("$ ${} $1 ${0}", 1, 2)
	require.Equal(t, "1 2 2 1", res)
}

func TestReamdeFormatm(t *testing.T) {
	res := Formatm("${name} $count ${weight|1.2}", Values{
		"name":   "name",
		"count":  12,
		"weight": 0.79,
	})
	require.Equal(t, "name 12 0.79", res)
}

type T struct {
	A     string
	Field int
}

func TestReadmeFormatg(t *testing.T) {
	var s = T{
		A:     "str",
		Field: 12,
	}
	var d struct {
		F     T
		Entry float64
	}
	d.F = s
	d.Entry = 0.5
	res := Formatg("$F.A $F.Field $Entry", d)
	require.Equal(t, "str 12 0.500000", res)

	v := map[int]string{
		1:  "bc",
		12: "bd",
	}
	res = Formatg("$1-$12", v)
	require.Equal(t, "bc-bd", res)
}

func TestReadmeDateArithmetics(t *testing.T) {
	tm := time.Date(2018, 1, 18, 22, 57, 37, 12, time.UTC)
	res := Formatm("${ date + 1 day | %Y-%m-%d %H:%M:%S }", Values{
		"date": tm,
	})
	require.Equal(t, "2018-01-19 22:57:37", res)
}

func TestFormatf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		data   func(string) string
		want   string
	}{
		{
			name:   "typical",
			format: "${name} ${value} ${location} - ${fake}",
			data: func(s string) string {
				switch s {
				case "name":
					return "Denis"
				case "value":
					return "1k"
				case "location":
					return "cbx"
				default:
					return "<nil>"
				}
			},
			want: "Denis 1k cbx - <nil>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Formatf(tt.format, tt.data); got != tt.want {
				t.Errorf("Formatf() = %v, want %v", got, tt.want)
			}
		})
	}
}
