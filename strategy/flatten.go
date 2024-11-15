package strategy

import (
	"iter"
	"slices"

	"github.com/wenooij/directsearch"
)

const flatMax = 4096

type flat struct {
	max int
	iter.Seq[directsearch.Strategy]
}

func (s flat) Strategy() iter.Seq[directsearch.Action] {
	limit := s.max
	if limit == 0 {
		limit = flatMax
	}
	return func(yield func(directsearch.Action) bool) {
		n := 0
		for e := range s.Seq {
			var hasValues bool
			for a := range e.Strategy() {
				if hasValues = true; !yield(a) {
					break
				}
			}
			if hasValues {
				n = 0
			} else if n++; limit < n {
				break
			}
		}
	}
}

// FlattenStrategies chains the strategies into a sinle Strategy by flattening its returned Strategies to concrete Actions.
//
// Flatten will stop after a excessive number of strategies with no actions.
func FlattenStrategies(s ...directsearch.Strategy) directsearch.Strategy {
	return FlattenStrategiesIter(slices.Values(s))
}

// FlattenStrategies chains the strategies into a sinle Strategy by flattening its returned Strategies to concrete Actions.
//
// Flatten will stop after a excessive number of strategies with no actions.
func FlattenStrategiesIter(s iter.Seq[directsearch.Strategy]) directsearch.Strategy {
	next, _ := iter.Pull(s)
	return Flatten(func() (directsearch.Strategy, bool) { return next() })
}

// Flatten converts a strategy iterator into a Strategy by flattening its returned Strategies to concrete Actions.
//
// Flatten will stop after a excessive number of strategies with no actions.
func Flatten(f func() (directsearch.Strategy, bool)) directsearch.Strategy {
	return flat{flatMax, func(yield func(directsearch.Strategy) bool) {
		for e, ok := f(); ok && yield(e); e, ok = f() {
		}
	}}
}

// FlattenForever converts a strategy iterator into a Strategy by flattening its returned Strategies to concrete Actions.
func FlattenForever(f func() directsearch.Strategy) directsearch.Strategy {
	return flat{flatMax, func(yield func(directsearch.Strategy) bool) {
		for yield(f()) {
		}
	}}
}
