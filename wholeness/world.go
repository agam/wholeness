package wholeness

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//////////////////////////////////////// Begin Interfaces ////////////////////////////////////////
type AgentID int

type World interface {
	Tick()
	Add(Agent, Position)
	UpdatePosition(id AgentID, old, new Position)

	RenderDebugDump()

	getRandomPosition() Position
	getAgentsAtPosition(Position) []AgentID
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

func (w *simpleWorld) getRandomPosition() Position {
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
	a.Init(p)
}

func (w *simpleWorld) Tick() {
	for id, agent := range w.agents {
		ctx := simpleContext{id: id, world: w}
		agent.Tick(ctx)
	}
}

func (w *simpleWorld) UpdatePosition(id AgentID, oldPosition, newPosition Position) {
	delete(w.positions[oldPosition], id)
	if _, ok := w.positions[newPosition]; !ok {
		w.positions[newPosition] = make(map[AgentID]bool)
	}
	w.positions[newPosition][id] = true
}

func (w *simpleWorld) getAgentsAtPosition(pos Position) []AgentID {
	agentMap := w.positions[pos]
	agentList := make([]AgentID, 0)
	for agentID := range agentMap {
		agentList = append(agentList, agentID)
	}
	return agentList
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
