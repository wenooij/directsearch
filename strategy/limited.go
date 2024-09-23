package strategy

import "github.com/wenooij/directsearch"

type Limited struct {
	Strategy directsearch.Strategy
	Count    int
	Limit    int
}

func (l Limited) Next() directsearch.Action {
	if l.Count < l.Limit {
		l.Count++
		return l.Strategy.Next()
	}
	return nil
}
