// Demo code for the bar chart primitive.
package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// spinners
	spinners := []*tvxwidgets.Spinner{

		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerDotsCircling),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerDotsUpDown),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerBounce),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerLine),

		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCircleQuarters),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerSquareCorners),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCircleHalves),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCorners),

		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerArrows),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerHamburger),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerStack),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerStar),

		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerGrowHorizontal),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerGrowVertical),
		tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerBoxBounce),
		tvxwidgets.NewSpinner().SetCustomStyle([]rune{'🕛', '🕐', '🕑', '🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚'}),
	}

	spinnerRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	spinnerRow.SetBorder(true).SetTitle("Spinners")

	for _, spinner := range spinners {
		spinnerRow.AddItem(spinner, 0, 1, false)
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

	sinData := func() [][]float64 {
		n := 220
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()

	bmLineChart := tvxwidgets.NewPlot()
	bmLineChart.SetBorder(true)
	bmLineChart.SetTitle("line chart (braille mode)")
	bmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorSteelBlue,
		tcell.ColorGreen,
	})
	bmLineChart.SetMarker(tvxwidgets.PlotMarkerBraille)
	bmLineChart.SetData(sinData)

	dmLineChart := tvxwidgets.NewPlot()
	dmLineChart.SetBorder(true)
	dmLineChart.SetTitle("line chart (dot mode)")
	dmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorDarkOrange,
	})
	dmLineChart.SetAxesLabelColor(tcell.ColorGold)
	dmLineChart.SetAxesColor(tcell.ColorGold)
	dmLineChart.SetMarker(tvxwidgets.PlotMarkerDot)
	dmLineChart.SetDotMarkerRune('\u25c9')

	sampleData1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sampleData2 := []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	dotChartData := [][]float64{sampleData1}
	dotChartData[0] = append(dotChartData[0], sampleData2...)
	dotChartData[0] = append(dotChartData[0], sampleData1[:5]...)
	dotChartData[0] = append(dotChartData[0], sampleData2[5:]...)
	dotChartData[0] = append(dotChartData[0], sampleData1[:7]...)
	dotChartData[0] = append(dotChartData[0], sampleData2[3:]...)

	dmLineChart.SetData(dotChartData)

	firstCol := tview.NewFlex().SetDirection(tview.FlexRow)
	firstCol.AddItem(barGraph, 11, 0, false)
	firstCol.AddItem(bmLineChart, 15, 0, false)
	firstCol.AddItem(spinnerRow, 3, 0, false)

	secondCol := tview.NewFlex().SetDirection(tview.FlexRow)
	secondCol.AddItem(amGauge, 3, 0, false)
	secondCol.AddItem(pmGauge, 3, 0, false)
	secondCol.AddItem(utilFlex, 5, 0, false)
	secondCol.AddItem(dmLineChart, 15, 0, false)

	screenLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	screenLayout.AddItem(firstCol, 50, 0, false)
	screenLayout.AddItem(secondCol, 50, 0, false)

	moveDotChartData := func() {
		newData := append(dotChartData[0], dotChartData[0][0])
		dotChartData[0] = newData[1:]
	}

	moveSinData := func(data [][]float64) [][]float64 {
		n := 220
		newData := make([][]float64, 2)
		newData[0] = make([]float64, n)
		newData[1] = make([]float64, n)

		for i := 0; i < n; i++ {
			if i+1 < len(data[0]) {
				newData[0][i] = data[0][i+1]
			}
			if i+1 < len(data[1]) {
				newData[1][i] = data[1][i+1]
			}
		}

		return newData
	}

	updateSpinner := func() {
		spinnerTick := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-spinnerTick.C:
				// update spinners
				for _, spinner := range spinners {
					spinner.Pulse()
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

				// move line charts
				sinData = moveSinData(sinData)
				bmLineChart.SetData(sinData)

				moveDotChartData()
				dmLineChart.SetData(dotChartData)

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
