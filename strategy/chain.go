package strategy

import "github.com/wenooij/directsearch"

// Chain the strategies together until each is exhausted.
//
// Chain is an alias for FlattenStrategies(s...).
func Chain(s ...directsearch.Strategy) directsearch.Strategy { return FlattenStrategies(s...) }
