package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_locateABuck(t *testing.T) {
	tests := []struct {
		rest string
		want int
	}{
		{
			rest: "$$abcd",
			want: -1,
		},
		{
			rest: "$abcd",
			want: 0,
		},
		{
			rest: "abcdef",
			want: -1,
		},
		{
			rest: "   $a",
			want: 3,
		},
		{
			rest: "$$abcde $a",
			want: 8,
		},
		{
			rest: " ${num} ${float|1.1}",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.rest, func(t *testing.T) {
			loc := locateABuck(tt.rest)
			if loc >= 0 {
				t.Log(tt.rest[loc:])
			}
			assert.Equalf(t, tt.want, loc, "locateABuck(%v)", tt.rest)
		})
	}
}
