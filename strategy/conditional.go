package strategy

import (
	"github.com/wenooij/directsearch"
)

// Conditional returns a metastrategy which returns active when Cond returns true otherwise fallback.
func Conditional(cond func() bool, active, fallback directsearch.Strategy) directsearch.MetaStrategy {
	return infiniteMetaStrategy(func() directsearch.Strategy {
		if cond() {
			return active
		}
		return fallback
	})
}

// Case represents a strategy for a single case of a swich conditional.
type Case struct {
	Cond func() bool
	directsearch.Strategy
}

// SwitchCase returns a flexible strategy where the first active Case is returned via linear probing
// otherwise fallback if none are active.
func SwitchCase(fallback directsearch.Strategy, cases ...Case) directsearch.MetaStrategy {
	return infiniteMetaStrategy(func() directsearch.Strategy {
		for _, c := range cases {
			if c.Cond() {
				return c.Strategy
			}
		}
		return fallback
	})
}
