package wholeness

import (
	"fmt"
	"github.com/barweiss/go-tuple"
	"github.com/rivo/tview"
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
	DestroyAgent(id AgentID)

	RenderDebugDump()
	RenderTable(*tview.Table)

	getRandomPosition() Position
	getAgentsAtPosition(Position) []AgentID
}

type AgentIDSet map[AgentID]bool
type AgentList []Agent

//////////////////////////////////////// End Interfaces ////////////////////////////////////////

type simpleWorld struct {
	dimension Position
	positions map[Position]AgentIDSet
	agents    map[AgentID]tuple.T2[Agent, Position]
	nextID    AgentID
}

func NewSimpleWorld(dim Position) World {
	return &simpleWorld{
		dimension: dim,
		positions: make(map[Position]AgentIDSet),
		agents:    make(map[AgentID]tuple.T2[Agent, Position]),
	}
}

func (w *simpleWorld) getRandomPosition() Position {
	return Position{
		Y: rand.Intn(w.dimension.Y - 1),
		X: rand.Intn(w.dimension.X - 1),
	}
}

func (w *simpleWorld) newAgentID() AgentID {
	w.nextID++
	return w.nextID
}

func (w *simpleWorld) Add(a Agent, p Position) {
	id := w.newAgentID()
	w.agents[id] = tuple.New2(a, p)
	if _, ok := w.positions[p]; !ok {
		w.positions[p] = make(map[AgentID]bool)
	}
	w.positions[p][id] = true
	a.Init(p)
}

func (w *simpleWorld) Tick() {
	for id, agentPosition := range w.agents {
		ctx := simpleContext{id: id, world: w}
		agentPosition.V1.Tick(ctx)
	}
}

func (w *simpleWorld) UpdatePosition(id AgentID, oldPosition, newPosition Position) {
	delete(w.positions[oldPosition], id)
	if _, ok := w.positions[newPosition]; !ok {
		w.positions[newPosition] = make(map[AgentID]bool)
	}
	w.positions[newPosition][id] = true

	agentPosition := w.agents[id]
	w.agents[id] = tuple.New2(agentPosition.V1, newPosition)
}

func (w *simpleWorld) getAgentsAtPosition(pos Position) []AgentID {
	agentMap := w.positions[pos]
	agentList := make([]AgentID, 0)
	for agentID := range agentMap {
		agentList = append(agentList, agentID)
	}
	return agentList
}

func (w *simpleWorld) DestroyAgent(id AgentID) {
	agentPosition, ok := w.agents[id]
	if !ok {
		return
	}
	pos := agentPosition.V2
	posAgents := w.positions[pos]
	delete(posAgents, id)
	delete(w.agents, id)
}

type renderCellCallback func(found bool, agents AgentList, pos Position)
type renderRowEndCallback func()

func (w *simpleWorld) renderHelper(callback renderCellCallback, endCallback renderRowEndCallback) {
	for i := 0; i < w.dimension.Y; i++ {
		for j := 0; j < w.dimension.X; j++ {
			pos := Position{Y: i, X: j}
			agents, ok := w.positions[pos]
			if !ok {
				callback(false, AgentList{}, pos)
			}
			agentList := make([]Agent, 0)
			for id, _ := range agents {
				agentList = append(agentList, w.agents[id].V1)
			}
			callback(true, agentList, pos)
		}
		endCallback()
	}
}

func (w *simpleWorld) RenderDebugDump() {
	fmt.Println("\n-------- BEGIN WORLD -------- ")
	fmt.Println()
	w.renderHelper(func(found bool, agents AgentList, _ Position) {
		if !found || len(agents) == 0 {
			fmt.Print("   ")
		} else {
			fmt.Printf(" %d ", len(agents))
		}
	}, func() {
		fmt.Printf("|")
		fmt.Println()
	})
	fmt.Println("\n-------- END WORLD -------- ")
	fmt.Println()
}

const emptyCell = "   "

func (w *simpleWorld) RenderTable(table *tview.Table) {
	w.renderHelper(func(found bool, agents AgentList, pos Position) {
		if !found || len(agents) == 0 {
			table.SetCell(pos.Y, pos.X, tview.NewTableCell(emptyCell))
		} else {
			switch len(agents) {
			case 1:
				table.SetCell(pos.Y, pos.X, tview.NewTableCell(fmt.Sprintf(" %c ", agents[0].Render())))
			case 2:
				table.SetCell(
					pos.Y, pos.X,
					tview.NewTableCell(fmt.Sprintf("%c %c",
						agents[0].Render(),
						agents[1].Render())))
			case 3:
				table.SetCell(
					pos.Y, pos.X,
					tview.NewTableCell(fmt.Sprintf("%c%c%c",
						agents[0].Render(),
						agents[1].Render(),
						agents[2].Render())))
			default:
				table.SetCell(pos.Y, pos.X, tview.NewTableCell(fmt.Sprintf(" %d ", len(agents))))
			}
		}
	}, func() {})
}
