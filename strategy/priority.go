package strategy

import (
	"container/heap"

	"github.com/wenooij/directsearch"
)

type PriorityInterface interface {
	Priority() float64
}

type PriorityEntry[T PriorityInterface] struct {
	Value T
	directsearch.Strategy
}

type Priority[T PriorityInterface] struct{ Entries []PriorityEntry[T] }

func (p Priority[T]) Next() directsearch.Action { return p.Entries[0].Next() }

func (p Priority[T]) Fix(i int) { heap.Fix((*byPriority[T])(&p.Entries), i) }
func (p Priority[T]) Init()     { heap.Init((*byPriority[T])(&p.Entries)) }
func (p *Priority[T]) Pop() PriorityEntry[T] {
	return heap.Pop((*byPriority[T])(&p.Entries)).(PriorityEntry[T])
}
func (p Priority[T]) Push(x PriorityEntry[T]) { heap.Push((*byPriority[T])(&p.Entries), x) }
func (p Priority[T]) Remove(i int) PriorityEntry[T] {
	return heap.Remove((*byPriority[T])(&p.Entries), i).(PriorityEntry[T])
}

func (p Priority[T]) DecreasePriority(i0 int) bool {
	h := byPriority[T](p.Entries)
	n := len(p.Entries)
	// Adapted from heap.down.
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

func (p Priority[T]) IncreasePriority(j int) {
	h := byPriority[T](p.Entries)
	// Adapted from heap.up.
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

type byPriority[T PriorityInterface] []PriorityEntry[T]

func (a byPriority[T]) Len() int           { return len(a) }
func (a byPriority[T]) Less(i, j int) bool { return a[i].Value.Priority() < a[j].Value.Priority() }
func (a byPriority[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *byPriority[T]) Push(x any)        { *a = append(*a, x.(PriorityEntry[T])) }
func (a *byPriority[T]) Pop() any          { n := len(*a) - 1; x := (*a)[n]; *a = (*a)[:n]; return x }
