package main

import (
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	sinData := func() [][]float64 {
		n := 220
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i+1)/5)
			// Avoid taking Cos(0) because it creates a high point of 2 that
			// will never be hit again and makes the graph look a little funny
			data[1][i] = 1 + math.Cos(float64(i+1)/5)
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
	bmLineChart.SetDrawXAxisLabel(false)

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

	dotModeChartData := [][]float64{sampleData1}
	dotModeChartData[0] = append(dotModeChartData[0], sampleData2...)
	dotModeChartData[0] = append(dotModeChartData[0], sampleData1[:5]...)
	dotModeChartData[0] = append(dotModeChartData[0], sampleData2[5:]...)
	dotModeChartData[0] = append(dotModeChartData[0], sampleData1[:7]...)
	dotModeChartData[0] = append(dotModeChartData[0], sampleData2[3:]...)
	dmLineChart.SetData(dotModeChartData)

	scatterPlotData := make([][]float64, 2)
	scatterPlotData[0] = []float64{1, 2, 3, 4, 5}
	scatterPlotData[1] = sinData[1][4:]
	dmScatterPlot := tvxwidgets.NewPlot()

	dmScatterPlot.SetBorder(true)
	dmScatterPlot.SetTitle("scatter plot (dot mode)")
	dmScatterPlot.SetLineColor([]tcell.Color{
		tcell.ColorMediumSlateBlue,
		tcell.ColorLightSkyBlue,
	})
	dmScatterPlot.SetPlotType(tvxwidgets.PlotTypeScatter)
	dmScatterPlot.SetMarker(tvxwidgets.PlotMarkerDot)
	dmScatterPlot.SetData(scatterPlotData)
	dmScatterPlot.SetDrawYAxisLabel(false)

	bmScatterPlot := tvxwidgets.NewPlot()
	bmScatterPlot.SetBorder(true)
	bmScatterPlot.SetTitle("scatter plot (braille mode)")
	bmScatterPlot.SetLineColor([]tcell.Color{
		tcell.ColorGold,
		tcell.ColorLightSkyBlue,
	})
	bmScatterPlot.SetPlotType(tvxwidgets.PlotTypeScatter)
	bmScatterPlot.SetMarker(tvxwidgets.PlotMarkerBraille)
	bmScatterPlot.SetData(scatterPlotData)

	firstRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	firstRow.AddItem(dmLineChart, 0, 1, false)
	firstRow.AddItem(bmLineChart, 0, 1, false)
	firstRow.SetRect(0, 0, 100, 15)

	secondRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	secondRow.AddItem(dmScatterPlot, 0, 1, false)
	secondRow.AddItem(bmScatterPlot, 0, 1, false)
	secondRow.SetRect(0, 0, 100, 15)

	layout := tview.NewFlex().SetDirection(tview.FlexRow)
	layout.AddItem(firstRow, 0, 1, false)
	layout.AddItem(secondRow, 0, 1, false)
	layout.SetRect(0, 0, 100, 30)

	if err := app.SetRoot(layout, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
