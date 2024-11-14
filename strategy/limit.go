package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// Limit returns a strategy limiting the number of Actions before stopping.
func Limit(n int, s directsearch.Strategy) directsearch.Strategy {
	i := 0
	next, stop := iter.Pull(s.Strategy())
	return strategy(func() (directsearch.Action, bool) {
		if i++; n < i {
			stop()
			return nil, false
		}
		return next()
	})
}
