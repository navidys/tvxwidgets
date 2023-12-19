package tvxwidgets_test

import (
	"github.com/gdamore/tcell/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"

	"github.com/navidys/tvxwidgets"
)

var _ = Describe("GaugeAm", Ordered, func() {
	var (
		app       *tview.Application
		headerBox *tview.Box
		gaugeAm   *tvxwidgets.ActivityModeGauge
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = tview.NewApplication()
		headerBox = tview.NewBox().SetBorder(true)
		gaugeAm = tvxwidgets.NewActivityModeGauge()
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
			appLayout.AddItem(headerBox, 1, 0, true)
			appLayout.AddItem(gaugeAm, 50, 0, true)
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
			Expect(gaugeAm.HasFocus()).To(Equal(false))

			app.SetFocus(gaugeAm)
			gaugeAm.Pulse()
			app.Draw()
			// gauge will not get focus
			Expect(gaugeAm.HasFocus()).To(Equal(false))
		})
	})

	Describe("GetRect", func() {
		It("primitivie size", func() {
			x, y, width, heigth := gaugeAm.GetRect()
			Expect(x).To(Equal(0))
			Expect(y).To(Equal(1))
			Expect(width).To(Equal(80))
			Expect(heigth).To(Equal(50))
		})
	})
})
