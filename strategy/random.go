package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

// Random selects a uniform random strategy from the list.
type Random struct {
	Strategies []directsearch.Strategy
	Rand       *rand.Rand
}

// Select selects the next random action.
func (r *Random) Select() directsearch.Strategy {
	idx := r.Rand.IntN(len(r.Strategies))
	return r.Strategies[idx]
}

// Next selects the next Action from the selected strategy.
func (r *Random) Next() directsearch.Action { s := r.Select(); return s.Next() }
