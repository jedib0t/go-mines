package minefield

import (
	"math/rand"
)

// Minefield describes a minefield and allows operations on the same.
type Minefield struct {
	Grid [][]State

	boom              bool
	mines             []Position
	numMines          int
	numMinesAround    map[Position]int
	numFlagged        int
	numFlaggedWrongly int
	numTurns          int
	numRows           int
	numCols           int
	rng               *rand.Rand
}

// New returns a new Minefield with the given options applied on top of some
// sensible defaults.
func New(opts ...Option) *Minefield {
	mf := &Minefield{}

	// apply all options
	for _, opt := range append(defaultOptions, opts...) {
		opt(mf)
	}

	// manufacture the grid
	mf.Grid = make([][]State, mf.numRows)
	for x := 0; x < mf.numRows; x++ {
		mf.Grid[x] = make([]State, mf.numCols)
		for y := 0; y < mf.numCols; y++ {
			mf.Grid[x][y] = Unknown
		}
	}

	// place the mines
	if mf.numMines > (mf.numRows * mf.numCols) {
		mf.numMines = mf.numRows * mf.numCols
	}
	for idx := 0; idx < mf.numMines; idx++ {
		for {
			val := mf.rng.Intn(mf.numRows * mf.numCols)
			x, y := val/mf.numCols, val%mf.numCols
			if !mf.HasMineAt(x, y) {
				mf.mines = append(mf.mines, Position{X: x, Y: y})
				break
			}
		}
	}

	// pre-calculate the number of mines around every position to make look-ups
	// fast
	mf.numMinesAround = make(map[Position]int, mf.numRows*mf.numCols)
	for x := 0; x < mf.numRows; x++ {
		for y := 0; y < mf.numCols; y++ {
			mf.numMinesAround[Position{X: x, Y: y}] = mf.calcNumMinesAround(x, y)
		}
	}
	return mf
}

// AllMinesFound returns true if all the mines in the field have been flagged
// and none incorrectly flagged.
func (mf *Minefield) AllMinesFound() bool {
	return mf.NumMinesRemaining() == 0 && mf.numFlaggedWrongly == 0
}

// Flag places a flag on a position marking that the user thinks there is a mine
// at that location.
func (mf *Minefield) Flag(x, y int) {
	if (x < 0 || x >= mf.numRows) || (y < 0 || y >= mf.numCols) {
		return
	}
	if mf.Grid[x][y] == Empty {
		return
	}

	mf.numTurns++
	if mf.Grid[x][y] != Flagged { // ensure it is not flagged at the moment
		mf.Grid[x][y] = Flagged
		mf.numFlagged++
		if !mf.HasMineAt(x, y) { // count if incorrectly flagged
			mf.numFlaggedWrongly++
		}
	} else { // already flagged? then toggle value.
		mf.Grid[x][y] = Unknown
		mf.numFlagged--
		if !mf.HasMineAt(x, y) { // revert count if incorrectly flagged
			mf.numFlaggedWrongly--
		}
	}
}

// HasMineAt returns true if there is a mine at the given location.
func (mf *Minefield) HasMineAt(x, y int) bool {
	for _, pos := range mf.mines {
		if pos.X == x && pos.Y == y {
			return true
		}
	}
	return false
}

// IsGameOver returns true if all mines have been found, and if a mine was
// accidentally triggered.
func (mf *Minefield) IsGameOver() bool {
	return mf.AllMinesFound() || mf.boom
}

// NumFlagged returns the number of locations flagged.
func (mf *Minefield) NumFlagged() int {
	return mf.numFlagged
}

// NumMinesAround returns the number of mines around the given location.
func (mf *Minefield) NumMinesAround(x, y int) int {
	return mf.numMinesAround[Position{X: x, Y: y}]
}

// NumMinesRemaining returns the number of mines that have not been flagged.
func (mf *Minefield) NumMinesRemaining() int {
	return mf.numMines - mf.numFlagged
}

// NumTurns returns the number of operations performed on the minefield.
func (mf *Minefield) NumTurns() int {
	return mf.numTurns
}

// Reveal reveals the given location. Note that the game is over if this is done
// on a position that has a mine.
func (mf *Minefield) Reveal(x, y int) {
	if (x < 0 || x >= mf.numRows) || (y < 0 || y >= mf.numCols) {
		return
	}
	if mf.Grid[x][y] == Empty {
		return
	}

	mf.numTurns++
	if mf.HasMineAt(x, y) { // game over
		mf.Grid[x][y] = Boom
		mf.boom = true
	} else {
		mf.revealEmpty(x, y)
	}
}

func (mf *Minefield) calcNumMinesAround(x, y int) int {
	rsp := 0
	for _, pos := range mf.mines {
		if (pos.X == x-1 && pos.Y == y-1) ||
			(pos.X == x && pos.Y == y-1) ||
			(pos.X == x+1 && pos.Y == y-1) ||
			(pos.X == x-1 && pos.Y == y) ||
			(pos.X == x+1 && pos.Y == y) ||
			(pos.X == x-1 && pos.Y == y+1) ||
			(pos.X == x && pos.Y == y+1) ||
			(pos.X == x+1 && pos.Y == y+1) {
			rsp++
		}
	}
	return rsp
}

func (mf *Minefield) revealEmpty(x, y int) {
	if (x < 0 || x >= mf.numRows) || (y < 0 || y >= mf.numCols) {
		return
	}
	if mf.Grid[x][y] == Empty {
		return
	}
	if mf.HasMineAt(x, y) {
		return
	}
	mf.Grid[x][y] = Empty
	if mf.NumMinesAround(x, y) > 0 {
		return
	}

	mf.revealEmpty(x-1, y-1)
	mf.revealEmpty(x, y-1)
	mf.revealEmpty(x+1, y-1)
	mf.revealEmpty(x-1, y)
	mf.revealEmpty(x+1, y)
	mf.revealEmpty(x-1, y+1)
	mf.revealEmpty(x, y+1)
	mf.revealEmpty(x+1, y+1)
}
