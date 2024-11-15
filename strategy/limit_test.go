package strategy

import (
	"testing"

	"github.com/wenooij/directsearch"
	"github.com/wenooij/directsearch/actions"
)

func TestLimit(t *testing.T) {
	i := 0
	s := Limit(10, MonotoneFunc(func() directsearch.Action {
		i++
		return actions.Int(i)
	}))

	for a := range s.Strategy() {
		t.Log(a.Action())
	}
}
