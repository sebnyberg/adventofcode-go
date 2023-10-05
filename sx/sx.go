package sx

type Permutator struct {
	permIndex int
	nperms    int
	// t is the number of swaps per number index
	t []int
	// c is the number of performed swaps per number index
	c []int
}

// Create a new Permutor which permutates an array of length n
func NewPermutator(n int) *Permutator {
	p := Permutator{
		t:      make([]int, n),
		c:      make([]int, n),
		nperms: 1,
	}
	for i := 0; i < n; i++ {
		p.t[i] = i + 1
		p.nperms *= i + 1
	}
	return &p
}

// Permutate permutates the provided array and returns false
// if no more permutations can be done.
func (p *Permutator) Permutate(swapFn func(i, j int)) bool {
	if p.permIndex == 0 {
		p.permIndex++
		return true
	}
	if p.permIndex >= p.nperms {
		return false
	}

	i := 0
	for p.c[i] >= p.t[i] {
		p.c[i] = 0
		i++
	}

	start, end := 0, i+1
	if i%2 == 0 {
		start = p.c[i]
	}
	p.c[i]++

	swapFn(start, end)
	p.permIndex++

	return true
}

func CountPerms(a []int) (perms int) {
	perms = 1
	for n := 2; n <= len(a); n++ {
		perms *= n
	}
	return perms
}

func Permute(i int, a []int) {
	n := len(a)
	if i == 0 || n == 1 {
		return
	}

	start, end := 0, 1
	fac := 1
	for size := 2; size < n; size++ {
		if i < fac {
			break
		}
		fac *= size
		if i%fac == 0 {
			end = size
			if size%2 == 1 {
				start = (i/fac)%(fac*size+1) - 1
			}
		}
	}

	a[start], a[end] = a[end], a[start]
}

func PermutationSwap(i int, n int) (start, end int) {
	if i == 0 {
		return 0, 0
	}
	start, end = 0, 1
	fac := 1
	for size := 2; size < n; size++ {
		if i < fac {
			break
		}
		fac *= size
		if i%fac == 0 {
			end = size
			if size%2 == 1 {
				start = (i/fac)%(fac*size+1) - 1
			}
		}
	}
	return start, end
}
