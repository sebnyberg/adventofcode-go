package a15d1

import (
	"testing"
	"text/scanner"

	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

func Test_day1(t *testing.T) {
	floor, basementPos := day1()
	require.Equal(t, 138, floor)
	require.Equal(t, 1771, basementPos)
}

func day1() (floor, basementPos int) {
	f := must.Open("input")
	var s scanner.Scanner
	s.Init(f)
	basementPos = -1
	i := 0
	for {
		i++
		switch s.Scan() {
		case '(':
			floor++
		case ')':
			floor--
		case scanner.EOF:
			return floor, basementPos
		}
		if floor == -1 && basementPos == -1 {
			basementPos = i
		}
	}
}
