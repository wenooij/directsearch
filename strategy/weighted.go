package strategy

import (
	"math/rand/v2"
	"sort"

	"github.com/wenooij/directsearch"
)

type WeightedInterface interface {
	Weight() float64
}

type WeightedEntry[T WeightedInterface] struct {
	Value T
	directsearch.Strategy
}

type Weighted[T WeightedInterface] struct {
	Rand        *rand.Rand
	TotalWeight float64
	Entries     []WeightedEntry[T]
}

func (w Weighted[T]) Init() { sort.Sort(byWeight[T](w.Entries)) }

func (w Weighted[T]) Next() directsearch.Action {
	t := w.Rand.Float64() * w.TotalWeight
	for _, e := range w.Entries {
		if t -= e.Value.Weight(); t <= 0 {
			return e.Next()
		}
	}
	return nil // Not usually possible.
}

type byWeight[T WeightedInterface] []WeightedEntry[T]

func (a byWeight[T]) Len() int           { return len(a) }
func (a byWeight[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWeight[T]) Less(i, j int) bool { return a[i].Value.Weight() > a[j].Value.Weight() }
