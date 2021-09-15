package demo

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"math/rand"
)

func QuickStart() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

func Box() {
	box := tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:] [::bu]title")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

func Button() {
	app := tview.NewApplication()
	button := tview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
		app.Stop()
	})
	button.SetBorder(true).SetRect(0, 0, 22, 3)
	if err := app.SetRoot(button, false).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func Checkbox() {
	app := tview.NewApplication()
	checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	checkbox.SetChangedFunc(func(checked bool) {
		checkbox.SetBackgroundColor([]tcell.Color{tcell.ColorSkyblue, tcell.ColorBlueViolet, tcell.ColorPink}[rand.Intn(3)])
	})
	if err := app.SetRoot(checkbox, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
