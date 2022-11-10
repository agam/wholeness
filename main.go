package main

import "fmt"
import "wholeness/wholeness"

func main() {
	fmt.Println("A new world")

	m := wholeness.NewNoOpModel()
	w := wholeness.NewSimpleWorld(wholeness.Position{X: 5, Y: 5})

	m.BigBang(w)
}
