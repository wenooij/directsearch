package strategy

import (
	"testing"

	"github.com/wenooij/directsearch/actions"
)

func TestGreedy(t *testing.T) {
	g := Greedy(
		Monotone(actions.Int(0)),
		Monotone(actions.Int(1)),
		Monotone(actions.Int(2)),
	)
	for a := range Limit(10, Flatten(g)).Strategy() {
		t.Log(a)
	}
}
