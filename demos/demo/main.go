// Demo code for the bar chart primitive.
package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tgraph"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	barGraph := tgraph.NewBarChart()
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")
	// display system metric usage
	barGraph.AddBar("cpu", 8, tcell.ColorBlue)
	barGraph.AddBar("mem", 2, tcell.ColorRed)
	barGraph.AddBar("swap", 4, tcell.ColorGreen)
	barGraph.AddBar("disk", 4, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	barGraph.SetRect(10, 10, 50, 20)

	update := func() {
		rand.Seed(time.Now().UnixNano())
		tick := time.NewTicker(1000 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				rangeLower := 0
				rangeUpper := 10
				randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("cpu", randomNum)
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("memory", randomNum)
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("swap", randomNum)
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("disk", randomNum)
				app.Draw()
			}
		}
	}
	go update()

	if err := app.SetRoot(barGraph, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
