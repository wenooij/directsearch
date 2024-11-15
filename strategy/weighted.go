package strategy

import (
	"iter"
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

// Weight is an interface for weighted items with positive weights.
type Weight interface{ Weight() float64 }

// WeightedStrategy is a pair combining a item with an associated strategy.
type WeightedStrategy[E Weight] struct {
	e E
	directsearch.Strategy
}

// MakeWeighted creates a new WeightedStrategy entry.
func MakeWeighted[E Weight](s directsearch.Strategy, e E) WeightedStrategy[E] {
	return WeightedStrategy[E]{Strategy: s, e: e}
}

// Weighted provides a metastrategy for selecting strategies based on associated weights.
//
// Use [Replace] to change Weights for ellements.
type Weighted[E Weight] struct {
	rand        *rand.Rand
	totalWeight float64
	entries     []WeightedStrategy[E]
}

// Append the given weight elements to the weighted strategy.
func (w *Weighted[E]) Append(es ...WeightedStrategy[E]) {
	w.entries = append(w.entries, es...)
	for _, e := range es {
		w.totalWeight += e.e.Weight()
	}
}

// Replace should be called to efficiently change the WeightedStrategy at i.
//
// Replace panics if i is out of bounds.
func (w *Weighted[E]) Replace(i int, e WeightedStrategy[E]) {
	old := w.entries[i]
	w.totalWeight += e.e.Weight() - old.e.Weight()
	w.entries[i] = e
}

// ReplaceStrategy should be called to efficiently change the Strategy at i.
//
// ReplaceStrategy panics if i is out of bounds.
func (w Weighted[E]) ReplaceStrategy(i int, s directsearch.Strategy) { w.entries[i].Strategy = s }

// ReplaceWeight should be called to efficiently change the WeightedStrategy at i.
//
// ReplaceWeight panics if i is out of bounds.
func (w *Weighted[E]) ReplaceItem(i int, e E) {
	old := w.entries[i]
	w.totalWeight += e.Weight() - old.e.Weight()
	w.entries[i].e = e
}

func (w Weighted[E]) Strategy() iter.Seq[directsearch.Action] {
	return FlattenForever(w.next).Strategy()
}

func (w Weighted[E]) next() directsearch.Strategy {
	t := w.rand.Float64() * w.totalWeight
	for _, e := range w.entries {
		if t -= e.e.Weight(); t <= 0 {
			return e.Strategy
		}
	}
	return Zero{} // Should normally be impossible.
}
