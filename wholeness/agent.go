package wholeness

type Agent interface {
	Init()
	Tick()
}

type AgentSet map[Agent]bool

type noOpAgent struct {
}

func (a *noOpAgent) Init() {}

func (a *noOpAgent) Tick() {}
