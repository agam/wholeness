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

func NewNoOpModel() SimpleModel {
	return SimpleModel{
		agents: []Agent{&noOpAgent{}, &noOpAgent{}},
	}
}

func NewMovingModel() SimpleModel {
	return SimpleModel{
		agents: []Agent{
			&constantMoveAgent{change: Position{Y: 1, X: 0}},
			&constantMoveAgent{change: Position{Y: 0, X: -1}},
		},
	}
}
