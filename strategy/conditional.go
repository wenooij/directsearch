package strategy

import (
	"github.com/wenooij/directsearch"
)

// Conditional strategy represents a flexible strategy where ActiveStrategy is used when Cond returns true otherwise FallbackStrategy is used.
type Conditional struct {
	Cond             func() bool
	ActiveStrategy   directsearch.Strategy
	FallbackStrategy directsearch.Strategy
}

// Next returns the next Action depending on the results of Cond.
func (c Conditional) Next() directsearch.Action {
	if c.Cond() {
		return c.ActiveStrategy.Next()
	}
	return c.FallbackStrategy.Next()
}

// Case represents a strategy for a single case of a swich conditional.
type Case struct {
	Cond func() bool
	directsearch.Strategy
}

// SwitchCase is a flexible strategy where the first active Case is returned otherwise FallbackStrategy is used.
type SwitchCase struct {
	Cases            []Case
	FallbackStrategy directsearch.Strategy
}

// Next returns the next Action depending on the result of the Case.
func (c SwitchCase) Next() directsearch.Action {
	for _, c := range c.Cases {
		if c.Cond() {
			return c.Next()
		}
	}
	return c.FallbackStrategy.Next()
}
