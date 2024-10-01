package strategy

import "github.com/wenooij/directsearch"

// Monotone represents a strategy which always returns the specified Action.
type Monotone struct{ Action directsearch.Action }

// Next returns the specified monotone Action.
func (a Monotone) Next() directsearch.Action { return a.Action }
