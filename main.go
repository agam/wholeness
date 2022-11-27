package main

import (
	"fmt"
	"github.com/rivo/tview"
)
import "wholeness/wholeness"

var app = tview.NewApplication()

func main() {
	fmt.Println("A new world")

	m := wholeness.NewMovingModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 5, Y: 5})

	m.BigBang(w)
	w.RenderDebugDump()
	w.Tick()
	w.RenderDebugDump()
	w.Tick()
	w.RenderDebugDump()
}
