package a15d4

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_day(t *testing.T) {
	require.Equal(t, "282749", findHashWithPrefix("yzbqklnj", "00000"))
	require.Equal(t, "9962624", findHashWithPrefix("yzbqklnj", "000000"))
}

func findHashWithPrefix(key string, prefix string) string {
	var ans string
	for i := 1; ; i++ {
		ans = key + strconv.Itoa(i)
		hash := md5.Sum([]byte(ans))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), prefix) {
			break
		}
	}
	return strings.TrimPrefix(ans, key)
}
