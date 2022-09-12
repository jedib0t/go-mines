package digital

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

// Number returns the given number in the style of a number displayed on a
// digital display.
func Number(n int64, opts ...Option) string {
	cfg := &config{}
	for _, opt := range append(defaultOptions, opts...) {
		opt(cfg)
	}
	isNegative := n < 0
	if n < 0 { // make it positive for the loop logic below
		n *= -1
	}

	// use division to find last digit and "prefix" row in a loop
	twRow := table.Row{}
	for n > -1 {
		q, r := n/10, n%10
		twRow = append(table.Row{getDigitalNumber(int(r))}, twRow...)
		n = q
		if n == 0 {
			break
		}
	}
	// prefix with '0's as much as needed
	for (isNegative && len(twRow) < cfg.minChars-1) || (!isNegative && len(twRow) < cfg.minChars) {
		twRow = append(table.Row{getDigitalNumber(0)}, twRow...)
	}
	// prefix with '-' if negative
	if isNegative {
		twRow = append(table.Row{getDigitalNumber(-1)}, twRow...)
	}

	tw := table.NewWriter()
	tw.AppendRow(twRow)
	tw.SetStyle(cfg.style)
	tw.Style().Color.Row = cfg.colors
	tw.Style().Box.PaddingLeft = cfg.padding
	tw.Style().Box.PaddingRight = cfg.padding
	return tw.Render()
}
