package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
	"math"
)

func main() {

	app := tview.NewApplication()

	xAxisZoomFactor := 10.0
	yAxisZoomFactor := 1.0
	xOffset := 0.0
	yOffset := 1.0
	xFunc1 := func(i int) float64 {
		return (float64(i) / xAxisZoomFactor) * math.Pi
	}
	yFunc1 := func(x float64) float64 {
		return (math.Sin(x+xOffset) + yOffset) * yAxisZoomFactor
	}

	data := func() [][]float64 {
		n := 170
		data := make([][]float64, 1)
		data[0] = make([]float64, n)
		xToYMap := make(map[float64]float64)
		mappedValues := []string{}
		for i := 0; i < n; i++ {
			xVal := xFunc1(i)
			yVal := yFunc1(xVal)

			xToYMap[xVal] = yVal

			mappedValues = append(mappedValues, fmt.Sprintf("%.2f -> %.2f", xVal, yVal))
			data[0][i] = yVal
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
	bmLineChart.SetXAxisLabelFunc(func(i int) string {
		label := fmt.Sprintf("%.1f", xFunc1(i))
		return label
	})
	bmLineChart.SetData(data)

	firstRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	firstRow.AddItem(bmLineChart, 0, 1, false)
	firstRow.SetRect(0, 0, 100, 15)

	layout := tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(firstRow, 0, 1, false)
	layout.SetRect(0, 0, 100, 30)

	if err := app.SetRoot(layout, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
