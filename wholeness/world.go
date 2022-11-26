package wholeness

import (
	"fmt"
	"math/rand"
)

//////////////////////////////////////// Begin Interfaces ////////////////////////////////////////
type AgentID int

type Position struct {
	X, Y int
}

type World interface {
	Tick()
	Add(Agent, Position)

	GetRandomPosition() Position
	RenderDebugDump()
}

type AgentContext interface {
	Move(AgentID, Position)
	Look(Position) []Agent
	Destroy(AgentID)
}

type AgentIDSet map[AgentID]bool

//////////////////////////////////////// End Interfaces ////////////////////////////////////////

type simpleWorld struct {
	dimension Position
	positions map[Position]AgentIDSet
	agents    map[AgentID]Agent
	nextID    AgentID
}

func NewSimpleWorld(dim Position) *simpleWorld {
	return &simpleWorld{
		dimension: dim,
		positions: make(map[Position]AgentIDSet),
		agents:    make(map[AgentID]Agent),
	}
}

func (w *simpleWorld) GetRandomPosition() Position {
	return Position{
		Y: rand.Intn(w.dimension.Y),
		X: rand.Intn(w.dimension.X),
	}
}

func (w *simpleWorld) newAgentID() AgentID {
	w.nextID++
	return w.nextID
}

func (w *simpleWorld) Add(a Agent, p Position) {
	id := w.newAgentID()
	w.agents[id] = a
	if _, ok := w.positions[p]; !ok {
		w.positions[p] = make(map[AgentID]bool)
	}
	w.positions[p][id] = true
}

func (w *simpleWorld) Tick() {
	for _, agent := range w.agents {
		agent.Tick(w)
	}
}

func (w *simpleWorld) Move(id AgentID, position Position) {
	//TODO implement me
	panic("implement me")
}

func (w *simpleWorld) Look(position Position) []Agent {
	//TODO implement me
	panic("implement me")
}

func (w *simpleWorld) Destroy(id AgentID) {
	//TODO implement me
	panic("implement me")
}

func (w *simpleWorld) RenderDebugDump() {
	fmt.Println("\n-------- BEGIN WORLD -------- ")
	fmt.Println()
	for i := 0; i < w.dimension.Y; i++ {
		fmt.Printf("|")
		for j := 0; j < w.dimension.X; j++ {
			pos := Position{Y: i, X: j}
			agents, ok := w.positions[pos]
			if !ok || len(agents) == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf(" %d ", len(agents))
			}
		}
		fmt.Printf("|")
		fmt.Println()
	}
	fmt.Println("\n-------- END WORLD -------- ")
	fmt.Println()
}
