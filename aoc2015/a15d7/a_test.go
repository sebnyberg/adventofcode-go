package a15d7_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

func Test_day(t *testing.T) {
	// Test input
	c := parseCircuit(must.Open("testinput"))

	require.EqualValues(t, 72, c.Value("d"))
	require.EqualValues(t, 507, c.Value("e"))
	require.EqualValues(t, 492, c.Value("f"))
	require.EqualValues(t, 114, c.Value("g"))
	require.EqualValues(t, 65412, c.Value("h"))
	require.EqualValues(t, 65079, c.Value("i"))
	require.EqualValues(t, 123, c.Value("x"))
	require.EqualValues(t, 456, c.Value("y"))

	// Real input
	c = parseCircuit(must.Open("input"))
	res := c.Value("a")
	require.EqualValues(t, 16076, c.Value("a"))

	// Reset memoization, and override b's value to the previous result
	c.values = make(map[string]int16)
	c.values["b"] = res
	require.EqualValues(t, 2797, c.Value("a"))
}

const (
	OpInit   = "INIT"
	OpNot    = "NOT"
	OpAnd    = "AND"
	OpOr     = "OR"
	OpRShift = "RSHIFT"
	OpLShift = "LSHIFT"
)

type Gate struct {
	Val int16
	Op  string
	In1 string
	In2 string
}

type Circuit struct {
	gates  map[string]Gate
	values map[string]int16
}

func (c *Circuit) Value(wire string) int16 {
	// Memoize values
	if val, exists := c.values[wire]; exists {
		return val
	}

	// Since all gates have been naively added even if the
	// incoming "wire" is actually a value, each root gate
	// will end up here requesting a "number" wire
	gate, exists := c.gates[wire]
	if !exists {
		// If the requested "wire" is actually a number,
		// return the number instead
		if n, err := strconv.Atoi(wire); err == nil {
			c.values[wire] = int16(n)
			return int16(n)
		}
		log.Fatalln("wire should exist but didnt: ", wire)
	}

	var val int16
	switch gate.Op {
	case OpInit:
		val = c.Value(gate.In1)
	case OpAnd:
		val = c.Value(gate.In1) & c.Value(gate.In2)
	case OpOr:
		val = c.Value(gate.In1) | c.Value(gate.In2)
	case OpNot:
		val = ^c.Value(gate.In1)
	case OpLShift:
		val = c.Value(gate.In1) << gate.Val
	case OpRShift:
		val = c.Value(gate.In1) >> gate.Val
	default:
		log.Fatalln("invalid gate", gate)
		return 0
	}
	c.values[wire] = val // Memoize result
	return val
}

func parseCircuit(f io.Reader) *Circuit {
	sc := bufio.NewScanner(f)
	c := &Circuit{
		gates:  make(map[string]Gate),
		values: make(map[string]int16),
	}
	for sc.Scan() {
		s := sc.Text()
		var wire string
		var gate Gate
		switch {
		case strings.Contains(s, OpAnd):
			gate.Op = OpAnd
			fmt.Sscanf(s, "%s AND %s -> %s", &gate.In1, &gate.In2, &wire)
		case strings.Contains(s, OpLShift):
			gate.Op = OpLShift
			fmt.Sscanf(s, "%s LSHIFT %d -> %s", &gate.In1, &gate.Val, &wire)
		case strings.Contains(s, OpRShift):
			gate.Op = OpRShift
			fmt.Sscanf(s, "%s RSHIFT %d -> %s", &gate.In1, &gate.Val, &wire)
		case strings.Contains(s, OpNot):
			gate.Op = OpNot
			fmt.Sscanf(s, "NOT %s -> %s", &gate.In1, &wire)
		case strings.Contains(s, OpOr):
			gate.Op = OpOr
			fmt.Sscanf(s, "%s OR %s -> %s", &gate.In1, &gate.In2, &wire)
		default: // Init
			gate.Op = OpInit
			fmt.Sscanf(s, "%s -> %s", &gate.In1, &wire)
		}
		c.gates[wire] = gate
	}
	return c
}
