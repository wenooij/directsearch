package strategy

import "github.com/wenooij/directsearch"

type Monotone struct{ Action directsearch.Action }

func (a Monotone) Next() directsearch.Action { return a.Action }
