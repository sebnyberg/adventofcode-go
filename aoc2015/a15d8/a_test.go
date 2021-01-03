package a15d8

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/sebnyberg/adventofcode/must"
	"github.com/stretchr/testify/require"
)

// Start: 3/1/2020 12:06
// End part 1: 12:47
// End part 2: 12:57

func Test_day(t *testing.T) {
	orig, decoded, encoded := readStrings(must.Open("testinput"))
	var norig, ndecoded, nencoded int
	for i := range orig {
		norig += utf8.RuneCountInString(orig[i])
		ndecoded += utf8.RuneCountInString(decoded[i])
		nencoded += utf8.RuneCountInString(encoded[i])
	}
	require.Equal(t, 23, norig)
	require.Equal(t, 11, ndecoded)
	require.Equal(t, 19, nencoded-norig)

	orig, decoded, encoded = readStrings(must.Open("input"))
	norig, ndecoded, nencoded = 0, 0, 0
	for i := range orig {
		norig += utf8.RuneCountInString(orig[i])
		ndecoded += utf8.RuneCountInString(decoded[i])
		nencoded += utf8.RuneCountInString(encoded[i])
	}
	require.Equal(t, 1342, norig-ndecoded)
	require.Equal(t, 2074, nencoded-norig)
}

func readStrings(r io.Reader) (orig, inmem, encoded []string) {
	sc := bufio.NewScanner(r)
	orig = make([]string, 0, 100)
	inmem = make([]string, 0, 100)
	encoded = make([]string, 0, 100)
	for sc.Scan() {
		row := sc.Text()
		orig = append(orig, row)
		inmem = append(inmem, decode(row))
		encoded = append(encoded, encode(row))
	}
	return orig, inmem, encoded
}

func Test_decode(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want string
	}{
		{`""`, ""},
		{`"\"`, `"`},
		{`"\xfb\""`, `û"`},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, decode(tc.in))
		})
	}
}

func decode(code string) string {
	runes := []rune(code)
	var decoded strings.Builder
	for i := 1; i < len(runes)-1; i++ {
		if runes[i] == '\\' {
			i++
			if runes[i] == 'x' {
				i++
				decoded.WriteRune(rune(must.ParseHex(string(runes[i : i+2]))))
				i++
				continue
			}
		}
		decoded.WriteRune(runes[i])
	}
	return decoded.String()
}

func Test_encode(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want string
	}{
		{`""`, `"\"\""`},
		{`"abc"`, `"\"abc\""`},
		{`"aaa\"aaa"`, `"\"aaa\\\"aaa\""`},
		{`"\x27"`, `"\"\\x27\""`},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, encode(tc.in))
		})
	}
}

func encode(s string) string {
	var sb strings.Builder
	sb.WriteRune('"')

	for _, ch := range s {
		switch ch {
		case '"', '\\':
			sb.WriteRune('\\')
		}
		sb.WriteRune(ch)
	}

	sb.WriteRune('"')
	return sb.String()
}
