package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

type Random struct {
	Strategies []directsearch.Strategy
	Rand       *rand.Rand
}

func (r *Random) Next() directsearch.Action {
	idx := r.Rand.IntN(len(r.Strategies))
	s := r.Strategies[idx]
	return s.Next()
}
