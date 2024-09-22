package directsearch

type Action any

type State struct {
	Text string
}

type Strategy interface {
	Next() Action
}

type Node interface {
	Load(State) error
	Result() (Result, bool)
}

type Result struct {
	Action Action
	Score  []float64
	State  State
}
