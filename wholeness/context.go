package wholeness

type AgentContext interface {
	Move(Position, Position)
	Look(Position) []AgentID
	Destroy(AgentID)
}

type simpleContext struct {
	id    AgentID
	world World
}

func (s simpleContext) Move(old, new Position) {
	s.world.UpdatePosition(s.id, old, new)
}

func (s simpleContext) Look(position Position) []AgentID {
	return s.world.getAgentsAtPosition(position)
}

func (s simpleContext) Destroy(id AgentID) {
	//TODO implement me
	panic("implement me")
}
