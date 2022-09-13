package game

import "github.com/jedib0t/go-mines/minefield"

const (
	symbolFlaggedWrong = "🅧"
	symbolMine         = "💣"
)

var (
	symbolSelected = "⬛"
	symbolStateMap = map[minefield.State]string{
		minefield.Unknown: "⬜",
		minefield.Empty:   "  ",
		minefield.Flagged: "🚩",
		minefield.Boom:    "💥",
	}
	symbolNumberShadedMap = map[int]string{
		0: "⬛",
		1: "➊",
		2: "➋",
		3: "➌",
		4: "➍",
		5: "➎",
		6: "➏",
		7: "➐",
		8: "➑",
	}
	symbolNumberMap = map[int]string{
		0: "⬜",
		1: "①",
		2: "②",
		3: "③",
		4: "④",
		5: "⑤",
		6: "⑥",
		7: "⑦",
		8: "⑧",
	}
)
