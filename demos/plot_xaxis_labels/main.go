package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
	"math"
	"time"
)

func main() {

	app := tview.NewApplication()

	// >>> Data Function <<<
	// With these values, the curve will start with a value of 0 and reach a
	// high point of 2 at x = 3.14 (Pi) and then return to 0 at x = 6.28 (2*Pi).

	// Play around with these values to get a feel for how they affect the curve
	// and how you might adapt this code to plot other functions.

	period := 2 * math.Pi
	horizontalStretchFactor := 1.0
	verticalStretchFactor := 1.0
	xOffset := 0.0
	yOffset := 0.0

	// >>> Graph View/Camera Controls <<<
	// These values influence which part of the curve is shown in
	// what "zoom level".

	xAxisZoomFactor := 3.0
	yAxisZoomFactor := 1.0
	xAxisShift := 0.0
	yAxisShift := 0.0

	// xFunc1 defines the x values that should be used for each vertical "slot" in the graph.
	xFunc1 := func(i int) float64 {
		return (float64(i) / xAxisZoomFactor) + xAxisShift
	}
	// yFunc1 defines the y values that result from a given input value x (this is the actual function).
	yFunc1 := func(x float64) float64 {
		return (math.Sin((x+xOffset)/horizontalStretchFactor) + yOffset) * verticalStretchFactor
	}

	// xLabelFunc1 defines a label for each vertical "slot". Which labels are shown is determined automatically
	// based on the available space.
	xLabelFunc1 := func(i int) string {
		xVal := xFunc1(i)
		labelVal := xVal
		label := fmt.Sprintf("%.1f", labelVal)
		return label
	}

	// computeDataArray computes the y values for n vertical slots based on the definitions above.
	computeDataArray := func() [][]float64 {
		n := 150
		data := make([][]float64, 1)
		data[0] = make([]float64, n)
		for i := 0; i < n; i++ {
			xVal := xFunc1(i)
			yVal := yFunc1(xVal)
			data[0][i] = yVal
		}

		return data
	}

	data := computeDataArray()

	bmLineChart := tvxwidgets.NewPlot()
	bmLineChart.SetBorder(true)
	bmLineChart.SetTitle("line chart (braille mode)")
	bmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorSteelBlue,
		tcell.ColorGreen,
	})
	bmLineChart.SetMarker(tvxwidgets.PlotMarkerBraille)
	bmLineChart.SetXAxisLabelFunc(xLabelFunc1)
	bmLineChart.SetYAxisAutoScaleMin(false)
	bmLineChart.SetYAxisAutoScaleMax(false)
	bmLineChart.SetYRange(
		(-1+yOffset+yAxisShift)/yAxisZoomFactor,
		(1+yOffset+yAxisShift)/yAxisZoomFactor,
	)
	bmLineChart.SetData(data)

	firstRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	firstRow.AddItem(bmLineChart, 0, 1, false)
	firstRow.SetRect(0, 0, 100, 15)

	layout := tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(firstRow, 0, 1, false)
	layout.SetRect(0, 0, 100, 30)

	animate := true

	rotateDataContinuously := func() {
		tick := time.NewTicker(100 * time.Millisecond)
		go func() {
			initialxAxisShift := xAxisShift
			for {
				select {
				case <-tick.C:
					if !animate {
						continue
					}

					xAxisShift = xAxisShift + 0.1
					if xAxisShift >= initialxAxisShift+period*4 {
						xAxisShift = initialxAxisShift
					}
					data = computeDataArray()
					bmLineChart.SetData(data)

					app.Draw()
				}
			}
		}()
	}

	go rotateDataContinuously()

	if err := app.SetRoot(layout, false).EnableMouse(true).SetMouseCapture(func(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
		if action == tview.MouseLeftClick {
			animate = !animate
		}
		return event, action
	}).Run(); err != nil {
		panic(err)
	}
}
