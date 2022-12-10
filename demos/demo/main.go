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

	// spinners
	spinners := [][]*tvxwidgets.Spinner{
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerDotsCircling),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerDotsUpDown),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerBounce),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerLine),
		},
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCircleQuarters),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerSquareCorners),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCircleHalves),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCorners),
		},
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerArrows),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerHamburger),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerStack),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerStar),
		},
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerGrowHorizontal),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerGrowVertical),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerBoxBounce),
			tvxwidgets.NewSpinner().SetCustomStyle([]rune{'🕛', '🕐', '🕑', '🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚'}),
		},
	}

	spinnerGrid := tview.NewGrid()
	spinnerGrid.SetBorder(true).SetTitle("Spinners")

	for rowIdx, row := range spinners {
		for colIdx, spinner := range row {
			spinnerGrid.AddItem(spinner, rowIdx, colIdx, 1, 1, 1, 1, false)
		}
	}

	// bar graph
	barGraph := tvxwidgets.NewBarChart()
	barGraph.SetBorder(true)
	barGraph.SetTitle("bar chart")
	barGraph.AddBar("eth0", 20, tcell.ColorBlue)
	barGraph.AddBar("eth1", 60, tcell.ColorRed)
	barGraph.AddBar("eth2", 80, tcell.ColorGreen)
	barGraph.AddBar("eth3", 100, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	// activity mode gauge
	amGauge := tvxwidgets.NewActivityModeGauge()
	amGauge.SetTitle("activity mode gauge")
	amGauge.SetPgBgColor(tcell.ColorOrange)
	amGauge.SetBorder(true)

	// percentage mode gauge
	pmGauge := tvxwidgets.NewPercentageModeGauge()
	pmGauge.SetTitle("percentage mode gauge")
	pmGauge.SetBorder(true)
	pmGauge.SetMaxValue(50)

	// cpu usage gauge
	cpuGauge := tvxwidgets.NewUtilModeGauge()
	cpuGauge.SetLabel("cpu usage:   ")
	cpuGauge.SetLabelColor(tcell.ColorLightSkyBlue)
	cpuGauge.SetBorder(false)
	// memory usage gauge
	memGauge := tvxwidgets.NewUtilModeGauge()
	memGauge.SetLabel("memory usage:")
	memGauge.SetLabelColor(tcell.ColorLightSkyBlue)
	memGauge.SetBorder(false)
	// swap usage gauge
	swapGauge := tvxwidgets.NewUtilModeGauge()
	swapGauge.SetLabel("swap usage:  ")
	swapGauge.SetLabelColor(tcell.ColorLightSkyBlue)
	swapGauge.SetBorder(false)

	// utilisation flex
	utilFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	utilFlex.AddItem(cpuGauge, 1, 0, false)
	utilFlex.AddItem(memGauge, 1, 0, false)
	utilFlex.AddItem(swapGauge, 1, 0, false)
	utilFlex.SetTitle("utilisation mode gauge")
	utilFlex.SetBorder(true)

	firstCol := tview.NewFlex().SetDirection(tview.FlexRow)
	firstCol.AddItem(barGraph, 11, 0, false)

	secondCol := tview.NewFlex().SetDirection(tview.FlexRow)
	secondCol.AddItem(amGauge, 3, 0, false)
	secondCol.AddItem(pmGauge, 3, 0, false)
	secondCol.AddItem(utilFlex, 5, 0, false)

	screenLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	screenLayout.AddItem(firstCol, 50, 0, false)
	screenLayout.AddItem(secondCol, 50, 0, false)
	screenLayout.AddItem(spinnerGrid, 15, 0, false)

	updateSpinner := func() {
		spinnerTick := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-spinnerTick.C:
				// update spinners
				for _, row := range spinners {
					for _, spinner := range row {
						spinner.Pulse()
					}
				}
				// update gauge
				amGauge.Pulse()

				app.Draw()
			}
		}
	}

	update := func() {
		value := 0
		maxValue := pmGauge.GetMaxValue()
		rand.Seed(time.Now().UnixNano())
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:

				if value > maxValue {
					value = 0
				} else {
					value = value + 1
				}
				pmGauge.SetValue(value)

				// update bar graph
				rangeLower := 0
				rangeUpper := 100
				randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("eth0", randomNum)
				cpuGauge.SetValue(float64(randomNum))
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("eth1", randomNum)
				memGauge.SetValue(float64(randomNum))
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("eth2", randomNum)
				swapGauge.SetValue(float64(randomNum))
				randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
				barGraph.SetBarValue("eth3", randomNum)
				app.Draw()
			}
		}
	}

	go updateSpinner()
	go update()

	if err := app.SetRoot(screenLayout, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
