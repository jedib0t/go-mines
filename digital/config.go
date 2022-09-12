package digital

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type config struct {
	colors   text.Colors
	minChars int
	padding  string
	style    table.Style
}
