package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// While returns a strategy which returns from s as long as cond returns true.
func While(cond func() bool, s directsearch.Strategy) directsearch.Strategy {
	next, _ := iter.Pull(s.Strategy())
	return Func(func() (directsearch.Action, bool) {
		if !cond() {
			return nil, false
		}
		return next()
	})
}

// Case represents a strategy for a single case of a swich conditional.
type Case struct {
	Cond        func() bool
	Fallthrough bool
	directsearch.Strategy
}

// SwitchCase returns a flexible strategy where the first active Case strategy is returned via linear probing
// until the first case without Fallthrough is exhausted or all cases are exhausted or fail to match.
func SwitchCase(cases ...Case) directsearch.Strategy {
	next := make([]func() (directsearch.Action, bool), len(cases))
	for i, c := range cases {
		next[i], _ = iter.Pull(c.Strategy.Strategy())
	}
	return Func(func() (directsearch.Action, bool) {
		var fallingThrough bool
		for i, c := range cases {
			if fallingThrough || !c.Cond() {
				continue
			}
			fallingThrough = false
			a, ok := next[i]()
			if !ok && c.Fallthrough {
				fallingThrough = true
				continue // Fallthrouh to next case.
			}
			return a, ok
		}
		return nil, false // Default case
	})
}
