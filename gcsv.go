package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	filePath, err := check()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
	}
	os.Setenv("FYNE_THEME", "light")
	csvapp := app.New()


	w := csvapp.NewWindow(filepath.Base(filePath))
	data , err := parse(filePath)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

	totalRows := len(data)
	totalCols := len(data[0]) + 1

	// Calculate max width for each column
	colWidths := make([]float32, totalCols)
	for col := 0; col < totalCols; col++ {
		maxWidth := float32(0)
		for row := 0; row < totalRows; row++ {
			var text string
			if col == 0 {
				if row == 0 {
					text = ""
				} else {
					text = fmt.Sprintf("%d", row)
				}
			} else {
				text = data[row][col-1]
			}
			lbl := canvas.NewText(text, color.Black)
			lbl.TextStyle = fyne.TextStyle{Bold: row == 0 || col == 0}
			lbl.TextSize = 14 // Default Fyne label size
			size := lbl.MinSize()
			if size.Width > maxWidth {
				maxWidth = size.Width
			}
		}
		// Add some padding
		colWidths[col] = maxWidth + 20
	}

	table := widget.NewTable(
		func() (int, int) {
			return totalRows, totalCols
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)

			// Set text for cell
			if id.Col == 0 {
				if id.Row == 0 {
					label.SetText("")
				} else {
					label.SetText(fmt.Sprintf("%d", id.Row))
				}
			} else {
				label.SetText(data[id.Row][id.Col-1])
			}

			// Bold header row and first column (row numbers), except [0][0]
			if id.Row == 0 || id.Col == 0 {
				if !(id.Row == 0 && id.Col == 0) {
					label.TextStyle = fyne.TextStyle{Bold: true}
				} else {
					label.TextStyle = fyne.TextStyle{}
				}
			} else {
				label.TextStyle = fyne.TextStyle{}
			}

			label.Refresh()
		},
	)

	// Set column widths based on content
	for col, width := range colWidths {
		table.SetColumnWidth(col, width)
	}

	w.SetContent(container.NewScroll(table))
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
