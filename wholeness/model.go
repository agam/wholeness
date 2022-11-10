package wholeness

type Model interface {
	BigBang(World)
}

type SimpleModel struct {
	agents []Agent
}

func (m *SimpleModel) BigBang(w World) {
	for _, agent := range m.agents {
		w.Add(agent, w.GetRandomPosition())
	}
}

type noOpModel struct {
	SimpleModel
}

func NewNoOpModel() *noOpModel {
	return &noOpModel{
		SimpleModel{
			agents: []Agent{&noOpAgent{}},
		},
	}
}
