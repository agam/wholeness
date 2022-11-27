package wholeness

type AgentContext interface {
	Move(Position, Position)
	Look(Position) []Agent
	Destroy(AgentID)
}

type simpleContext struct {
	id    AgentID
	world World
}

func (s simpleContext) Move(old, new Position) {
	s.world.UpdatePosition(s.id, old, new)
}

func (s simpleContext) Look(position Position) []Agent {
	//TODO implement me
	panic("implement me")
}

func (s simpleContext) Destroy(id AgentID) {
	//TODO implement me
	panic("implement me")
}
