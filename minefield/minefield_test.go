package minefield

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mf := New(WithSeed(1))
	assert.Equal(t, defaultNumCols, mf.numCols)
	assert.Equal(t, defaultNumMines, mf.numMines)
	assert.Equal(t, defaultNumRows, mf.numRows)
	assert.Len(t, mf.mines, defaultNumMines)
	assert.Len(t, mf.numMinesAround, mf.numRows*mf.numCols)
}

func TestMinefield_AllMinesFound(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	assert.False(t, mf.AllMinesFound())

	x, y := mf.mines[0].X, mf.mines[0].Y
	mf.Flag(x, y)
	assert.True(t, mf.AllMinesFound())
}

func TestMinefield_Flag(t *testing.T) {
	mf := New(WithSeed(1))
	assert.Equal(t, Unknown, mf.Grid[0][0])

	x, y := mf.mines[0].X, mf.mines[0].Y
	mf.Flag(x, y)
	assert.Equal(t, Flagged, mf.Grid[x][y])
}

func TestMinefield_HasMineAt(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	assert.False(t, mf.HasMineAt(0, 0))
	assert.True(t, mf.HasMineAt(4, 1))
}

func TestMinefield_IsGameOver(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	assert.False(t, mf.IsGameOver())
	mf.Reveal(4, 1)
	assert.False(t, mf.AllMinesFound())
	assert.True(t, mf.IsGameOver())

	mf = New(WithNumMines(1), WithSeed(1))
	assert.False(t, mf.IsGameOver())
	mf.Flag(4, 1)
	assert.True(t, mf.AllMinesFound())
	assert.True(t, mf.IsGameOver())
}

func TestMinefield_NumFlagged(t *testing.T) {
	mf := New()
	mf.Flag(1, 1)
	assert.Equal(t, 1, mf.NumFlagged())
	mf.Flag(1, 2)
	assert.Equal(t, 2, mf.NumFlagged())
	mf.Flag(1, 3)
	assert.Equal(t, 3, mf.NumFlagged())
}

func TestMinefield_NumMinesAround(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	assert.Equal(t, 0, mf.NumMinesAround(1, 1))

	assert.Equal(t, 1, mf.NumMinesAround(3, 0))
	assert.Equal(t, 1, mf.NumMinesAround(3, 1))
	assert.Equal(t, 1, mf.NumMinesAround(3, 2))
	assert.Equal(t, 1, mf.NumMinesAround(4, 0))
	assert.Equal(t, 0, mf.NumMinesAround(4, 1))
	assert.Equal(t, 1, mf.NumMinesAround(4, 2))
	assert.Equal(t, 1, mf.NumMinesAround(5, 0))
	assert.Equal(t, 1, mf.NumMinesAround(5, 1))
	assert.Equal(t, 1, mf.NumMinesAround(5, 2))
}

func TestMinefield_NumMinesRemaining(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	assert.Equal(t, 1, mf.NumMinesRemaining())
	mf.Flag(4, 1)
	assert.Equal(t, 0, mf.NumMinesRemaining())
}

func TestMinefield_NumTurns(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	assert.Equal(t, 0, mf.NumTurns())
	mf.Flag(4, 1)
	assert.Equal(t, 1, mf.NumTurns())
	mf.Reveal(5, 1)
	assert.Equal(t, 2, mf.NumTurns())
}

func TestMinefield_Reveal(t *testing.T) {
	mf := New(WithNumMines(1), WithSeed(1))
	mf.Reveal(1, 1)
	assert.Equal(t, Empty, mf.Grid[1][1])
	mf.Reveal(4, 1)
	assert.Equal(t, Boom, mf.Grid[4][1])
}
