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
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	barGraph := tgraph.NewBarChart()
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")
	// display system metric usage
	barGraph.AddBar("cpu", 80, tcell.ColorBlue)
	barGraph.AddBar("mem", 20, tcell.ColorRed)
	barGraph.AddBar("swap", 40, tcell.ColorGreen)
	barGraph.AddBar("disk", 40, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	bar2 := tgraph.NewBarChart()
	bar2.SetBorder(true)
	bar2.SetTitle("System Resource Usage")
	// display system metric usage
	bar2.AddBar("cpu", 80, tcell.ColorBlue)
	bar2.AddBar("mem", 20, tcell.ColorRed)
	bar2.AddBar("swap", 40, tcell.ColorGreen)
	bar2.AddBar("disk", 40, tcell.ColorOrange)
	bar2.SetMaxValue(80)

	flex.AddItem(barGraph, 20, 1, false)

	flex.SetRect(0, 0, 50, 50)

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
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("disk", randomNum)
				app.Draw()
			}
		}
	}
	go update()

	if err := app.SetRoot(flex, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
