package main

import (
	"wholeness/wholeness"
)

func main() {
	m := wholeness.NewBlowupModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 5, Y: 5})

	driver := wholeness.NewDriver(m, w)
	driver.RunTUI()
}
