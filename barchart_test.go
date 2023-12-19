package tvxwidgets_test

import (
	"github.com/gdamore/tcell/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"

	"github.com/navidys/tvxwidgets"
)

var _ = Describe("Barchart", Ordered, func() {
	var (
		app       *tview.Application
		headerBox *tview.Box
		barchart  *tvxwidgets.BarChart
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = tview.NewApplication()
		headerBox = tview.NewBox().SetBorder(true)
		barchart = tvxwidgets.NewBarChart()
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
			appLayout.AddItem(headerBox, 1, 0, true)
			appLayout.AddItem(barchart, 50, 0, true)
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
			Expect(barchart.HasFocus()).To(Equal(false))

			app.SetFocus(barchart)
			app.Draw()
			Expect(barchart.HasFocus()).To(Equal(true))
		})
	})

	Describe("GetRect", func() {
		It("primitivie size", func() {
			x, y, width, heigth := barchart.GetRect()
			Expect(x).To(Equal(0))
			Expect(y).To(Equal(1))
			Expect(width).To(Equal(80))
			Expect(heigth).To(Equal(50))
		})
	})

})
