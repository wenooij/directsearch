package strategy

import "github.com/wenooij/directsearch"

type RoundRobin struct {
	Strategies []directsearch.Strategy
	NextIndex  int
}

func (r *RoundRobin) Next() directsearch.Action {
	s := r.Strategies[r.NextIndex]
	if r.NextIndex++; r.NextIndex == len(r.Strategies) {
		r.NextIndex = 0
	}
	return s.Next()
}
