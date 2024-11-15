package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// Zero is a strategy which returns no values.
type Zero struct{}

func (Zero) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {}
}
