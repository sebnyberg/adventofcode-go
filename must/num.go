package must

import "strconv"

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}
