package strategy

import "github.com/wenooij/directsearch"

type RoundRobin struct {
	Actions   []directsearch.Action
	NextIndex int
}

func (r *RoundRobin) Next() directsearch.Action {
	a := r.Actions[r.NextIndex]
	if r.NextIndex++; r.NextIndex == len(r.Actions) {
		r.NextIndex = 0
	}
	return a
}
