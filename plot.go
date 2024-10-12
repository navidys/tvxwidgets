package tvxwidgets

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Marker represents plot drawing marker (braille or dot).
type Marker uint

const (
	// plot marker.
	PlotMarkerBraille Marker = iota
	PlotMarkerDot
)

// PlotYAxisLabelDataType represents plot y axis type (integer or float).
type PlotYAxisLabelDataType uint

const (
	PlotYAxisLabelDataInt PlotYAxisLabelDataType = iota
	PlotYAxisLabelDataFloat
)

// PlotType represents plot type (line chart or scatter).
type PlotType uint

const (
	PlotTypeLineChart PlotType = iota
	PlotTypeScatter
)

const (
	plotHorizontalScale   = 1
	plotXAxisLabelsHeight = 1
	plotXAxisLabelsGap    = 2
	plotYAxisLabelsGap    = 1
)

type brailleCell struct {
	cRune rune
	color tcell.Color
}

// Plot represents a plot primitive used for different charts.
type Plot struct {
	*tview.Box
	data [][]float64
	// maxVal is the maximum y-axis (vertical) value found in any of the lines in the data set.
	maxVal float64
	// minVal is the minimum y-axis (vertical) value found in any of the lines in the data set.
	minVal             float64
	marker             Marker
	ptype              PlotType
	dotMarkerRune      rune
	lineColors         []tcell.Color
	axesColor          tcell.Color
	axesLabelColor     tcell.Color
	drawAxes           bool
	drawXAxisLabel     bool
	xAxisLabelFunc     func(int) string
	drawYAxisLabel     bool
	yAxisLabelDataType PlotYAxisLabelDataType
	yAxisAutoScaleMin  bool
	yAxisAutoScaleMax  bool
	brailleCellMap     map[image.Point]brailleCell
	mu                 sync.Mutex
}

// NewPlot returns a plot widget.
func NewPlot() *Plot {
	return &Plot{
		Box:                tview.NewBox(),
		marker:             PlotMarkerDot,
		ptype:              PlotTypeLineChart,
		dotMarkerRune:      dotRune,
		axesColor:          tcell.ColorDimGray,
		axesLabelColor:     tcell.ColorDimGray,
		drawAxes:           true,
		drawXAxisLabel:     true,
		xAxisLabelFunc:     func(i int) string { return strconv.Itoa(i) },
		drawYAxisLabel:     true,
		yAxisLabelDataType: PlotYAxisLabelDataFloat,
		yAxisAutoScaleMin:  false,
		yAxisAutoScaleMax:  true,
		lineColors: []tcell.Color{
			tcell.ColorSteelBlue,
		},
	}
}

// Draw draws this primitive onto the screen.
func (plot *Plot) Draw(screen tcell.Screen) {
	plot.Box.DrawForSubclass(screen, plot)

	switch plot.marker {
	case PlotMarkerDot:
		plot.drawDotMarkerToScreen(screen)
	case PlotMarkerBraille:
		plot.drawBrailleMarkerToScreen(screen)
	}

	plot.drawAxesToScreen(screen)
}

// SetRect sets rect for this primitive.
func (plot *Plot) SetRect(x, y, width, height int) {
	plot.Box.SetRect(x, y, width, height)
}

// SetLineColor sets chart line color.
func (plot *Plot) SetLineColor(color []tcell.Color) {
	plot.lineColors = color
}

// SetYAxisLabelDataType sets Y axis label data type (integer or float).
func (plot *Plot) SetYAxisLabelDataType(dataType PlotYAxisLabelDataType) {
	plot.yAxisLabelDataType = dataType
}

// SetYAxisAutoScaleMin enables YAxis min value autoscale.
func (plot *Plot) SetYAxisAutoScaleMin(autoScale bool) {
	plot.yAxisAutoScaleMin = autoScale
}

// SetYAxisAutoScaleMax enables YAxix max value autoscale.
func (plot *Plot) SetYAxisAutoScaleMax(autoScale bool) {
	plot.yAxisAutoScaleMax = autoScale
}

// SetAxesColor sets axes x and y lines color.
func (plot *Plot) SetAxesColor(color tcell.Color) {
	plot.axesColor = color
}

// SetAxesLabelColor sets axes x and y label color.
func (plot *Plot) SetAxesLabelColor(color tcell.Color) {
	plot.axesLabelColor = color
}

// SetDrawAxes set true in order to draw axes to screen.
func (plot *Plot) SetDrawAxes(draw bool) {
	plot.drawAxes = draw
}

// SetDrawXAxisLabel set true in order to draw x axis label to screen.
func (plot *Plot) SetDrawXAxisLabel(draw bool) {
	plot.drawXAxisLabel = draw
}

// SetXAxisLabelFunc sets x axis label function.
func (plot *Plot) SetXAxisLabelFunc(f func(int) string) {
	plot.xAxisLabelFunc = f
}

// SetDrawYAxisLabel set true in order to draw y axis label to screen.
func (plot *Plot) SetDrawYAxisLabel(draw bool) {
	plot.drawYAxisLabel = draw
}

// SetMarker sets marker type braille or dot mode.
func (plot *Plot) SetMarker(marker Marker) {
	plot.marker = marker
}

// SetPlotType sets plot type (linechart or scatter).
func (plot *Plot) SetPlotType(ptype PlotType) {
	plot.ptype = ptype
}

// SetData sets plot data.
func (plot *Plot) SetData(data [][]float64) {
	plot.mu.Lock()
	defer plot.mu.Unlock()

	plot.brailleCellMap = make(map[image.Point]brailleCell)
	plot.data = data

	if plot.yAxisAutoScaleMax {
		plot.maxVal = getMaxFloat64From2dSlice(data)
	}

	if plot.yAxisAutoScaleMin {
		plot.minVal = getMinFloat64From2dSlice(data)
	}
}

func (plot *Plot) SetMaxVal(maxVal float64) {
	plot.maxVal = maxVal
}

func (plot *Plot) SetMinVal(minVal float64) {
	plot.minVal = minVal
}

func (plot *Plot) SetYRange(minVal float64, maxVal float64) {
	plot.minVal = minVal
	plot.maxVal = maxVal
}

// SetDotMarkerRune sets dot marker rune.
func (plot *Plot) SetDotMarkerRune(r rune) {
	plot.dotMarkerRune = r
}

// Figure out the text width necessary to display the largest data value.
func (plot *Plot) getYAxisLabelsWidth() int {
	return len(fmt.Sprintf("%.2f", plot.maxVal))
}

// GetPlotRect returns the rect for the inner part of the plot, ie not including axes.
func (plot *Plot) GetPlotRect() (int, int, int, int) {
	x, y, width, height := plot.Box.GetInnerRect()
	plotYAxisLabelsWidth := plot.getYAxisLabelsWidth()

	if plot.drawAxes {
		x = x + plotYAxisLabelsWidth + 1
		width = width - plotYAxisLabelsWidth - 1
		height = height - plotXAxisLabelsHeight - 1
	} else {
		x++
		width--
	}

	return x, y, width, height
}

func (plot *Plot) getData() [][]float64 {
	plot.mu.Lock()
	data := plot.data
	plot.mu.Unlock()

	return data
}

func (plot *Plot) drawAxesToScreen(screen tcell.Screen) {
	if !plot.drawAxes {
		return
	}

	x, y, width, height := plot.Box.GetInnerRect()
	plotYAxisLabelsWidth := plot.getYAxisLabelsWidth()

	axesStyle := tcell.StyleDefault.Background(plot.GetBackgroundColor()).Foreground(plot.axesColor)

	// draw Y axis line
	drawLine(screen,
		x+plotYAxisLabelsWidth,
		y,
		height-plotXAxisLabelsHeight-1,
		verticalLine, axesStyle)

	// draw X axis line
	drawLine(screen,
		x+plotYAxisLabelsWidth+1,
		y+height-plotXAxisLabelsHeight-1,
		width-plotYAxisLabelsWidth-1,
		horizontalLine, axesStyle)

	tview.PrintJoinedSemigraphics(screen,
		x+plotYAxisLabelsWidth,
		y+height-plotXAxisLabelsHeight-1,
		tview.BoxDrawingsLightUpAndRight, axesStyle)

	if plot.drawXAxisLabel {
		plot.drawXAxisLabelToScreen(screen, plotYAxisLabelsWidth, x, y, width, height)
	}

	if plot.drawYAxisLabel {
		plot.drawYAxisLabelToScreen(screen, plotYAxisLabelsWidth, x, y, height)
	}
}

func (plot *Plot) drawXAxisLabelToScreen(
	screen tcell.Screen, plotYAxisLabelsWidth int, x int, y int, width int, height int,
) {
	for labelX := x + plotYAxisLabelsWidth; labelX < x+width-1; {
		labelIndex := (labelX-(x+plotYAxisLabelsWidth)-1)/(plotHorizontalScale) + 1
		label := plot.xAxisLabelFunc(labelIndex)

		tview.Print(screen, label, labelX, y+height-plotXAxisLabelsHeight, width, tview.AlignLeft, plot.axesLabelColor)

		labelX += (len(label) + plotXAxisLabelsGap) * plotHorizontalScale
	}
}

func (plot *Plot) drawYAxisLabelToScreen(screen tcell.Screen, plotYAxisLabelsWidth int, x int, y int, height int) {
	verticalOffset := plot.minVal
	verticalScale := (plot.maxVal - plot.minVal) / float64(height-plotXAxisLabelsHeight-1)
	previousLabel := ""

	for i := 0; i*(plotYAxisLabelsGap+1) < height-1; i++ {
		var label string
		if plot.yAxisLabelDataType == PlotYAxisLabelDataFloat {
			label = fmt.Sprintf("%.2f", float64(i)*verticalScale*(plotYAxisLabelsGap+1)+verticalOffset)
		} else {
			label = strconv.Itoa(int(float64(i)*verticalScale*(plotYAxisLabelsGap+1) + verticalOffset))
		}

		// Prevent same label being shown twice.
		// Mainly relevant for integer labels with small data sets (in value)
		if label == previousLabel {
			continue
		}

		previousLabel = label

		tview.Print(screen,
			label,
			x,
			y+height-(i*(plotYAxisLabelsGap+1))-2, //nolint:gomnd
			plotYAxisLabelsWidth,
			tview.AlignLeft, plot.axesLabelColor)
	}
}

//nolint:cyclop,gocognit
func (plot *Plot) drawDotMarkerToScreen(screen tcell.Screen) {
	x, y, width, height := plot.GetPlotRect()
	chartData := plot.getData()
	verticalOffset := -plot.minVal

	switch plot.ptype {
	case PlotTypeLineChart:
		for i, line := range chartData {
			style := tcell.StyleDefault.Background(plot.GetBackgroundColor()).Foreground(plot.lineColors[i])

			for j := 0; j < len(line) && j*plotHorizontalScale < width; j++ {
				val := line[j]
				if math.IsNaN(val) {
					continue
				}

				lheight := int(((val + verticalOffset) / plot.maxVal) * float64(height-1))
				if lheight > height {
					continue
				}

				if (x+(j*plotHorizontalScale) < x+width) && (y+height-1-lheight < y+height) {
					tview.PrintJoinedSemigraphics(screen, x+(j*plotHorizontalScale), y+height-1-lheight, plot.dotMarkerRune, style)
				}
			}
		}

	case PlotTypeScatter:
		for i, line := range chartData {
			style := tcell.StyleDefault.Background(plot.GetBackgroundColor()).Foreground(plot.lineColors[i])

			for j, val := range line {
				if math.IsNaN(val) {
					continue
				}

				lheight := int(((val + verticalOffset) / plot.maxVal) * float64(height-1))
				if lheight > height {
					continue
				}

				if (x+(j*plotHorizontalScale) < x+width) && (y+height-1-lheight < y+height) {
					tview.PrintJoinedSemigraphics(screen, x+(j*plotHorizontalScale), y+height-1-lheight, plot.dotMarkerRune, style)
				}
			}
		}
	}
}

func (plot *Plot) drawBrailleMarkerToScreen(screen tcell.Screen) {
	x, y, width, height := plot.GetPlotRect()

	plot.calcBrailleLines()

	// print to screen
	for point, cell := range plot.getBrailleCells() {
		style := tcell.StyleDefault.Background(plot.GetBackgroundColor()).Foreground(cell.color)
		if point.X < x+width && point.Y < y+height {
			tview.PrintJoinedSemigraphics(screen, point.X, point.Y, cell.cRune, style)
		}
	}
}

func calcDataPointHeight(val, maxVal, minVal float64, height int) int {
	return int(((val - minVal) / (maxVal - minVal)) * float64(height-1))
}

func calcDataPointHeightIfInBounds(val float64, maxVal float64, minVal float64, height int) (int, bool) {
	if math.IsNaN(val) {
		return 0, false
	}

	result := calcDataPointHeight(val, maxVal, minVal, height)
	if (val > maxVal) || (val < minVal) || (result > height) {
		return result, false
	}

	return result, true
}

func (plot *Plot) calcBrailleLines() {
	x, y, _, height := plot.GetPlotRect()
	chartData := plot.getData()

	for i, line := range chartData {
		if len(line) <= 1 {
			continue
		}

		previousHeight := 0
		lastValWasOk := false

		for j, val := range line {
			lheight, currentValIsOk := calcDataPointHeightIfInBounds(val, plot.maxVal, plot.minVal, height)

			if !lastValWasOk && !currentValIsOk {
				// nothing valid to draw, skip to next data point
				continue
			}

			if !lastValWasOk { //nolint:gocritic
				// current data point is single valid data point, draw it individually
				plot.setBraillePoint(
					calcBraillePoint(x, j+1, y, height, lheight),
					plot.lineColors[i],
				)
			} else if !currentValIsOk {
				// last data point was single valid data point, draw it individually
				plot.setBraillePoint(
					calcBraillePoint(x, j, y, height, previousHeight),
					plot.lineColors[i],
				)
			} else {
				// we have two valid data points, draw a line between them
				plot.setBrailleLine(
					calcBraillePoint(x, j, y, height, previousHeight),
					calcBraillePoint(x, j+1, y, height, lheight),
					plot.lineColors[i],
				)
			}

			lastValWasOk = currentValIsOk
			previousHeight = lheight
		}
	}
}

func calcBraillePoint(x, j, y, maxY, height int) image.Point {
	return image.Pt(
		(x+(j*plotHorizontalScale))*2, //nolint:gomnd
		(y+maxY-height-1)*4,           //nolint:gomnd
	)
}

func (plot *Plot) setBraillePoint(p image.Point, color tcell.Color) {
	if p.X < 0 || p.Y < 0 {
		return
	}

	point := image.Pt(p.X/2, p.Y/4) //nolint:gomnd
	plot.brailleCellMap[point] = brailleCell{
		plot.brailleCellMap[point].cRune | brailleRune[p.Y%4][p.X%2],
		color,
	}
}

func (plot *Plot) setBrailleLine(p0, p1 image.Point, color tcell.Color) {
	for _, p := range plot.brailleLine(p0, p1) {
		plot.setBraillePoint(p, color)
	}
}

func (plot *Plot) getBrailleCells() map[image.Point]brailleCell {
	cellMap := make(map[image.Point]brailleCell)
	for point, cvCell := range plot.brailleCellMap {
		cellMap[point] = brailleCell{cvCell.cRune + brailleOffsetRune, cvCell.color}
	}

	return cellMap
}

func (plot *Plot) brailleLine(p0, p1 image.Point) []image.Point {
	points := []image.Point{}
	leftPoint, rightPoint := p0, p1

	if leftPoint.X > rightPoint.X {
		leftPoint, rightPoint = rightPoint, leftPoint
	}

	xDistance := absInt(leftPoint.X - rightPoint.X)
	yDistance := absInt(leftPoint.Y - rightPoint.Y)
	slope := float64(yDistance) / float64(xDistance)
	slopeSign := 1

	if rightPoint.Y < leftPoint.Y {
		slopeSign = -1
	}

	targetYCoordinate := float64(leftPoint.Y)
	currentYCoordinate := leftPoint.Y

	for i := leftPoint.X; i < rightPoint.X; i++ {
		points = append(points, image.Pt(i, currentYCoordinate))
		targetYCoordinate += (slope * float64(slopeSign))

		for currentYCoordinate != int(targetYCoordinate) {
			points = append(points, image.Pt(i, currentYCoordinate))

			currentYCoordinate += slopeSign
		}
	}

	return points
}
