package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

// Random returns a metastrategy which uniformly samples s.
func Random(r *rand.Rand, s ...directsearch.Strategy) directsearch.MetaStrategy {
	n := len(s)
	if n == 0 {
		return Zero{}
	}
	return infiniteMetaStrategy(func() directsearch.Strategy {
		i := r.IntN(n)
		return s[i]
	})
}
