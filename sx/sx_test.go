package sx

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Permutator(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}
	perms := make(map[string]struct{})
	p := NewPermutator(len(a))
	i := 0
	for p.Permutate(func(i, j int) { a[i], a[j] = a[j], a[i] }) {
		var sb strings.Builder
		for _, n := range a {
			sb.WriteString(strconv.Itoa(n))
		}
		s := sb.String()
		if _, exists := perms[s]; exists {
			t.Logf("tried to add %v at index %v and failed\n", s, i)
			t.Logf("+%v\n", perms)
			t.FailNow()
		}
		t.Logf("i: %v\ts: %v\n", i, s)
		i++
		perms[s] = struct{}{}
	}
	require.Equal(t, i, 120)
}

func Test_Permutator_Strings(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	perms := make(map[string]struct{})
	p := NewPermutator(len(a))
	i := 0
	for p.Permutate(func(i, j int) { a[i], a[j] = a[j], a[i] }) {
		s := strings.Join(a, "")
		if _, exists := perms[s]; exists {
			t.Logf("tried to add %v at index %v and failed\n", s, i)
			t.Logf("+%v\n", perms)
			t.FailNow()
		}
		t.Logf("i: %v\ts: %v\n", i, s)
		i++
		perms[s] = struct{}{}
	}
	require.Equal(t, i, 24)
}

func Test_Permutation(t *testing.T) {
	a := []int{4, 3, 2, 1}
	perms := make(map[string]struct{})
	for i := 0; i < CountPerms(a); i++ {
		Permute(i, a)
		var sb strings.Builder
		for _, n := range a {
			sb.WriteString(strconv.Itoa(n))
		}
		s := sb.String()
		if _, exists := perms[s]; exists {
			t.Logf("tried to add %v at index %v and failed\n", s, i)
			t.Logf("+%v\n", perms)
			t.FailNow()
		}
		perms[s] = struct{}{}
	}
}

func Test_PermutationSwap(t *testing.T) {
	results := make([][]int, 0)
	for i := 0; i < 19; i++ {
		start, end := PermutationSwap(i, 4)
		results = append(results, []int{start, end})
	}
	require.Equal(t,
		[][]int{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 1},
			{0, 2},
			{0, 1},
			{0, 3},
			{0, 1},
			{0, 2},
			{0, 1},
			{0, 2},
			{0, 1},
			{1, 3},
			{0, 1},
			{0, 2},
			{0, 1},
			{0, 2},
			{0, 1},
			{2, 3},
		},
		results,
	)
}
