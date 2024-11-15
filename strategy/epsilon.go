package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

// Epsilon selects eps proportional to the rate epsilon.
// Otherwise it returns the fallback strategy.
func Epsilon(r *rand.Rand, epsilon float64, eps, fallback directsearch.Strategy) directsearch.Strategy {
	return FlattenForever(func() directsearch.Strategy {
		if r.Float64() < epsilon {
			return eps
		}
		return fallback
	})
}
