package wholeness

//////////////////////////////////////// No-Op
type noOpAgent struct {
}

func (a *noOpAgent) Init(_ Position) {}

func (a *noOpAgent) Tick(_ AgentContext) {}

func NewNoOpModel() SimpleModel {
	return SimpleModel{
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
	newPosition := Position{
		Y: a.current.Y + a.change.Y,
		X: a.current.X + a.change.X,
	}
	ctx.Move(a.current, newPosition)
	a.current = newPosition
}

func NewMovingModel() SimpleModel {
	return SimpleModel{
		agents: []Agent{
			&constantMoveAgent{change: Position{Y: 1, X: 0}},
			&constantMoveAgent{change: Position{Y: 0, X: -1}},
		},
	}
}
