package game

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
	userQuit = false

	// demo
	demoRNG   = rand.New(rand.NewSource(1))
	demoSpeed = time.Second / 5

	// configurations
	maxCols  = 80
	maxRows  = 40
	numCols  = 15
	numMines = roundToNearest10th((numRows * numCols) / 5)
	numRows  = 15
)

// Play starts the game.
func Play() {
	defer cleanup()
	generateMineField()

	// render forever in a separate routine
	chStop := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go renderAsync(chStop, &wg)

	for {
		if mf.IsGameOver() || userQuit {
			break
		}

		if *flagDemo {
			demo()
		} else {
			getUserInput()
		}
	}

	renderGame() // one final render
	chStop <- true
	wg.Wait()
}

func demo() {
	moveAndDo := func(x, y int, f func(x, y int)) {
		moveCursorTo(x, y, demoSpeed/2)
		f(x, y)
		time.Sleep(demoSpeed)
	}

	for mf.NumMinesRemaining() > 0 {
		val := demoRNG.Intn(*flagNumRows * *flagNumCols)
		x, y := val/(*flagNumCols), val%(*flagNumCols)

		if mf.HasMineAt(x, y) {
			if mf.Grid[x][y] != minefield.Flagged {
				moveAndDo(x, y, mf.Flag)
			}
		} else if mf.Grid[x][y] == minefield.Unknown {
			moveAndDo(x, y, mf.Reveal)
		}
	}
}

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
		return
	}
	if *flagDemo && key != keyboard.KeyEsc && key != keyboard.KeyCtrlC && char != 'q' && char != 'Q' {
		return
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
