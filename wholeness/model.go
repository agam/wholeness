package wholeness

type Model struct {
	agentMap map[Position]Agent
	agents   []Agent
}

func (m *Model) BigBang(w World) {
	for _, agent := range m.agents {
		w.Add(agent, w.getRandomPosition())
	}
	for pos, agent := range m.agentMap {
		w.Add(agent, pos)
	}
}
