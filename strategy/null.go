package strategy

import "github.com/wenooij/directsearch"

// Null is a strategy which always returns the nil Action.
type Null struct{}

// Next returns the nil Action.
func (Null) Next() directsearch.Action { return nil }

// IgnoreNulls wraps a strategy and attempts to ignore successive nil Actions up to the given Limit. Otherwise it returns the FallbackStrategy.
type IgnoreNulls struct {
	directsearch.Strategy
	Limit            int
	FallbackStrategy directsearch.Strategy
}

// Init initializes the Limit to the default value.
func (s *IgnoreNulls) Init() {
	if s.Limit == 0 {
		s.Limit = 100
	}
}

// Next returns the next non-nil Action or returns an Action from the FallbackStrategy.
func (s *IgnoreNulls) Next() directsearch.Action {
	for attempts := 0; attempts < s.Limit; attempts++ {
		if a := s.Strategy.Next(); a != nil {
			return a
		}
	}
	return s.FallbackStrategy.Next()
}
