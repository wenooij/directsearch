package strategy

import (
	"github.com/wenooij/directsearch"
)

// Greedy maintains a list of strategies and always uses the first.
//
// If there are no available strategies Greedy returns nil.
type Greedy struct{ Strategies []directsearch.Strategy }

// Next returns the next Action from the first strategy or nil.
func (g Greedy) Next() directsearch.Action {
	if len(g.Strategies) == 0 {
		return nil
	}
	return g.Strategies[0].Next()
}
