package tvxwidgets_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rivo/tview"
)

var _ = Describe("Dialog", Ordered, func() {
	var (
		app       *tview.Application
		headerBox *tview.Box
		msgDialog *tvxwidgets.MessageDialog
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = tview.NewApplication()
		headerBox = tview.NewBox().SetBorder(true)
		msgDialog = tvxwidgets.NewMessageDialog(tvxwidgets.InfoDialog)
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := tview.NewFlex().SetDirection(tview.FlexRow)
			appLayout.AddItem(headerBox, 0, 1, true)
			appLayout.AddItem(msgDialog, 0, 1, true)
			err := app.SetScreen(screen).SetRoot(appLayout, true).Run()
			if err != nil {
				panic(err)
			}
		}()
	})

	AfterAll(func() {
		app.Stop()
	})

	Describe("NewMessageDialog", func() {
		It("returns a new message dialog primitive", func() {
			tests := []struct {
				msgType int
				bgColor tcell.Color
			}{
				{msgType: tvxwidgets.InfoDialog, bgColor: tcell.ColorSteelBlue},
				{msgType: tvxwidgets.ErrorDailog, bgColor: tcell.ColorOrangeRed},
			}

			for _, test := range tests {
				msgDialog.SetType(test.msgType)
				app.Draw()
				Expect(msgDialog.GetBackgroundColor()).To(Equal(test.bgColor))
			}
		})
	})

	Describe("Focus", func() {
		It("checks primitivie focus", func() {
			app.SetFocus(headerBox)
			app.Draw()
			Expect(msgDialog.HasFocus()).To(Equal(false))

			app.SetFocus(msgDialog)
			app.Draw()
			Expect(msgDialog.HasFocus()).To(Equal(true))
		})
	})
})
