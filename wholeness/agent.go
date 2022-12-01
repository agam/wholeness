package wholeness

type Agent interface {
	Init(Position)
	Tick(AgentContext)
	Render() rune
	IsFixed() bool
}

type AgentSet map[Agent]bool
