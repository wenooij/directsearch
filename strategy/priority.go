package strategy

import (
	"container/heap"

	"github.com/wenooij/directsearch"
)

type PriorityEntry struct {
	Priority float64
	Strategy directsearch.Strategy
}

type Priority struct{ Entries []PriorityEntry }

func (p Priority) Next() directsearch.Action { return p.Entries[0].Strategy.Next() }

func (p Priority) Fix(i int)            { heap.Fix((*byPriority)(&p.Entries), i) }
func (p Priority) Init()                { heap.Init((*byPriority)(&p.Entries)) }
func (p *Priority) Pop() PriorityEntry  { return heap.Pop((*byPriority)(&p.Entries)).(PriorityEntry) }
func (p Priority) Push(x PriorityEntry) { heap.Push((*byPriority)(&p.Entries), x) }
func (p Priority) Remove(i int) PriorityEntry {
	return heap.Remove((*byPriority)(&p.Entries), i).(PriorityEntry)
}

func (p Priority) DecreasePriority(i0 int) bool {
	h := byPriority(p.Entries)
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

func (p Priority) IncreasePriority(j int) {
	h := byPriority(p.Entries)
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

type byPriority []PriorityEntry

func (a byPriority) Len() int           { return len(a) }
func (a byPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }
func (a byPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *byPriority) Push(x any)        { *a = append(*a, x.(PriorityEntry)) }
func (a *byPriority) Pop() any          { n := len(*a) - 1; x := (*a)[n]; *a = (*a)[:n]; return x }
