package directsearch

type Action interface {
	String() string
}

type State interface {
	String() string
}

type Environment interface {
	State() State
	Do(Action) (Reward, error)
}

type Strategy interface {
	Next() Action
}

type Agent interface {
	Strategy() Strategy
	Update(Reward)
}

type Reward struct {
	Score []float64
}
