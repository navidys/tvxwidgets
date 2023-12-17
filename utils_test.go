package tvxwidgets_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"
)

var _ = Describe("Utils", func() {
	Describe("getColorName", func() {
		It("returns color name", func() {
			tests := []struct {
				color     tcell.Color
				colorName string
			}{
				{color: tcell.ColorWhite, colorName: "white"},
				{color: tcell.ColorBlack, colorName: "black"},
				{color: tcell.NewRGBColor(0, 1, 2), colorName: ""},
			}

			for _, test := range tests {
				Expect(tvxwidgets.GetColorName(test.color)).To(Equal(test.colorName))
			}
		})
	})

	Describe("getMessageWidth", func() {
		It("returns width size for dialogs based on messages", func() {
			tests := []struct {
				msg   string
				width int
			}{
				{msg: "test", width: 4},
				{msg: "test01\ntest001", width: 7},
				{msg: "", width: 0},
			}

			for _, test := range tests {
				Expect(tvxwidgets.GetMessageWidth(test.msg)).To(Equal(test.width))
			}
		})
	})

	Describe("getMaxFloat64From2dSlice", func() {
		It("returns max values in 2D float64 slices.", func() {
			tests := []struct {
				have  [][]float64
				wants float64
			}{
				{have: [][]float64{}, wants: 0},
				{have: [][]float64{
					{5, -1, 0, -10, 12},
					{15, -11, 0, -110, 22},
				}, wants: 22},
				{have: [][]float64{
					{-5, -1, -2, -10, -12},
					{-15, -11, -1, -110, -22},
				}, wants: -1},
			}

			for _, test := range tests {
				Expect(tvxwidgets.GetMaxFloat64From2dSlice(test.have)).To(Equal(test.wants))
			}
		})
	})

	Describe("getMaxFloat64FromSlice", func() {
		It("returns max values in float64 slices", func() {
			tests := []struct {
				have  []float64
				wants float64
			}{
				{have: []float64{}, wants: 0},
				{have: []float64{5, -1, 0, -10, 12}, wants: 12},
				{have: []float64{-10, -20, -9, -1}, wants: -1},
			}

			for _, test := range tests {
				Expect(tvxwidgets.GetMaxFloat64FromSlice(test.have)).To(Equal(test.wants))
			}
		})
	})

	Describe("absInt", func() {
		It("return absint", func() {
			tests := []struct {
				have  int
				wants int
			}{
				{have: 2, wants: 2},
				{have: -2, wants: 2},
				{have: 0, wants: 0},
			}

			for _, test := range tests {
				Expect(tvxwidgets.AbsInt(test.have)).To(Equal(test.wants))
			}
		})
	})

	Describe("drawLine", func() {
		It("draws horizontal or vertival line on screen", func() {
			screen := tcell.NewSimulationScreen("UTF-8")
			screenWidth := 70
			screenHeight := 30
			lineStartX := 0
			lineStartY := 0
			lineLenght := 20
			screen.SetSize(screenWidth, screenHeight)
			screen.Init()
			screen.Clear()

			// draw and test horizental line
			tvxwidgets.DrawLine(screen, lineStartX, lineStartY, lineLenght, 0, tcell.StyleDefault)
			screen.Show()

			cellRune, _, _, _ := screen.GetContent(lineStartX, lineStartY)
			Expect(cellRune).To(Equal(tview.BoxDrawingsLightTripleDashHorizontal))

			// draw and test vertical line
			screen.Clear()
			tvxwidgets.DrawLine(screen, lineStartX, lineStartY, lineLenght, 1, tcell.StyleDefault)
			screen.Show()

			cellRune, _, _, _ = screen.GetContent(lineStartX, lineStartY)
			Expect(cellRune).To(Equal(tview.BoxDrawingsLightTripleDashVertical))
		})
	})
})
