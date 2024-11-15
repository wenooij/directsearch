package strategy

import (
	"testing"

	"github.com/wenooij/directsearch/actions"
)

func TestCoalesce(t *testing.T) {
	s := Coalesce(
		First(Monotone(actions.Int(0))),
		First(Monotone(actions.Int(1))),
		First(Monotone(actions.Int(2))),
	)
	for a := range s.Strategy() {
		t.Log(a)
	}
}
