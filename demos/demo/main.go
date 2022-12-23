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
	spinnerRow.SetBorder(true).SetTitle("spinners")

	for _, spinner := range spinners {
		spinnerRow.AddItem(spinner, 0, 1, false)
	}

	// bar graph
	barGraph := newBarChart()
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

	// plot (line charts)
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

	bmLineChart := newBrailleModeLineChart()
	bmLineChart.SetData(sinData)

	dmLineChart := newDotModeLineChart()

	sampleData1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sampleData2 := []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	dotChartData := [][]float64{sampleData1}
	dotChartData[0] = append(dotChartData[0], sampleData2...)
	dotChartData[0] = append(dotChartData[0], sampleData1[:5]...)
	dotChartData[0] = append(dotChartData[0], sampleData2[5:]...)
	dotChartData[0] = append(dotChartData[0], sampleData1[:7]...)
	dotChartData[0] = append(dotChartData[0], sampleData2[3:]...)

	dmLineChart.SetData(dotChartData)

	// sparkline
	iowaitSparkline := tvxwidgets.NewSparkline()
	iowaitSparkline.SetBorder(false)
	iowaitSparkline.SetDataTitle("Disk IO (iowait)")
	iowaitSparkline.SetDataTitleColor(tcell.ColorDarkOrange)
	iowaitSparkline.SetLineColor(tcell.ColorMediumPurple)

	systemSparkline := tvxwidgets.NewSparkline()
	systemSparkline.SetBorder(false)
	systemSparkline.SetDataTitle("Disk IO (system)")
	systemSparkline.SetDataTitleColor(tcell.ColorDarkOrange)
	systemSparkline.SetLineColor(tcell.ColorSteelBlue)

	iowaitData := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	systemData := []float64{0, 0, 1, 2, 9, 5, 3, 1, 2, 0, 6, 3, 2, 2, 6, 8, 5, 2, 1, 5, 8, 6, 1, 4, 1, 1, 4, 3, 6}

	ioSparkLineData := func() []float64 {
		for i := 0; i < 5; i++ {
			iowaitData = append(iowaitData, iowaitData...)
		}

		return iowaitData
	}()

	systemSparklineData := func() []float64 {
		for i := 0; i < 5; i++ {
			systemData = append(systemData, systemData...)
		}

		return systemData
	}()

	iowaitSparkline.SetData(ioSparkLineData)
	systemSparkline.SetData(systemSparklineData)

	sparklineGroupLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	sparklineGroupLayout.SetBorder(true)
	sparklineGroupLayout.SetTitle("sparkline")
	sparklineGroupLayout.AddItem(iowaitSparkline, 0, 1, false)
	sparklineGroupLayout.AddItem(tview.NewBox(), 1, 0, false)
	sparklineGroupLayout.AddItem(systemSparkline, 0, 1, false)

	// first row layout
	firstRowfirstCol := tview.NewFlex().SetDirection(tview.FlexRow)
	firstRowfirstCol.AddItem(barGraph, 0, 1, false)

	firstRowSecondCol := tview.NewFlex().SetDirection(tview.FlexRow)
	firstRowSecondCol.AddItem(amGauge, 0, 3, false)
	firstRowSecondCol.AddItem(pmGauge, 0, 3, false)
	firstRowSecondCol.AddItem(utilFlex, 0, 5, false)

	firstRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	firstRow.AddItem(firstRowfirstCol, 0, 1, false)
	firstRow.AddItem(firstRowSecondCol, 0, 1, false)

	// second row
	plotRowLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	plotRowLayout.AddItem(bmLineChart, 0, 1, false)
	plotRowLayout.AddItem(dmLineChart, 0, 1, false)

	screenLayout := tview.NewFlex().SetDirection(tview.FlexRow)
	screenLayout.AddItem(firstRow, 11, 0, false)
	screenLayout.AddItem(plotRowLayout, 15, 0, false)
	screenLayout.AddItem(sparklineGroupLayout, 6, 0, false)
	screenLayout.AddItem(spinnerRow, 3, 0, false)

	screenLayout.SetRect(0, 0, 100, 40)

	// upgrade datat functions
	moveDotChartData := func() {
		newData := append(dotChartData[0], dotChartData[0][0])
		dotChartData[0] = newData[1:]
	}

	moveDiskIOData := func() ([]float64, []float64) {

		newIOWaitData := ioSparkLineData[1:]
		newIOWaitData = append(newIOWaitData, ioSparkLineData[0])
		ioSparkLineData = newIOWaitData

		newSystemData := systemSparklineData[1:]
		newSystemData = append(newSystemData, systemSparklineData[0])
		systemSparklineData = newSystemData

		return newIOWaitData, newSystemData
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

	// update screen ticker
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

				randNum1 := float64(rand.Float64() * 100)
				randNum2 := float64(rand.Float64() * 100)
				randNum3 := float64(rand.Float64() * 100)
				randNum4 := float64(rand.Float64() * 100)

				barGraph.SetBarValue("eth0", int(randNum1))
				cpuGauge.SetValue(randNum1)
				barGraph.SetBarValue("eth1", int(randNum2))
				memGauge.SetValue(randNum2)
				barGraph.SetBarValue("eth2", int(randNum3))
				swapGauge.SetValue(randNum3)
				barGraph.SetBarValue("eth3", int(randNum4))

				// move line charts
				sinData = moveSinData(sinData)
				bmLineChart.SetData(sinData)

				moveDotChartData()
				dmLineChart.SetData(dotChartData)

				d1, d2 := moveDiskIOData()
				iowaitSparkline.SetData(d1)
				systemSparkline.SetData(d2)

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

func newDotModeLineChart() *tvxwidgets.Plot {
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

	return dmLineChart
}

func newBrailleModeLineChart() *tvxwidgets.Plot {
	bmLineChart := tvxwidgets.NewPlot()
	bmLineChart.SetBorder(true)
	bmLineChart.SetTitle("line chart (braille mode)")
	bmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorSteelBlue,
		tcell.ColorGreen,
	})
	bmLineChart.SetMarker(tvxwidgets.PlotMarkerBraille)

	return bmLineChart
}

func newBarChart() *tvxwidgets.BarChart {
	barGraph := tvxwidgets.NewBarChart()
	barGraph.SetBorder(true)
	barGraph.SetTitle("bar chart")
	barGraph.AddBar("eth0", 20, tcell.ColorBlue)
	barGraph.AddBar("eth1", 60, tcell.ColorRed)
	barGraph.AddBar("eth2", 80, tcell.ColorGreen)
	barGraph.AddBar("eth3", 100, tcell.ColorOrange)

	return barGraph
}
