// Demo code for the bar chart primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tgraph"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	barGraph := tgraph.NewBarChart()
	barGraph.SetRect(4, 2, 50, 20)
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")
	// display system metric usage
	barGraph.AddBar("cpu", 80, tcell.ColorBlue)
	barGraph.AddBar("mem", 20, tcell.ColorRed)
	barGraph.AddBar("swap", 40, tcell.ColorGreen)
	barGraph.AddBar("disk", 40, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	if err := app.SetRoot(barGraph, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
