package strategy

import (
	"github.com/wenooij/directsearch"
)

// Cycle returns a metastrategy which cycles over s.
//
// This is often called "round robin" ordering.
func Cycle(s ...directsearch.Strategy) directsearch.MetaStrategy {
	n := len(s)
	if n == 0 {
		return Zero{}
	}
	i := 0
	return infiniteMetaStrategy(func() directsearch.Strategy {
		if n <= i {
			i = 0
		}
		e := s[i]
		i++
		return e
	})
}
