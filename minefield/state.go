package minefield

// State defines the state of a single Position in a Minefield.
type State int

// State values.
const (
	Unknown State = iota
	Empty
	Flagged
	Boom
)
