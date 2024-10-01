package strategy

import (
	"math/rand/v2"
	"sort"

	"github.com/wenooij/directsearch"
)

// WeightedItem defines an interface for items which have an associated weight.
type WeightedItem interface {
	Weight() float64
	SetWeight(float64)
}

// WeightedEntry is a pair combining a item with an associated strategy.
type WeightedEntry[E WeightedItem] struct {
	Item E
	directsearch.Strategy
}

type wrappedWeightItem[E WeightedItem] struct {
	w    *Weighted[E]
	Item E
}

func (w wrappedWeightItem[E]) Weight() float64 { return w.Item.Weight() }

func (w wrappedWeightItem[E]) SetWeight(weight float64) {
	w.w.TotalWeight += weight - w.Item.Weight()
	w.Item.SetWeight(weight)
}

// Weighted provides a strategy for selecting strategies based on associated weights.
type Weighted[E WeightedItem] struct {
	Rand        *rand.Rand
	TotalWeight float64
	Entries     []WeightedEntry[wrappedWeightItem[E]]
}

func (w *Weighted[E]) wrapWeightItem(e E) wrappedWeightItem[E] { return wrappedWeightItem[E]{w, e} }

// Sort the strategies by descending weight.
func (w Weighted[E]) Sort() { sort.Sort(byWeight[E](w.Entries)) }

// ChangeWeight should be called when changing the weight to keep the values.
func (w *Weighted[E]) ChangeWeight(i int, weight float64) {
	e := w.Entries[i]
	w.TotalWeight += weight - e.Item.Weight()
	e.Item.SetWeight(weight)
}

// Select a random strategy proportional to the weights
func (w Weighted[E]) Select() directsearch.Strategy {
	t := w.Rand.Float64() * w.TotalWeight
	for _, e := range w.Entries {
		if t -= e.Item.Weight(); t <= 0 {
			return e.Strategy
		}
	}
	return Null{} // Not normally possible.
}

func (w Weighted[E]) Next() directsearch.Action { s := w.Select(); return s.Next() }

type byWeight[E WeightedItem] []WeightedEntry[wrappedWeightItem[E]]

func (a byWeight[E]) Len() int           { return len(a) }
func (a byWeight[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWeight[E]) Less(i, j int) bool { return a[i].Item.Weight() > a[j].Item.Weight() }
