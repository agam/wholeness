package main

import (
	"fmt"
	"github.com/rivo/tview"
)
import "wholeness/wholeness"

var app = tview.NewApplication()

func main() {
	fmt.Println("A new world")

	consoleDriver()
	//tuiDriver()
}

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
	table := tview.NewTable().SetBorder(true)
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}

	m := wholeness.NewBouncingModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 10, Y: 10})

	m.BigBang(w)

	for i := 0; i < 10; i++ {
	}
}
