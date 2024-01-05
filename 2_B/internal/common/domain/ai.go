package domain

type AI struct {
	Player
}

func NewAI() IPlayer {
	return &AI{
		Player{},
	}
}
