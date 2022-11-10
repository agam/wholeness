package wholeness

import (
	"fmt"
	"math/rand"
)

type Position struct {
	X, Y int
}

type World interface {
	Tick()
	Add(Agent, Position)

	GetRandomPosition() Position
	RenderDebugDump()
}

type AgentSetPositions map[Position]AgentSet

type simpleWorld struct {
	dimension Position
	positions AgentSetPositions
}

func NewSimpleWorld(dim Position) *simpleWorld {
	return &simpleWorld{
		dimension: dim,
		positions: make(AgentSetPositions),
	}
}

func (w *simpleWorld) GetRandomPosition() Position {
	return Position{
		Y: rand.Intn(w.dimension.Y),
		X: rand.Intn(w.dimension.X),
	}
}

func (w *simpleWorld) Add(a Agent, p Position) {

}

func (w *simpleWorld) Tick() {

}

func (w *simpleWorld) RenderDebugDump() {
	fmt.Println("\n=== Dumping world ===\n\n")
}
