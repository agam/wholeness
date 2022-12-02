package wholeness

type Model struct {
	agentMap map[Position]Agent
	agents   []Agent
	name     string
}

func (m *Model) BigBang(w World) {
	for _, agent := range m.agents {
		w.Add(agent, w.getRandomPosition())
	}
	for pos, agent := range m.agentMap {
		w.Add(agent, pos)
	}
}
