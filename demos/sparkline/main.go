package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	iowaitSparkline := tvxwidgets.NewSparkline()
	iowaitSparkline.SetBorder(false)
	iowaitSparkline.SetDataTitle("Disk I/O (iowait)")
	iowaitSparkline.SetBorderColor(tcell.ColorDimGray)
	iowaitSparkline.SetTitleColor(tcell.ColorDimGray)
	iowaitSparkline.SetDataTitleColor(tcell.ColorDarkOrange)
	iowaitSparkline.SetLineColor(tcell.ColorMediumPurple)

	systemSparkline := tvxwidgets.NewSparkline()
	systemSparkline.SetBorder(false)
	systemSparkline.SetDataTitle("Disk I/O (system)")
	systemSparkline.SetBorderColor(tcell.ColorDimGray)
	systemSparkline.SetTitleColor(tcell.ColorDimGray)
	systemSparkline.SetDataTitleColor(tcell.ColorDarkOrange)
	systemSparkline.SetLineColor(tcell.ColorSteelBlue)

	iowaitData := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	systemData := []float64{0, 0, 1, 2, 9, 5, 3, 1, 2, 0, 6, 3, 2, 2, 6, 8, 5, 2, 1, 5, 8, 6, 1, 4, 1, 1, 4, 3, 6}

	ioSparkLineData := func() []float64 {
		for i := 0; i < 5; i++ {
			iowaitData = append(iowaitData, iowaitData...)
		}

		return iowaitData
	}()

	systemSparklineData := func() []float64 {
		for i := 0; i < 5; i++ {
			systemData = append(systemData, systemData...)
		}

		return systemData
	}()

	iowaitSparkline.SetData(ioSparkLineData)
	systemSparkline.SetData(systemSparklineData)

	sparklineGroupLayout := tview.NewFlex().SetDirection(tview.FlexRow)
	sparklineGroupLayout.SetBorder(true)
	sparklineGroupLayout.SetBorderColor(tcell.ColorDimGray)
	sparklineGroupLayout.SetTitle("DISK IO")
	sparklineGroupLayout.SetTitleColor(tcell.ColorDarkOrange)
	sparklineGroupLayout.AddItem(iowaitSparkline, 0, 1, false)
	sparklineGroupLayout.AddItem(tview.NewBox(), 1, 0, false)
	sparklineGroupLayout.AddItem(systemSparkline, 0, 1, false)

	moveData := func() ([]float64, []float64) {

		newIOWaitData := ioSparkLineData[1:]
		newIOWaitData = append(newIOWaitData, ioSparkLineData[0])
		ioSparkLineData = newIOWaitData

		newSystemData := systemSparklineData[1:]
		newSystemData = append(newSystemData, systemSparklineData[0])
		systemSparklineData = newSystemData

		return newIOWaitData, newSystemData
	}

	update := func() {
		tick := time.NewTicker(500 * time.Millisecond)

		for {
			select {
			case <-tick.C:
				d1, d2 := moveData()
				iowaitSparkline.SetData(d1)
				systemSparkline.SetData(d2)

				app.Draw()
			}
		}
	}

	go update()

	if err := app.SetRoot(sparklineGroupLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
