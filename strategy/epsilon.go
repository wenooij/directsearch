package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

// Epsilon selects EpsilonStrategy proportional to the rate Epsilon
type Epsilon struct {
	Rand             *rand.Rand
	Epsilon          float64
	EpsilonStrategy  directsearch.Strategy
	FallbackStrategy directsearch.Strategy
}

func (e *Epsilon) Next() directsearch.Action {
	if e.Rand.Float64() < e.Epsilon {
		return e.EpsilonStrategy.Next()
	}
	return e.FallbackStrategy.Next()
}
