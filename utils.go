package tvxwidgets

import "github.com/gdamore/tcell/v2"

const (
	// gauge cell
	prgCell = "▉"
)

// getColorName returns convert tcell color to its name
func getColorName(color tcell.Color) string {
	for name, c := range tcell.ColorNames {
		if c == color {
			return name
		}
	}
	return ""
}
