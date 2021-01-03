package must

import "strconv"

func ParseHex(s string) int {
	i, err := strconv.ParseInt(s, 16, 64)
	Check(err)
	return int(i)
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}
