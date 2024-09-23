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

type Case struct {
	Cond func() bool
	directsearch.Strategy
}

type SwitchCase struct {
	Cases            []Case
	FallbackStrategy directsearch.Strategy
}

func (c SwitchCase) Next() directsearch.Action {
	for _, c := range c.Cases {
		if c.Cond() {
			return c.Next()
		}
	}
	return c.FallbackStrategy.Next()
}
