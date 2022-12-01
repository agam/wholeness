package main

import (
	"wholeness/wholeness"
)

func main() {
	m := wholeness.NewBouncingModel(30, 3)
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 30, Y: 30})

	driver := wholeness.NewDriver(m, w)
	driver.RunTUI()
}
