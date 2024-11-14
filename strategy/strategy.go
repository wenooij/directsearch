package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

type strategy func() (directsearch.Action, bool)

func (s strategy) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for a, ok := s(); ok && yield(a); a, ok = s() {
		}
	}
}

type infiniteStrategy func() directsearch.Action

func (s infiniteStrategy) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for yield(s()) {
		}
	}
}

type metaStrategy func() (directsearch.Strategy, bool)

func (s metaStrategy) MetaStrategy() iter.Seq[directsearch.Strategy] {
	return func(yield func(directsearch.Strategy) bool) {
		for e, ok := s(); ok && yield(e); e, ok = s() {
		}
	}
}

type infiniteMetaStrategy func() directsearch.Strategy

func (s infiniteMetaStrategy) MetaStrategy() iter.Seq[directsearch.Strategy] {
	return func(yield func(directsearch.Strategy) bool) {
		for yield(s()) {
		}
	}
}
