package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

type EpsilonGreedy struct {
	Rand            *rand.Rand
	Epsilon         float64
	Strategy        directsearch.Strategy
	EpsilonStrategy directsearch.Strategy
}

func (e *EpsilonGreedy) Next() directsearch.Action {
	if e.Rand.Float64() < e.Epsilon {
		return e.EpsilonStrategy.Next()
	}
	return e.Strategy.Next()
}
