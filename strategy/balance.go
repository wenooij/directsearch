package strategy

import "math"

type Balance struct {
	Priority
}

func (b Balance) AddRuns(i, n int) {
	if n <= 0 {
		panic("AddRuns: n <= 0")
	}
	e := b.Entries[i]
	bs := e.PriorityInterface.(BalanceeStats)
	bs.Runs += float64(n)
	e.PriorityInterface = bs
	b.Entries[i] = e
	b.Priority.DecreasePriority(i)
}

type BalanceeStats struct {
	Runs      float64
	Score     []float64
	Objective func([]float64) float64
}

func (s BalanceeStats) Priority() float64 {
	if s.Runs == 0 {
		return math.Inf(-1)
	}
	return -s.Objective(s.Score) / s.Runs
}
