package a15d5

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

func Test_countNiceStrings(t *testing.T) {
	require.Equal(t, 255, countNiceStrings(must.Open("input"), isNicePart1))
	require.Equal(t, 55, countNiceStrings(must.Open("input"), isNicePart2))
}

func countNiceStrings(f io.Reader, niceFunc func(string) bool) int {
	sc := bufio.NewScanner(f)
	var nice int

	for sc.Scan() {
		row := sc.Text()
		if niceFunc(row) {
			nice++
		}
	}
	return nice
}

func isNicePart1(s string) bool {
	// Check naughty strings
	naughtyStrings := []string{"ab", "cd", "pq", "xy"}
	for _, naughtyString := range naughtyStrings {
		if strings.Contains(s, naughtyString) {
			return false
		}
	}

	vowels := "aeiou"

	// Simultaneously check vowels / two letters in a row
	nvowels := 0
	var prev rune // Compare previous rune with current to check recurrence
	var twoInARow bool
	for i, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			nvowels++
		}
		if i > 0 && prev == ch {
			twoInARow = true
		}
		if nvowels >= 3 && twoInARow {
			return true
		}
		prev = ch
	}

	return false
}

func Test_isNicePart2(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want bool
	}{
		{"xyxbbb", false},
		{"xyxbbbb", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, isNicePart2(tc.in))
		})
	}
}

func isNicePart2(s string) bool {
	// Find a pair that appears twice in the string
	for i := 0; i < len(s)-2; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			goto Continue
		}
	}
	return false

Continue:
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] {
			return true
		}
	}

	return false
}
