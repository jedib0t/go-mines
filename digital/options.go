package digital

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Option helps customize the operations of the functions in this package.
type Option func(c *config)

var (
	defaultOptions = []Option{
		WithColors(text.Colors{text.BgBlack, text.FgHiRed}),
		WithMinChars(3),
		WithPadding(""),
		WithTableStyle(table.StyleColoredDark),
	}
)

// WithColors defines the colors used to print the numbers/letters.
func WithColors(colors text.Colors) Option {
	return func(c *config) {
		c.colors = colors
	}
}

// WithMinChars defines the minimum number of characters in the board.
func WithMinChars(n int) Option {
	return func(c *config) {
		c.minChars = n
	}
}

// WithPadding defines the padding used for each character.
func WithPadding(padding string) Option {
	return func(c *config) {
		c.padding = padding
	}
}

// WithTableStyle defines the table.Style to be used for the table used.
func WithTableStyle(s table.Style) Option {
	return func(c *config) {
		c.style = s
	}
}
