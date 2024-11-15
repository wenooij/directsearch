package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

type flat struct {
	iter.Seq[directsearch.Strategy]
}

func (s flat) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for e := range s.Seq {
			for a := range e.Strategy() {
				if !yield(a) {
					break
				}
			}
		}
	}
}

// FlattenFlatten converts a strategy iterator into a Strategy by flattening its returned Strategies to concrete Actions.
func Flatten(f func() (directsearch.Strategy, bool)) directsearch.Strategy {
	return flat{func(yield func(directsearch.Strategy) bool) {
		for e, ok := f(); ok && yield(e); e, ok = f() {
		}
	}}
}

// FlattenForever converts a strategy iterator into a Strategy by flattening its returned Strategies to concrete Actions.
func FlattenForever(f func() directsearch.Strategy) directsearch.Strategy {
	return flat{func(yield func(directsearch.Strategy) bool) {
		for yield(f()) {
		}
	}}
}
