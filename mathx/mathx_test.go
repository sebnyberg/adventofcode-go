package mathx_test

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/adventofcode/mathx"
	"github.com/stretchr/testify/require"
)

func Test_Min(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{3, 1}, 1},
		{[]int{3}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, mathx.Min(tc.in...))
		})
	}
}
