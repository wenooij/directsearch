package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

type strategyFunc func() (directsearch.Action, bool)

func (s strategyFunc) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for a, ok := s(); ok && yield(a); a, ok = s() {
		}
	}
}

// Func returns a strategy which yields from f until exhausted.
func Func(f func() (directsearch.Action, bool)) directsearch.Strategy { return strategyFunc(f) }

type monotoneFunc func() directsearch.Action

func (s monotoneFunc) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for yield(s()) {
		}
	}
}

// MonotoneFunc returns a stratey which yields from f.
func MonotoneFunc(f func() directsearch.Action) directsearch.Strategy { return monotoneFunc(f) }

// Monotone returns a strategy which yields the specified Action.
func Monotone(a directsearch.Action) directsearch.Strategy {
	return MonotoneFunc(func() directsearch.Action { return a })
}
