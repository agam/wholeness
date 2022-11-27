package wholeness

type Agent interface {
	Init(Position)
	Tick(AgentContext)
}

type AgentSet map[Agent]bool
