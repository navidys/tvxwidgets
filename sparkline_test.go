package tvxwidgets_test

import (
	"github.com/gdamore/tcell/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"

	"github.com/navidys/tvxwidgets"
)

var _ = Describe("Sparkline", Ordered, func() {
	var (
		app       *tview.Application
		headerBox *tview.Box
		sparkline *tvxwidgets.Sparkline
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = tview.NewApplication()
		headerBox = tview.NewBox().SetBorder(true)
		sparkline = tvxwidgets.NewSparkline()
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
			appLayout.AddItem(headerBox, 1, 0, true)
			appLayout.AddItem(sparkline, 50, 0, true)
			err := app.SetScreen(screen).SetRoot(appLayout, true).Run()
			if err != nil {
				panic(err)
			}
		}()
	})

	AfterAll(func() {
		app.Stop()
	})

	Describe("Focus", func() {
		It("checks primitivie focus", func() {
			app.SetFocus(headerBox)
			app.Draw()
			Expect(sparkline.HasFocus()).To(Equal(false))

			app.SetFocus(sparkline)
			app.Draw()
			Expect(sparkline.HasFocus()).To(Equal(true))
		})
	})

	Describe("GetRect", func() {
		It("primitivie size", func() {
			x, y, width, heigth := sparkline.GetRect()
			Expect(x).To(Equal(0))
			Expect(y).To(Equal(1))
			Expect(width).To(Equal(80))
			Expect(heigth).To(Equal(50))
		})
	})

	Describe("DataTitle and Color", func() {
		It("checks data title text and color", func() {
			tests := []struct {
				title string
				color tcell.Color
			}{
				{title: "test01", color: tcell.ColorDarkOrange},
				{title: "abc123", color: tcell.ColorBlue},
			}

			for _, test := range tests {
				sparkline.SetDataTitle(test.title)
				sparkline.SetDataTitleColor(test.color)
				app.Draw()

				for x := 0; x < len(test.title); x++ {
					prune, _, style, _ := screen.GetContent(x, 1)
					fg, _, _ := style.Decompose()

					Expect(fg).To(Equal(test.color))
					Expect(string(prune)).To(Equal(string(test.title[x])))
				}
			}
		})
	})
})
