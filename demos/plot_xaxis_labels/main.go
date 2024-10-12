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

	// With these values, the curve will start with a value of 0 and reach a
	// high point of 2 at x = 3.14 (Pi) and then return to 0 at x = 6.28 (2*Pi).

	// Play around with these values to get a feel for how they affect the curve
	// and how you might adapt this code to plot other functions.

	period := 2 * math.Pi
	xAxisStrechFactor := 1.0
	yAxisStretchFactor := 1.0
	//xOffset := -(period)*1 - period/4
	xOffset := -period / 4
	//xOffset := 0.0
	yOffset := 1.0

	// These values influence which part of the curve is shown in
	// what "zoom level".

	// Note: There is no way to only zoom the visible range on the x-axis, so we
	// have to zoom the actual data values instead (see xAxisZoomFactor).
	xAxisZoomFactor := 10.0
	// TODO: needs custom min/max values for y-axis, coming soon TM in #68
	//yAxisZoomFactor := 1.0

	// Note: There is no way to only shift the visible range on the x-axis, so we
	// have to shift the actual data values instead (see xOffset).
	//xAxisShift := xOffset
	//yAxisShift := 0.0

	xFunc1 := func(i int) float64 {
		return (((float64(i)) / xAxisZoomFactor) * math.Pi) + xOffset
	}
	yFunc1 := func(x float64) float64 {
		return (math.Sin((x+xOffset)/xAxisStrechFactor) + yOffset) * yAxisStretchFactor
	}

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
	bmLineChart.SetXAxisLabelFunc(func(i int) string {
		label := fmt.Sprintf("%.1f", xFunc1(i))
		return label
	})
	// TODO: enable when #68 is done
	//bmLineChart.SetYAxisAutoScaleMin(false)
	//bmLineChart.SetYAxisAutoScaleMax(false)
	//bmLineChart.SetYRange(
	//	(-1+yOffset+yAxisShift)/yAxisZoomFactor,
	//	(1+yOffset+yAxisShift)/yAxisZoomFactor,
	//)
	bmLineChart.SetData(data)

	firstRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	firstRow.AddItem(bmLineChart, 0, 1, false)
	firstRow.SetRect(0, 0, 100, 15)

	layout := tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(firstRow, 0, 1, false)
	layout.SetRect(0, 0, 100, 30)

	rotateDataContinuously := func() {
		tick := time.NewTicker(100 * time.Millisecond)
		go func() {
			initialOffsetX := xOffset
			for {
				select {
				case <-tick.C:
					//xOffset = xOffset + 0.1
					if xOffset >= initialOffsetX+period {
						xOffset = initialOffsetX
					}
					data := computeDataArray()
					bmLineChart.SetData(data)

					app.Draw()
				}
			}
		}()
	}

	go rotateDataContinuously()

	if err := app.SetRoot(layout, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
