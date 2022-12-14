package wholeness

import (
	"fmt"
	"math/rand"
)

//////////////////////////////////////// No-Op
type noOpAgent struct {
}

func (a *noOpAgent) Init(_ Position) {}

func (a *noOpAgent) Tick(_ AgentContext) {}

func (a *noOpAgent) Render() rune {
	return '.'
}

func (a *noOpAgent) IsFixed() bool {
	return true
}

func NewNoOpModel() *Model {
	return &Model{
		agents: []Agent{&noOpAgent{}, &noOpAgent{}},
	}
}

//////////////////////////////////////// Constant-Movement
type constantMoveAgent struct {
	current Position
	change  Position
}

func (a *constantMoveAgent) Init(pos Position) {
	a.current = pos
}

func (a *constantMoveAgent) Tick(ctx AgentContext) {
	newPosition := a.current.Add(a.change)
	ctx.Move(a.current, newPosition)
	a.current = newPosition
}

func (a *constantMoveAgent) Render() rune {
	return '*'
}

func (a *constantMoveAgent) IsFixed() bool {
	return false
}

func NewMovingModel() *Model {
	return &Model{
		agents: []Agent{
			&constantMoveAgent{change: Position{Y: 1, X: 0}},
			&constantMoveAgent{change: Position{Y: 0, X: -1}},
		},
	}
}

//////////////////////////////////////// Bouncing around
var (
	left  = Position{Y: 0, X: -1}
	right = Position{Y: 0, X: 1}
	up    = Position{Y: -1, X: 0}
	down  = Position{Y: 1, X: 0}

	directions = []Position{left, right, up, down}
)

func getRandomDirection() Position {
	randIndex := rand.Intn(len(directions))
	return directions[randIndex]
}

func getOppositeDirection(dir Position) Position {
	switch dir {
	case left:
		return right
	case right:
		return left
	case up:
		return down
	case down:
		return up
	default:
		return Position{0, 0}
	}
}

type bouncingAgent struct {
	current Position
	drift   Position
}

func (a *bouncingAgent) Init(pos Position) {
	a.current = pos
	a.drift = getRandomDirection()
}

func (a *bouncingAgent) Tick(ctx AgentContext) {
	// Keep moving until we "hit" something; if we do, switch directions
	nextPos := a.current.Add(a.drift)
	contents := ctx.Look(nextPos)
	if len(contents) == 0 {
		ctx.Move(a.current, nextPos)
		a.current = nextPos
	} else {
		a.drift = getOppositeDirection(a.drift)
		// Skip a turn while we pivot
	}
}

func (a *bouncingAgent) Render() rune {
	return '????'
}

func (a *bouncingAgent) IsFixed() bool {
	return false
}

func NewBouncingModel(dim int, numballs int) *Model {
	fixedAgents := make(map[Position]Agent, 0)
	// Make a "solid" box
	for i := 0; i < dim; i++ {
		fixedAgents[Position{Y: 0, X: i}] = &noOpAgent{}
		fixedAgents[Position{Y: dim - 1, X: i}] = &noOpAgent{}
		fixedAgents[Position{Y: i, X: 0}] = &noOpAgent{}
		fixedAgents[Position{Y: i, X: dim - 1}] = &noOpAgent{}
	}

	floatingAgents := make([]Agent, 0)
	for i := 0; i < numballs; i++ {
		floatingAgents = append(floatingAgents, &bouncingAgent{})
	}

	return &Model{
		agentMap: fixedAgents,
		agents:   floatingAgents,
		name:     "Bouncing Balls",
	}
}

//////////////////////////////////////// Blowing up
type blowupAgent struct {
	current Position
	drift   Position
}

func (a *blowupAgent) Init(pos Position) {
	a.current = pos
	a.drift = getRandomDirection()
}

func (a *blowupAgent) Tick(ctx AgentContext) {
	nextPos := a.current.Add(a.drift)
	contents := ctx.Look(nextPos)
	if len(contents) != 0 {
		fmt.Println("BOOM !! ")
		ctx.SelfDestruct()
		return
	}
	ctx.Move(a.current, nextPos)
	a.current = nextPos
}

func (a *blowupAgent) Render() rune {
	return 'O'
}

func (a *blowupAgent) IsFixed() bool {
	return false
}

func NewBlowupModel() *Model {
	fixedAgents := make(map[Position]Agent, 0)
	// Make a "solid" box
	const dim = 4
	for i := 0; i < dim; i++ {
		fixedAgents[Position{Y: 0, X: i}] = &noOpAgent{}
		fixedAgents[Position{Y: dim, X: i}] = &noOpAgent{}
		fixedAgents[Position{Y: i, X: 0}] = &noOpAgent{}
		fixedAgents[Position{Y: i, X: dim}] = &noOpAgent{}
	}

	// Add a "bomb" inside.
	floatingAgents := make([]Agent, 0)
	floatingAgents = append(floatingAgents, &blowupAgent{})

	return &Model{
		agentMap: fixedAgents,
		agents:   floatingAgents,
	}
}
