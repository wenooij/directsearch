package strategy

import (
	"container/heap"
	"iter"

	"github.com/wenooij/directsearch"
)

// Prioritized is an interface for items with a real value priority.
//
// Smaller priority values are higher in priority.
type Priority interface {
	Priority() float64
}

// PrioritizedStrategy wraps a Priority with an associated strategy.
type PrioritizedStrategy[E Priority] struct {
	e E
	directsearch.Strategy
}

func (p PrioritizedStrategy[E]) Item() E { return p.e }

func MakePrioritized[E Priority](s directsearch.Strategy, e E) PrioritizedStrategy[E] {
	return PrioritizedStrategy[E]{Strategy: s, e: e}
}

// Prioritized defines a strategy which selects the highest priority entry greedily.
//
// Smaller priority values are higher in priority.
type Prioritized[E Priority] struct{ entries []PrioritizedStrategy[E] }

func (p Prioritized[E]) MetaStrategy() iter.Seq[directsearch.Strategy] {
	return infiniteMetaStrategy(p.next).MetaStrategy()
}

func (p Prioritized[E]) next() directsearch.Strategy { return p.entries[0].Strategy }

// Init the priority list.
func (p Prioritized[E]) Init() { heap.Init((*byPriority[E])(&p.entries)) }
func (p *Prioritized[E]) Pop() PrioritizedStrategy[E] {
	return heap.Pop((*byPriority[E])(&p.entries)).(PrioritizedStrategy[E])
}
func (p Prioritized[E]) Push(x PrioritizedStrategy[E]) { heap.Push((*byPriority[E])(&p.entries), x) }
func (p Prioritized[E]) Remove(i int) PrioritizedStrategy[E] {
	return heap.Remove((*byPriority[E])(&p.entries), i).(PrioritizedStrategy[E])
}

// Replace the Entry at i.
func (p Prioritized[E]) Replace(i int, e PrioritizedStrategy[E]) {
	p.entries[i] = e
	p.fix(i)
}

// ReplaceItem replaces the priority item at i.
func (p Prioritized[E]) ReplaceItem(i int, e E) {
	p.entries[i].e = e
	p.fix(i)
}

// Replace the Strategy at i.
func (p Prioritized[E]) ReplaceStrategy(i int, s directsearch.Strategy) { p.entries[i].Strategy = s }

// Fix the priority at i.
func (p Prioritized[E]) fix(i int) { heap.Fix((*byPriority[E])(&p.entries), i) }

type byPriority[E Priority] []PrioritizedStrategy[E]

func (a byPriority[E]) Len() int { return len(a) }
func (a byPriority[E]) Less(i, j int) bool {
	return a[i].e.Priority() < a[j].e.Priority()
}
func (a byPriority[E]) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a *byPriority[E]) Push(x any)   { *a = append(*a, x.(PrioritizedStrategy[E])) }
func (a *byPriority[E]) Pop() any     { n := len(*a) - 1; x := (*a)[n]; *a = (*a)[:n]; return x }
