package strategy

import (
	"iter"
	"math/rand/v2"
	"slices"

	"github.com/wenooij/directsearch"
)

// Any returns a strategy which uniformly samples s.
func Any(r *rand.Rand, s ...directsearch.Strategy) directsearch.Strategy {
	n := len(s)
	if n == 0 {
		return Zero{}
	}
	return FlattenForever(func() directsearch.Strategy {
		i := r.IntN(n)
		return s[i]
	})
}

// Permute returns a random permutation of s.
func Permute(r *rand.Rand, s ...directsearch.Strategy) directsearch.Strategy {
	n := len(s)
	cp := make([]int, n)
	for i := range cp {
		cp[i] = i
	}
	r.Shuffle(len(cp), func(i, j int) { cp[i], cp[j] = cp[j], cp[i] })
	next, _ := iter.Pull(slices.Values(cp))
	return Flatten(func() (directsearch.Strategy, bool) {
		i, ok := next()
		if !ok {
			return nil, false
		}
		e := s[i]
		return e, true
	})
}
