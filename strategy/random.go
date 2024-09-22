package strategy

import (
	"math/rand/v2"

	"github.com/wenooij/directsearch"
)

type Random struct {
	Actions []directsearch.Action
	Rand    *rand.Rand
}

func (r *Random) Next() directsearch.Action {
	return r.Actions[r.Rand.IntN(len(r.Actions))]
}
