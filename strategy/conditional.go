package strategy

import (
	"github.com/wenooij/directsearch"
)

type Conditional struct {
	Condition        func() bool
	Strategy         directsearch.Strategy
	FallbackStrategy directsearch.Strategy
}

func (c Conditional) Next() directsearch.Action {
	if c.Condition() {
		return c.Strategy.Next()
	}
	return c.FallbackStrategy.Next()
}
