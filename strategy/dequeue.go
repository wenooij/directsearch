package strategy

import (
	"iter"

	"github.com/wenooij/directsearch"
)

// Dequeue is a simple list of strategies which supports taking from the front or back.
type Dequeue []directsearch.Strategy

func (q Dequeue) Len() int { return len(q) }

func (q Dequeue) At(i int) directsearch.Strategy { return q[i] }

func (q *Dequeue) Append(es ...directsearch.Strategy) { *q = append(*q, es...) }

func (q Dequeue) Front() directsearch.Strategy { return FlattenStrategies(q...) }

func (q Dequeue) reverse() iter.Seq[directsearch.Strategy] {
	return func(yield func(directsearch.Strategy) bool) {
		for i := q.Len() - 1; i >= 0; i-- {
			if !yield(q.At(i)) {
				break
			}
		}
	}
}

func (q Dequeue) Back() directsearch.Strategy { return FlattenStrategiesIter(q.reverse()) }
