package strategy

import "github.com/wenooij/directsearch"

// Limited represents a strategy where LimitedStrategy is used Limit times after which FallbackStrategy is used.
//
// Calling Reset resets the Limited strategy counter.
type Limited struct {
	Count            int
	Limit            int
	LimitedStrategy  directsearch.Strategy
	FallbackStrategy directsearch.Strategy
}

// Reset resets the counter to 0.
func (l *Limited) Reset() { l.Count = 0 }

// Next returns the next action depending on the current counter.
func (l *Limited) Next() directsearch.Action {
	if l.Count < l.Limit {
		l.Count++
		return l.LimitedStrategy.Next()
	}
	return l.FallbackStrategy.Next()
}
