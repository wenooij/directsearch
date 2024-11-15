package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// Coalesce maintains a list of strategies and always uses the first non-exhausted one.
//
// The elements of s can change while iterating but the length is fixed.
func Coalesce(s ...directsearch.Strategy) directsearch.Strategy {
	n := len(s)
	if n == 0 {
		return Zero{}
	}
	return Func(func() (directsearch.Action, bool) {
		for _, e := range s {
			next, _ := iter.Pull(e.Strategy())
			v, ok := next()
			if ok {
				return v, true
			}
		}
		return nil, false
	})
}
