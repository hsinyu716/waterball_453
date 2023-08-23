package common

type AI struct {
	PlayerAdapter
}

func (A AI) TakeTurn() *TurnMove {
	return nil
}

func (A AI) GainPoint() {
}

func (A AI) GetPoint() int {
	return 0
}
