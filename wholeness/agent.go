package wholeness

type Agent interface {
	Init(Position)
	Tick(AgentContext)
	Render() rune
}

type AgentSet map[Agent]bool
