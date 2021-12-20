package tgraph

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	barChartYAxisLabelWidth = 2
)

// BarChartItem represents a single bar in bar chart
type BarChartItem struct {
	label string
	value int
	color tcell.Color
}

// BarChart represents bar chart primitive.
type BarChart struct {
	*tview.Box
	// bar items
	bars []BarChartItem
	// maximum value of bars
	maxVal int
	// barGap gap between two bars
	barGap int
	// barWidth width of bars
	barWidth int
	//hasBorder true if primitive has border
	hasBorder bool
}

// NewBarChart returns a new bar chart primitive.
func NewBarChart() *BarChart {
	chart := &BarChart{
		Box:      tview.NewBox(),
		barGap:   2,
		barWidth: 3,
	}
	return chart
}

// Focus is called when this primitive receives focus
func (c *BarChart) Focus(delegate func(p tview.Primitive)) {
	delegate(c.Box)
}

// HasFocus returns whether or not this primitive has focus.
func (c *BarChart) HasFocus() bool {
	return c.Box.HasFocus()
}

// Draw draws this primitive onto the screen.
func (c *BarChart) Draw(screen tcell.Screen) {
	style := tcell.StyleDefault
	style = style.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor)
	c.Box.DrawForSubclass(screen, c)
	x, y, width, height := c.Box.GetInnerRect()
	maxValY := y + 1
	xAxisStartY := height - 2
	barStartY := height - 3
	borderPadding := 0
	if c.hasBorder {
		borderPadding = 1
	}
	// set max value if not set
	c.initMaxValue()
	maxValueSr := fmt.Sprintf("%d", c.maxVal)
	maxValLenght := len(maxValueSr) + 1
	if maxValLenght < barChartYAxisLabelWidth {
		maxValLenght = barChartYAxisLabelWidth
	}
	// draw graph y-axis
	for i := borderPadding; i+y < height; i++ {
		tview.PrintJoinedSemigraphics(screen, x+maxValLenght, y+i, tview.Borders.Vertical, style)
	}
	// draw graph x-axix
	for i := maxValLenght; i+x < width-borderPadding; i++ {
		tview.PrintJoinedSemigraphics(screen, x+i, xAxisStartY, tview.Borders.Horizontal, style)
	}
	tview.PrintJoinedSemigraphics(screen, x+maxValLenght, xAxisStartY, tview.BoxDrawingsLightVerticalAndRight, style)
	tview.PrintJoinedSemigraphics(screen, x+maxValLenght-1, xAxisStartY, '0', style)

	mxValRune := []rune(maxValueSr)
	for i := 0; i < len(mxValRune); i++ {
		tview.PrintJoinedSemigraphics(screen, x+borderPadding+i, maxValY, mxValRune[i], style)
	}

	// draw bars
	startX := x + maxValLenght + c.barGap
	labelY := height - 1
	valueMaxHeight := barStartY - borderPadding - 1
	for _, item := range c.bars {
		if startX > width {
			return
		}
		// set labels
		r := []rune(item.label)
		for j := 0; j < len(r); j++ {
			tview.PrintJoinedSemigraphics(screen, startX+j, labelY, r[j], style)
		}
		// bar style
		bStyle := style.Foreground(item.color)
		barHeight := c.getHeight(valueMaxHeight, item.value)
		for k := 0; k < barHeight; k++ {
			for l := 0; l < c.barWidth; l++ {
				tview.PrintJoinedSemigraphics(screen, startX+l, barStartY-k, '\u2588', bStyle)
			}

		}
		// bar value
		vSt := fmt.Sprintf("%d", item.value)
		vRune := []rune(vSt)
		for i := 0; i < len(vRune); i++ {
			tview.PrintJoinedSemigraphics(screen, startX+i, barStartY-barHeight, vRune[i], bStyle)
		}

		// calculate next startX for next bar
		rWidth := len(r)
		if rWidth < c.barWidth {
			rWidth = c.barWidth
		}

		startX = startX + c.barGap + rWidth
	}

}

// SetBorder sets border for this primitive
func (c *BarChart) SetBorder(status bool) {
	c.hasBorder = status
	c.Box.SetBorder(status)
}

func (c *BarChart) GetRect() (int, int, int, int) {
	return c.Box.GetRect()
}

// SetRect sets rect for this primitive.
func (c *BarChart) SetRect(x, y, width, height int) {
	c.Box.SetRect(x, y, width, height)
}

// InputHandler returns input handler function for this primitive
func (c *BarChart) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return c.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {

	})
}

// SetMaxValue sets maximum value of bars
func (c *BarChart) SetMaxValue(max int) {
	c.maxVal = max
}

// AddBar adds new bar item to the bar chart primitive.
func (c *BarChart) AddBar(label string, value int, color tcell.Color) {
	c.bars = append(c.bars, BarChartItem{
		label: label,
		value: value,
		color: color,
	})
}

// SetBarValue sets bar values
func (c *BarChart) SetBarValue(name string, value int) {
	for i := 0; i < len(c.bars); i++ {
		if c.bars[i].label == name {
			c.bars[i].value = value
		}
	}
}

func (c *BarChart) getHeight(maxHeight int, value int) int {
	height := 0

	if value > c.maxVal {
		return maxHeight
	}
	height = (value * maxHeight) / c.maxVal
	return height
}

func (c *BarChart) initMaxValue() {
	// set max value if not set
	if c.maxVal == 0 {
		for _, b := range c.bars {
			if b.value > c.maxVal {
				c.maxVal = b.value
			}
		}
	}
}
