package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/jedib0t/go-mines/minefield"
)

var (
	// game state
	cursor   = minefield.Position{X: 0, Y: 0}
	mf       *minefield.Minefield
	userQuit bool

	// configurations
	maxCols  = 80
	maxRows  = 40
	numCols  = 15
	numMines = roundToNearest10th((numRows * numCols) / 5)
	numRows  = 15
)

func generateMineField() {
	mf = minefield.New(
		minefield.WithNumCols(*flagNumCols),
		minefield.WithNumMines(*flagNumMines),
		minefield.WithNumRows(*flagNumRows),
		minefield.WithSeed(*flagSeed),
	)
}

func getUserInput() {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		logErrorAndExit("failed to get input: %v", err)
	}

	switch key {
	case keyboard.KeyEsc, keyboard.KeyCtrlC:
		handleActionQuit()
	case keyboard.KeyCtrlR:
		handleActionReset()
	case keyboard.KeyArrowDown:
		if cursor.X+1 < *flagNumRows {
			cursor.X++
		}
	case keyboard.KeyArrowUp:
		if cursor.X-1 >= 0 {
			cursor.X--
		}
	case keyboard.KeyArrowRight:
		if cursor.Y+1 < *flagNumCols {
			cursor.Y++
		}
	case keyboard.KeyArrowLeft:
		if cursor.Y-1 >= 0 {
			cursor.Y--
		}
	case keyboard.KeySpace:
		mf.Reveal(cursor.X, cursor.Y)
	default:
		if char == 'f' || char == 'F' {
			mf.Flag(cursor.X, cursor.Y)
		} else if char == 'q' || char == 'Q' {
			handleActionQuit()
		} else {
			handleActionInput(char)
		}
	}
}

func play() {
	// render forever in a separate routine
	chStop := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go renderAsync(chStop, &wg)

	for {
		if mf.IsGameOver() || userQuit {
			break
		}

		if *flagWin {
			cheatAndWin()
		} else {
			getUserInput()
		}
	}

	renderGame() // one final render
	chStop <- true
	wg.Wait()
}

func cheatAndWin() {
	rng := rand.New(rand.NewSource(1))
	interval := time.Second / 2

	for mf.NumMinesRemaining() > 0 {
		val := rng.Intn(*flagNumRows * *flagNumCols)
		x, y := val/(*flagNumCols), val%(*flagNumCols)
		cursor = minefield.Position{X: x, Y: y}
		if mf.HasMineAt(x, y) {
			if mf.Grid[x][y] != minefield.Flagged {
				mf.Flag(x, y)
				time.Sleep(interval)
			}
		} else if mf.Grid[x][y] == minefield.Unknown {
			mf.Reveal(x, y)
			time.Sleep(interval)
		}
	}
}
