package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-mines/minefield"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	// colors
	colorsFlaggedWrong = text.Colors{text.FgHiRed}

	// misc
	linesRendered = 0
	renderedGame  = ""

	// controls
	renderEnabled = true
	renderMutex   = sync.Mutex{}
)

func renderAsync(chStop chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.Tick(time.Second / time.Duration(*flagRefreshRate))
	for {
		select {
		case <-chStop: // render one final time and return
			renderGame()
			return
		case <-timer: // render as part of regular cycle
			renderGame()
		}
	}
}

func renderGame() {
	renderMutex.Lock()
	defer renderMutex.Unlock() // unnecessary
	if !renderEnabled {
		return
	}

	style := table.StyleColoredBlackOnBlueWhite
	if mf.AllMinesFound() {
		style = table.StyleColoredBlackOnGreenWhite
	} else if mf.IsGameOver() {
		style = table.StyleColoredBlackOnRedWhite
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"\n" + renderHeader() + "\n"})
	tw.AppendRow(table.Row{renderMineField()})
	tw.AppendFooter(table.Row{"\n" + renderFooter() + "\n"})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
	})
	tw.SetStyle(style)
	tw.Style().Format.Footer = text.FormatDefault
	tw.Style().Title.Align = text.AlignCenter

	out := tw.Render()
	if out != renderedGame {
		for linesRendered > 0 {
			fmt.Print(text.CursorUp.Sprint())
			fmt.Print(text.EraseLine.Sprint())
			linesRendered--
		}

		linesRendered = strings.Count(out, "\n") + 1
		fmt.Println(out)
		renderedGame = out
	}
}

func renderMineField() string {
	twMineField := table.NewWriter()
	for x, row := range mf.Grid {
		twRow := table.Row{}
		for y, col := range row {
			numMinesAround := mf.NumMinesAround(x, y)

			colStr := symbolStateMap[col]
			if userQuit {
				if col == minefield.Flagged && !mf.HasMineAt(x, y) {
					colStr = colorsFlaggedWrong.Sprint(symbolFlaggedWrong)
				} else if mf.HasMineAt(x, y) {
					colStr = symbolMine
				}
			} else if x == cursor.X && y == cursor.Y {
				switch col {
				case minefield.Unknown:
					colStr = symbolSelected
				case minefield.Empty:
					colStr = symbolNumberShadedMap[numMinesAround]
				default:
					colStr = symbolStateMap[col]
				}
			} else {
				if col == minefield.Empty {
					if numMinesAround > 0 {
						colStr = symbolNumberMap[numMinesAround]
					}
				} else if mf.IsGameOver() && col == minefield.Flagged && !mf.HasMineAt(x, y) {
					colStr = colorsFlaggedWrong.Sprint(symbolFlaggedWrong)
				} else if mf.IsGameOver() && mf.HasMineAt(x, y) {
					if col != minefield.Boom && col != minefield.Flagged {
						colStr = symbolMine
					}
				}
			}
			twRow = append(twRow, colStr)
		}
		twMineField.AppendRow(twRow)
	}
	twMineField.SetStyle(table.StyleRounded)
	twMineField.Style().Options.SeparateRows = true
	twMineField.Style().Color.RowAlternate = twMineField.Style().Color.Row
	return twMineField.Render()
}
