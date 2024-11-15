package strategy

import "github.com/wenooij/directsearch"

// Chain the strategies together until each is exhausted.
func Chain(s ...directsearch.Strategy) directsearch.Strategy {
	i := 0
	return Flatten(func() (directsearch.Strategy, bool) {
		if len(s) <= i {
			return nil, false
		}
		e := s[i]
		i += 1
		return e, true
	})
}
