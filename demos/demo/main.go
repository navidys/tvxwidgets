// Demo code for the bar chart primitive.
package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// bar graph
	barGraph := tvxwidgets.NewBarChart()
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")
	barGraph.AddBar("cpu", 20, tcell.ColorBlue)
	barGraph.AddBar("mem", 60, tcell.ColorRed)
	barGraph.AddBar("swap", 80, tcell.ColorGreen)
	barGraph.AddBar("disk", 100, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	// activity mode gauge
	amGauge := tvxwidgets.NewActivityModeGauge()
	amGauge.SetTitle("activity mode gauge")
	amGauge.SetPgBgColor(tcell.ColorOrange)
	amGauge.SetRect(10, 4, 50, 3)
	amGauge.SetBorder(true)

	// percetage mode gauge
	pmGauge := tvxwidgets.NewPercentageModeGauge()
	pmGauge.SetTitle("percentage mode gauge")
	pmGauge.SetRect(10, 4, 50, 3)
	pmGauge.SetBorder(true)
	pmGauge.SetMaxValue(50)

	gaugeFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	gaugeFlex.AddItem(amGauge, 3, 0, false)
	gaugeFlex.AddItem(pmGauge, 3, 0, false)

	screenLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	screenLayout.AddItem(barGraph, 40, 0, false)
	screenLayout.AddItem(gaugeFlex, 40, 0, false)

	//screenLayout.SetRect(0, 0, 100, 15)

	update := func() {
		rand.Seed(time.Now().UnixNano())
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				// update gauge
				amGauge.Pulse()
				pmValue := pmGauge.GetValue()
				if pmValue > pmGauge.GetMaxValue() {
					pmValue = 0
					pmGauge.SetBackgroundColor(tcell.ColorOrange)
				} else {
					pmValue = pmValue + 1
				}
				pmGauge.SetValue(pmValue)

				// update bar graph
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

	if err := app.SetRoot(screenLayout, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
