package main

import (
	"time"

	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	grid := tview.NewGrid()
	grid.SetBorder(true).SetTitle("Spinners")

	spinners := [][]*tvxwidgets.Spinner{
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerDotsCircling),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerDotsUpDown),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerBounce),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerLine),
		},
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCircleQuarters),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerSquareCorners),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCircleHalves),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerCorners),
		},
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerArrows),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerHamburger),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerStack),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerStar),
		},
		{
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerGrowHorizontal),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerGrowVertical),
			tvxwidgets.NewSpinner().SetStyle(tvxwidgets.SpinnerBoxBounce),
			tvxwidgets.NewSpinner().SetCustomStyle([]rune{'🕛', '🕐', '🕑', '🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚'}),
		},
	}

	for rowIdx, row := range spinners {
		for colIdx, spinner := range row {
			grid.AddItem(spinner, rowIdx, colIdx, 1, 1, 1, 1, false)
		}
	}

	update := func() {
		tick := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-tick.C:
				for _, row := range spinners {
					for _, spinner := range row {
						spinner.Pulse()
					}
				}
				app.Draw()
			}
		}
	}
	go update()

	if err := app.SetRoot(grid, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
