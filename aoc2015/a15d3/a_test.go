package a15d3

import (
	"bufio"
	"log"
	"testing"

	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

func Test_day(t *testing.T) {
	deliveries := deliverPresents()
	var houses int
	for range deliveries {
		houses++
	}
	require.Equal(t, 2631, houses)
}

type Point struct {
	X int
	Y int
}

func deliverPresents() map[Point]int {
	f := must.Open("input")
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanRunes)
	santaPos := []Point{{0, 0}, {0, 0}}
	deliveries := map[Point]int{{0, 0}: 2}
	for i := 0; sc.Scan(); i++ {
		move(&santaPos[i%2], sc.Text())
		deliveries[santaPos[i%2]]++
	}
	return deliveries
}

func move(pos *Point, ch string) {
	switch ch {
	case ">":
		pos.X++
	case "<":
		pos.X--
	case "v":
		pos.Y--
	case "^":
		pos.Y++
	default:
		log.Fatalf("bad character %v\n", string(ch))
	}
}
