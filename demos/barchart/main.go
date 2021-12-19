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
	barGraph.SetRect(0, 0, 50, 20)
	barGraph.SetBorder(true)
	barGraph.SetTitle("Bar Graph Demo")
	// display system metric usage
	barGraph.AddBar("cpu", 80, tcell.ColorBlue)
	barGraph.AddBar("memory", 20, tcell.ColorRed)
	barGraph.AddBar("swap", 40, tcell.ColorGreen)
	barGraph.SetMaxValue(100)
	/*
		update := func() {
			rand.Seed(time.Now().UnixNano())
			tick := time.NewTicker(1000 * time.Millisecond)
			for {
				select {
				case <-tick.C:
					rangeLower := 0
					rangeUpper := 100
					randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
					barGraph.SetBarValue("cpu", randomNum)
					randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
					barGraph.SetBarValue("memory", randomNum)
					randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
					barGraph.SetBarValue("swap", randomNum)
					app.Draw()
				}
			}
		}
		go update()
	*/
	if err := app.SetRoot(barGraph, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
