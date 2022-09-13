package game

import (
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-mines/digital"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	titleText1 = `
   _____   __                      
  /     \ |__| ____   ____   ______
 /  \ /  \|  |/    \_/ __ \ /  ___/
/    \    \  |   |  \  ___/_\___ \ 
\____/\_  /__|___|  /\___  /____  \
        \/        \/     \/     \/ 
 
`
	titleText = titleText1
)

var (
	colorKey   = text.Colors{text.Italic, text.FgWhite}
	colorTitle = []text.Color{
		text.FgHiWhite,
		text.FgHiYellow,
		text.FgHiCyan,
		text.FgHiBlue,
		text.FgHiGreen,
		text.FgHiMagenta,
		text.FgHiRed,
	}
	colorTitleAnimated = false
	colorTitleIdx      = 0
	colorTitleOnce     = sync.Once{}

	spacesToPad = -1
	titleOnce   sync.Once
)

func initHeaderAndFooter() {
	titleOnce.Do(func() {
		// title Colors
		if colorTitleAnimated {
			colorTitleOnce.Do(func() {
				go func() {
					for {
						time.Sleep(time.Second / 2)
						if colorTitleIdx < len(colorTitle)-1 {
							colorTitleIdx++
						} else {
							colorTitleIdx = 0
						}
					}
				}()
			})
		}
	})

	spacesToPad = (*flagNumCols * 5) - (9 /*mines*/ + 9 /*turns*/ + text.LongestLineLen(titleText)) - 13 /*magic*/
}

func renderFooter() string {
	return colorKey.Sprint("flag: <F/f>  | navigate: ▶ ▲ ▼ ◀  | reveal: <space> | quit: <Q/q>")
}

func renderHeader() string {
	gameTimeSeconds := time.Now().Sub(timeStart).Seconds()
	title := renderTitle()
	numMinesRemaining := digital.Number(int64(mf.NumMinesRemaining()), digital.WithMinChars(3))
	numSeconds := digital.Number(int64(gameTimeSeconds), digital.WithMinChars(3))
	padding := ""
	if spacesToPad > 0 {
		padding = text.Pad("", spacesToPad/2, ' ')
	}

	tw := table.NewWriter()
	tw.AppendRow(table.Row{numMinesRemaining, padding, title, padding, numSeconds})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, VAlign: text.VAlignBottom},
		{Number: 5, VAlign: text.VAlignBottom},
	})
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.DrawBorder = false
	tw.Style().Options.SeparateColumns = false
	return tw.Render()
}

func renderTitle() string {
	colors := text.Colors{colorTitle[colorTitleIdx], text.Bold}
	tw := table.NewWriter()
	for _, line := range strings.Split(titleText, "\n") {
		if line != "" {
			tw.AppendRow(table.Row{colors.Sprint(line)})
		}
	}
	tw.Style().Options = table.OptionsNoBordersAndSeparators
	return tw.Render()
}
