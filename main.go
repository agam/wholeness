package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)
import "wholeness/wholeness"

var app = tview.NewApplication()

func main() {
	fmt.Println("A new world")

	//consoleDriver()
	tuiDriver()
}

// TODO: factor out a common "driver" interface
func consoleDriver() {
	m := wholeness.NewBlowupModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 5, Y: 5})

	m.BigBang(w)
	w.RenderDebugDump()

	for i := 0; i < 10; i++ {
		w.RenderDebugDump()
		w.Tick()
	}
}

func tuiDriver() {
	table := tview.NewTable()
	table.SetBorder(true).SetBorderColor(tcell.ColorWhite)

	m := wholeness.NewBouncingModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 10, Y: 10})

	m.BigBang(w)

	step := func() {
		w.RenderTable(table)
		w.Tick()
		time.Sleep(time.Second)
		app.Draw()
	}

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			go step()
			return nil
		}
		return event
	})

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
