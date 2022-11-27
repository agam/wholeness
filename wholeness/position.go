package wholeness

type Position struct {
	X, Y int
}

func (p Position) Add(delta Position) Position {
	return Position{
		Y: p.Y + delta.Y,
		X: p.X + delta.X,
	}
}
