package strategy

import "math"

// Balance implements an explore-exploit algorithm to select its strategies.
type Balance struct {
	Priority[BalanceStats]
}

func (b Balance) Runs(i int) float64 { return b.Entries[i].Item.Runs }

func (b Balance) SetRuns(i int, runs float64) {
	if runs <= 0 {
		panic("runs <= 0")
	}
	b.Entries[i].Item.addRuns(runs)
	b.Entries[i].Item.recomputePriority()
	b.Priority.Fix(i)
}

func (b Balance) AddRuns(i int, runs float64) {
	if runs <= 0 {
		panic("runs <= 0")
	}
	b.Entries[i].Item.addRuns(runs)
	b.Entries[i].Item.recomputePriority()
	b.Priority.DecreasePriority(i)
}

func (b Balance) AddScore(i int, score []float64) {
	b.Entries[i].Item.addScore(score)
	b.Entries[i].Item.recomputePriority()
	b.Priority.Fix(i)
}

func (b Balance) AddScoreRuns(i int, score []float64, runs int) {
	if runs <= 0 {
		panic("runs <= 0")
	}
	b.Entries[i].Item.addScore(score)
	b.Entries[i].Item.addRuns(float64(runs))
	b.Entries[i].Item.recomputePriority()
	b.Priority.Fix(i)
}

type BalanceStats struct {
	Runs          float64
	Score         []float64
	Objective     func([]float64) float64
	PriorityValue float64
}

func (s BalanceStats) Priority() float64 { return s.PriorityValue }

func (s *BalanceStats) addScore(score []float64) {
	n := len(s.Score) - 1
	_ = score[n] // Bounds check.
	for i := 0; i <= n; i++ {
		s.Score[i] += score[i]
	}
}

func (s *BalanceStats) addRuns(n float64) { s.Runs += n }

func (s *BalanceStats) recomputePriority() {
	if s.Runs == 0 {
		s.PriorityValue = math.Inf(-1)
		return
	}
	s.PriorityValue = -s.Objective(s.Score) / s.Runs
}
