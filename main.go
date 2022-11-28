package main

import (
	"fmt"
	"github.com/rivo/tview"
)
import "wholeness/wholeness"

var app = tview.NewApplication()

func main() {
	fmt.Println("A new world")

	m := wholeness.NewBlowupModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 5, Y: 5})

	m.BigBang(w)
	w.RenderDebugDump()

	for i := 0; i < 10; i++ {
		w.RenderDebugDump()
		w.Tick()
	}
}
