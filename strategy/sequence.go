package strategy

import "github.com/wenooij/directsearch"

type Sequence struct {
	Strategies []directsearch.Strategy
	NextIndex  int
}

func (s Sequence) Next() directsearch.Action {
	for ; s.NextIndex < len(s.Strategies); s.NextIndex++ {
		if a := s.Strategies[s.NextIndex].Next(); a != nil {
			return a
		}
	}
	return nil
}
