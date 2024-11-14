package strategy

import (
	"github.com/wenooij/directsearch"
)

// Monotone represents a strategy which always returns the specified Action.
func Monotone(a directsearch.Action) directsearch.Strategy {
	return infiniteStrategy(func() directsearch.Action { return a })
}
