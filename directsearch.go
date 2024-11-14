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

// Strategy represents an iterator which yields Actions.
type Strategy interface {
	Strategy() iter.Seq[Action]
}

// MetaStrategy represents an interactor which yields new strategies.
//
// It is possible to implement MetaStrategies that yield MetaStrategies by
// wrapping the returned MetaStrategies using the [strategies.Flat] wrapper.
type MetaStrategy interface {
	MetaStrategy() iter.Seq[Strategy]
}

// Agent defines a player who interacts with the environment
// and is capable of learning through Reward updates.
type Agent interface {
	Strategy() Strategy
	Update(Reward)
}

// Reward represents a value which is opaque with respect to any
// given problem, and a flag which determines it the reward value
// is valid or not.
//
// Typically a Valid is set to false when an illegal Action was performed
// on an Environment. The Environment may decide which Reward to deliver
// to the agent which submitted the illeal action and the resulting change
// of state.
type Reward struct {
	Valid bool
	Value []float64
}
