package wholeness

type Model interface {
	BigBang(World)
}

type SimpleModel struct {
	agents []Agent
}

func (m *SimpleModel) BigBang(w World) {
	for _, agent := range m.agents {
		w.Add(agent, w.getRandomPosition())
	}
}

type FixedModel struct {
	agentMap map[Position]Agent
}

func (m *FixedModel) BigBang(w World) {
	for pos, agent := range m.agentMap {
		w.Add(agent, pos)
	}
}
