package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	flagHelp        = flag.Bool("help", false, "Show this help-text?")
	flagNumCols     = flag.Int("cols", numRows, fmt.Sprintf("Number of Columns (max: %d)", maxCols))
	flagNumMines    = flag.Int("mines", numMines, "Number of Mines (max: 50% of grid)")
	flagNumRows     = flag.Int("rows", numCols, fmt.Sprintf("Number of Rows (max: %d)", maxRows))
	flagRefreshRate = flag.Int("refresh-rate", 20, "Refresh-rate per second")
	flagSeed        = flag.Int64("seed", 0, "Randomizer Seed value (will use current time if ZERO)")
	flagWin         = flag.Bool("win", false, "Cheat and Win?")
)

func initFlags() {
	flag.Parse()
	if *flagHelp {
		printHelp()
	}

	// control rows/cols
	if *flagNumCols > maxCols {
		*flagNumCols = maxCols
	}
	if *flagNumRows > maxRows {
		*flagNumRows = maxRows
	}
	if (*flagNumCols != numCols || *flagNumRows != numRows) && *flagNumMines == numMines {
		*flagNumMines = roundToNearest10th(*flagNumRows * *flagNumCols / 5)
	}
	if *flagNumMines > (*flagNumRows**flagNumCols)/2 {
		*flagNumMines = roundToNearest10th((*flagNumRows * *flagNumCols) / 2)
	}
	if *flagSeed == 0 {
		*flagSeed = time.Now().Unix()
	}
	rand.Seed(*flagSeed)
}

func printHelp() {
	fmt.Println(`go-mines: A GoLang implementation of the Minesweeper game.

Version: ` + version + `

Flags
=====`)
	flag.PrintDefaults()
	os.Exit(0)
}
