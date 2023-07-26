package common

type Card struct {
	rank RankEnumType
	suit SuitEnumType
}

func NewCard(rank RankEnumType, suit SuitEnumType) *Card {
	return &Card{
		rank: rank,
		suit: suit,
	}
}

func (c *Card) GetRank() RankEnumType {
	return c.rank
}

func (c *Card) GetSuit() SuitEnumType {
	return c.suit
}

type RankEnumType int

type SuitEnumType int

const (
	TWO RankEnumType = iota
	THREE
	FORE
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

const (
	Club SuitEnumType = iota
	Diamond
	Heart
	Spade
)
