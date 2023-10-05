package a15d9

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/sebnyberg/adventofcode/must"
	"github.com/sebnyberg/adventofcode/sx"
	"github.com/stretchr/testify/require"
)

// Start: 3/1/2020 14:28
// 17:11 realise that this is not a directed graph... rewrite

func Test_day(t *testing.T) {
	routes := scanRoutes(must.Open("input"))
	cities := make([]string, 0, len(routes))
	for city := range routes {
		cities = append(cities, city)
	}
	p := sx.NewPermutator(len(cities))
	minDist := math.MaxInt64
	for p.Permutate(func(i, j int) { cities[i], cities[j] = cities[j], cities[i] }) {
		dist := 0
		for i := 0; i < len(cities)-1; i++ {
			dist += routes[cities[i]][cities[i+1]]
		}
		if dist < minDist {
			minDist = dist
		}
	}
	require.Equal(t, 0, minDist)

	t.FailNow()
}

func scanRoutes(r io.ReadCloser) map[string]map[string]int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	routes := make(map[string]map[string]int)
	for sc.Scan() {
		row := sc.Text()
		var from, to string
		var distance int
		fmt.Sscanf(row, "%s to %s = %d", &from, &to, &distance)
		if _, exists := routes[from]; !exists {
			routes[from] = make(map[string]int)
		}
		if _, exists := routes[to]; !exists {
			routes[to] = make(map[string]int)
		}
		routes[from][to] = distance
		routes[to][from] = distance
	}

	return routes
}
