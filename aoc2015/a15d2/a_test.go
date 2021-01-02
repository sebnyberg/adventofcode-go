package a15d2

import (
	"bufio"
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/adventofcode/mathx"
	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

func Test_day(t *testing.T) {
	paper, ribbon := day()
	require.Equal(t, 1586300, paper)
	require.Equal(t, 3737498, ribbon)
	t.Fail()
}

func day() (paper, ribbon int) {
	f := must.Open("input")
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var l, w, h int
		fmt.Sscanf(sc.Text(), "%dx%dx%d", &l, &w, &h)

		paper += 2*(l*w+w*h+h*l) + mathx.Min(l*w, w*h, h*l)

		sides := []int{l, w, h}
		sort.Ints(sides)
		ribbon += 2*(sides[0]+sides[1]) + l*w*h
	}
	return paper, ribbon
}
