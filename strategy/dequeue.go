package strategy

import (
	"iter"
	"slices"

	"github.com/wenooij/directsearch"
)

// Dequeue is a simple list of strategies which supports taking from the front or back.
type Dequeue []directsearch.Strategy

func (q Dequeue) Len() int { return len(q) }

func (q Dequeue) At(i int) directsearch.Strategy { return q[i] }

func (q *Dequeue) Append(es ...directsearch.Strategy) { *q = append(*q, es...) }

func (q Dequeue) Front() directsearch.Strategy {
	next, _ := iter.Pull(slices.Values(q))
	return Flatten(func() (directsearch.Strategy, bool) {
		e, ok := next()
		if !ok {
			return nil, false
		}
		return e, true
	})
}

func (q Dequeue) Back() directsearch.Strategy {
	i := q.Len() - 1
	return Flatten(func() (directsearch.Strategy, bool) {
		if i < 0 || q.Len() <= i {
			return nil, false
		}
		e := q.At(i)
		i--
		return e, true
	})
}
