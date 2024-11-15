package directsearch

import "iter"

// Action represents an arbitrary change to the environment
// with a unique string representation.
//
// Action is the unique identifier for the Action and should not change.
type Action interface {
	Action() string
}

// State represents the serializable state of an Environment.
//
// String is a human readable description of the state that should also
// uniquely identify the state.
type State interface {
	State() string
}

// Environment represents an interactive environment with a State
// and a means of enacting Actions on the state with a resulting Reward.
//
// The Reward will contain information related to the Reward received
// after doing the Action on the given Environment. Reward.Valid may be
// false if the Action is not valid in the given context.
type Environment interface {
	State
	Do(Action) Reward
}

// Strategy is an iterator which yields Actions.
//
// When defined in terms of the current Environment state,
// Strategy provides an ideal abstraction for programming the agent.
//
// Abstractions in the strategy subpackage will make implementing common
// pattens even easier.
type Strategy interface {
	Strategy() iter.Seq[Action]
}

// Agent defines a player who interacts with the environment
// and is capable of learning through Reward updates.
//
// Depending on the Agent implementation, Update may be called
// periodically after Actions are performed, or in a batched way.
type Agent interface {
	Strategy() Strategy
	Update(Reward)
}

// Reward represents a value which is opaque with respect to any
// given problem, and a flag which determines it the reward value
// is valid or not.
//
// The Value slice provides a way to represent multiple scalar objective dimensions
// and it is up to the particular agent to score them appropriately.
// Values may correspond to minmax player scores or multiple objectives in optimazation.
//
// Typically a Valid is set to false when an illegal Action was performed
// on an Environment. The Environment may decide which Reward to deliver
// to the agent which submitted the illeal action and the resulting change
// of state.
//
// If Valid is true and Value is nil the Agent may interpret missing values as zeroes.
type Reward struct {
	Valid bool
	Value []float64
}
