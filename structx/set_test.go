package structx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IntSet(t *testing.T) {
	s := NewIntSet(1, 2, 3)
	require.ElementsMatch(t, []int{1, 2, 3}, s.GetAll())
	s.Remove(1)
	require.ElementsMatch(t, []int{2, 3}, s.GetAll())
}
