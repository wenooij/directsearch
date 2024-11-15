package strategy

import (
	"iter"
	"slices"

	"github.com/wenooij/directsearch"
)

// Cycle returns a strategy which cycles over s.
//
// This is often called "round robin" ordering.
func Cycle(s ...directsearch.Strategy) directsearch.Strategy {
	n := len(s)
	if n == 0 {
		return Zero{}
	}
	next, _ := iter.Pull(slices.Values(s))
	return FlattenForever(func() directsearch.Strategy {
		e, ok := next()
		if !ok {
			next, _ = iter.Pull(slices.Values(s))
			e, _ = next()
		}
		return e
	})
}
