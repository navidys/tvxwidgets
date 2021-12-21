# tvxwidgets

tvxwidgets provides extra widgets for [tview](https://github.com/rivo/tview).  
`NOTE:` The project is at its early stages and under development, feel free to contribute and report bugs.

![Screenshot](tvxwidgets.gif)

## Example

```go
package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	gauge := tvxwidgets.NewActivityModeGauge()
	gauge.SetTitle("activity mode gauge")
	gauge.SetPgBgColor(tcell.ColorOrange)
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetBorder(true)
	if err := app.SetRoot(gauge, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
```

## Available Widgets

* bar chart
* activity mode gauge
* percentage mode gauge
* utilisation mode gauge
