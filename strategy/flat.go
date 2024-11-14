package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

type flat struct{ directsearch.MetaStrategy }

func (s flat) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for e := range s.MetaStrategy.MetaStrategy() {
			for a := range e.Strategy() {
				if !yield(a) {
					break
				}
			}
		}
	}
}

// Flatten converts a MetaStrategy into a Strategy by flattening its returned Strategies to concrete Actions.
//
// You can save implementing Strategy yourself by simply wrapping it with Flat instead.
func Flatten(s directsearch.MetaStrategy) directsearch.Strategy { return flat{s} }
