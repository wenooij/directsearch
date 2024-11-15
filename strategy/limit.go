package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// Limit returns a strategy limiting the number of Actions before stopping.
func Limit(n int, s directsearch.Strategy) directsearch.Strategy {
	next, stop := iter.Pull(s.Strategy())
	return Func(func() (directsearch.Action, bool) {
		if n <= 0 {
			stop()
			return nil, false
		}
		n--
		return next()
	})
}

// First returns a strategy which picks at most one Action from s.
//
// It is equivalent to calling Limit(1, s).
func First(s directsearch.Strategy) directsearch.Strategy { return Limit(1, s) }
