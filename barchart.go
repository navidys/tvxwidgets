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
	xAxisStartY := height - 2
	barStartY := height - 3
	borderPadding := 0
	if c.hasBorder {
		borderPadding = 1
	}
	// draw graph y-axis
	for i := borderPadding; i+y < height; i++ {
		screen.SetContent(x+barChartYAxisLabelWidth, y+i, tview.Borders.Vertical, nil, style)
	}
	// draw graph x-axix
	for i := 0; i+x < width-borderPadding; i++ {
		screen.SetContent(x+borderPadding+i, xAxisStartY, tview.Borders.Horizontal, nil, style)
	}
	screen.SetContent(x+barChartYAxisLabelWidth, xAxisStartY, tview.Borders.Cross, nil, style)

	// set max value if not set
	c.initMaxValue()

	// draw bars
	startX := x + barChartYAxisLabelWidth + c.barGap
	labelY := height - 1
	valueMaxHeight := barStartY - borderPadding - 1
	for _, item := range c.bars {
		// set labels
		r := []rune(item.label)
		for j := 0; j < len(r); j++ {
			screen.SetContent(startX+j, labelY, r[j], nil, style)
		}
		// bar style
		bStyle := style.Foreground(item.color)
		barHeight := c.getHeight(valueMaxHeight, item.value)
		for k := 0; k < barHeight; k++ {
			for l := 0; l < c.barWidth; l++ {
				screen.SetContent(startX+l, barStartY-k, '\u2588', nil, bStyle)
			}

		}
		// bar value
		vSt := fmt.Sprintf("%d", item.value)
		vRune := []rune(vSt)
		for i := 0; i < len(vRune); i++ {
			screen.SetContent(startX+i, barStartY-barHeight, vRune[i], nil, bStyle)
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
