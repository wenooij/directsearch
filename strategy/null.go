package strategy

import "github.com/wenooij/directsearch"

type Null struct{}

func (Null) Next() directsearch.Action { return nil }

type IgnoreNulls struct {
	directsearch.Strategy
	Limit            int
	FallbackStrategy directsearch.Strategy
}

func (s *IgnoreNulls) Init() {
	if s.Limit == 0 {
		s.Limit = 100
	}
}

func (s *IgnoreNulls) Next() directsearch.Action {
	for attempts := 0; attempts < s.Limit; attempts++ {
		if a := s.Strategy.Next(); a != nil {
			return a
		}
	}
	return s.FallbackStrategy.Next()
}
