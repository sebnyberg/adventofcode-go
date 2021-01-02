package a15d6

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/sebnyberg/adventofcode/mathx"
	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

func Test_day(t *testing.T) {
	// Part 1
	lights := setUpLights1(must.Open("input"))
	var turnedOn int
	for _, row := range lights {
		for _, light := range row {
			if light {
				turnedOn++
			}
		}
	}
	require.Equal(t, 400410, turnedOn)

	// Part 2
	lightBrightness := setUpLights2(must.Open("input"))
	var totalBrightness int
	for _, row := range lightBrightness {
		for _, brightness := range row {
			totalBrightness += brightness
		}
	}
	require.Equal(t, 15343601, totalBrightness)
}

type Point struct {
	x int
	y int
}

func setUpLights1(f io.Reader) (lights [1000][1000]bool) {
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		row := sc.Text()
		var command string
		var from, to Point
		row = strings.TrimPrefix(row, "turn ")
		fmt.Sscanf(row, "%s %d,%d through %d,%d", &command, &from.x, &from.y, &to.x, &to.y)
		switch command {
		case "on":
			for x := from.x; x <= to.x; x++ {
				for y := from.y; y <= to.y; y++ {
					lights[x][y] = true
				}
			}
		case "off":
			for x := from.x; x <= to.x; x++ {
				for y := from.y; y <= to.y; y++ {
					lights[x][y] = false
				}
			}
		case "toggle":
			for x := from.x; x <= to.x; x++ {
				for y := from.y; y <= to.y; y++ {
					lights[x][y] = !lights[x][y]
				}
			}
		default:
			log.Fatalf("failed to parse command: %q\n", command)
		}
	}

	return lights
}

func setUpLights2(f io.Reader) (lights [1000][1000]int) {
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		row := sc.Text()
		var command string
		var from, to Point
		row = strings.TrimPrefix(row, "turn ")
		fmt.Sscanf(row, "%s %d,%d through %d,%d", &command, &from.x, &from.y, &to.x, &to.y)
		switch command {
		case "on":
			for x := from.x; x <= to.x; x++ {
				for y := from.y; y <= to.y; y++ {
					lights[x][y]++
				}
			}
		case "off":
			for x := from.x; x <= to.x; x++ {
				for y := from.y; y <= to.y; y++ {
					lights[x][y] = mathx.Max(0, lights[x][y]-1)
				}
			}
		case "toggle":
			for x := from.x; x <= to.x; x++ {
				for y := from.y; y <= to.y; y++ {
					lights[x][y] += 2
				}
			}
		default:
			log.Fatalf("failed to parse command: %q\n", command)
		}
	}

	return lights
}
