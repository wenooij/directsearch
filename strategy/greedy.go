package strategy

import (
	"github.com/wenooij/directsearch"
)

type Greedy struct {
	Actions []directsearch.Action
}

func (g Greedy) Next() directsearch.Action {
	return g.Actions[0]
}
