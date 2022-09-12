package minefield

import (
	"math/rand"
	"time"
)

var (
	defaultNumCols  = 20
	defaultNumRows  = 20
	defaultNumMines = (defaultNumRows * defaultNumCols) / 4
	defaultSeed     = time.Now().UnixNano()
)

// Option helps customize the Minefield.
type Option func(mf *Minefield)

var (
	defaultOptions = []Option{
		WithNumCols(defaultNumCols),
		WithNumMines(defaultNumMines),
		WithNumRows(defaultNumRows),
		WithSeed(defaultSeed),
	}
)

// WithNumMines defines the number of mines to be placed.
func WithNumMines(numMines int) Option {
	return func(mf *Minefield) {
		mf.numMines = numMines
	}
}

// WithNumCols defines the number of columns.
func WithNumCols(numCols int) Option {
	return func(mf *Minefield) {
		mf.numCols = numCols
	}
}

// WithNumRows defines the number of rows.
func WithNumRows(numRows int) Option {
	return func(mf *Minefield) {
		mf.numRows = numRows
	}
}

// WithSeed defines the seed value for the RNG used.
func WithSeed(seed int64) Option {
	return func(mf *Minefield) {
		mf.rng = rand.New(rand.NewSource(seed))
	}
}
