package strategy

import "github.com/wenooij/directsearch"

// Buffer allows for replay and rewind of strategies.
type Buffer struct {
	actions    []directsearch.Action
	strategies []directsearch.Strategy

	pos int
}

func (b *Buffer) Reset() { b.actions = b.actions[:0]; b.strategies = b.strategies[:0]; b.pos = 0 }

func (b *Buffer) Append(es ...directsearch.Strategy) { b.strategies = append(b.strategies, es...) }
