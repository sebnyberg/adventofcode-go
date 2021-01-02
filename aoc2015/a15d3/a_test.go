package a15d3

import (
	"bufio"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_day(t *testing.T) {
	deliveries := deliverPresents()
	var houses int
	for range deliveries {
		houses++
	}
	require.Equal(t, 0, houses)
}

type Point struct {
	X int
	Y int
}

func deliverPresents() map[Point]int {
	f := must.Open("input")
	// f := strings.NewReader("^>v<")
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanRunes)
	pos := Point{0, 0}
	deliveries := map[Point]int{pos: 1}
	for sc.Scan() {
		switch ch := sc.Text(); {
		case ch == ">":
			pos.X++
		case ch == "<":
			pos.X--
		case ch == "v":
			pos.Y--
		case ch == "^":
			pos.Y++
		default:
			log.Fatalf("bad character %v\n", string(ch))
		}
		fmt.Printf("+%v\n", deliveries)
		deliveries[pos]++
	}
	return deliveries
}
