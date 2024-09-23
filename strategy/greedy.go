package strategy

import (
	"github.com/wenooij/directsearch"
)

type Greedy struct{ Strategies []directsearch.Strategy }

func (g Greedy) Next() directsearch.Action { return g.Strategies[0].Next() }
