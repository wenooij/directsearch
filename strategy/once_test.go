package strategy

import (
	"testing"

	"github.com/wenooij/directsearch/actions"
)

func TestOnce(t *testing.T) {
	s := First(Monotone(actions.Int(1)))

	for a := range s.Strategy() {
		t.Log(a.Action())
	}
}
