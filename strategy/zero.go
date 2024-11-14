package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// Zero is a strategy and metastrategy which returns no values.
type Zero struct{}

func (Zero) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {}
}

func (Zero) MetaStrategy() iter.Seq[directsearch.Strategy] {
	return func(yield func(directsearch.Strategy) bool) {}
}
