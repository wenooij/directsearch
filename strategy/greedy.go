package strategy

import "github.com/wenooij/directsearch"

// Greedy maintains a list of strategies and always uses the first.
//
// If there are no available strategies Greedy stops.
// The elements of s can change while iterating but the length is fixed.
func Greedy(s ...directsearch.Strategy) directsearch.MetaStrategy {
	if len(s) == 0 {
		return Zero{}
	}
	return infiniteMetaStrategy(func() directsearch.Strategy { return s[0] })
}
