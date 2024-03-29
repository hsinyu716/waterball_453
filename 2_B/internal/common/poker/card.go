package poker

type Card struct {
	rank RankEnumType
	suit SuitEnumType
}

type Cards struct {
	Cards []*Card
}

type CompareResult int

const (
	// CompareResultSmaller is a CompareResult of type Smaller.
	CompareResultSmaller CompareResult = iota
	// CompareResultBigger is a CompareResult of type Bigger.
	CompareResultBigger
	// CompareResultEqual is a CompareResult of type Equal.
	CompareResultEqual
	// CompareResultInvalid is a CompareResult of type Invalid.
	CompareResultInvalid
)

func (c *Card) Compare(card *Card) CompareResult {
	if c.rank < card.rank {
		return CompareResultSmaller
	}
	if c.rank == card.rank && c.suit < card.suit {
		return CompareResultSmaller
	}
	return CompareResultBigger
}

// DiffRank return diff rank
func (c *Card) DiffRank(card *Card) RankEnumType {
	return c.rank - card.rank
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
	THREE RankEnumType = iota
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
	TWO
)

var (
	RankMap = map[RankEnumType]string{
		THREE: "3",
		FORE:  "4",
		FIVE:  "5",
		SIX:   "6",
		SEVEN: "7",
		EIGHT: "8",
		NINE:  "9",
		TEN:   "10",
		JACK:  "J",
		QUEEN: "Q",
		KING:  "K",
		ACE:   "A",
		TWO:   "2",
	}

	RankStringMap = map[string]RankEnumType{
		"3":  THREE,
		"4":  FORE,
		"5":  FIVE,
		"6":  SIX,
		"7":  SEVEN,
		"8":  EIGHT,
		"9":  NINE,
		"10": TEN,
		"J":  JACK,
		"Q":  QUEEN,
		"K":  KING,
		"A":  ACE,
		"2":  TWO,
	}

	SuitMap = map[SuitEnumType]string{
		Club:    "C",
		Diamond: "D",
		Heart:   "H",
		Spade:   "S",
	}

	SuitStringMap = map[string]SuitEnumType{
		"C": Club,
		"D": Diamond,
		"H": Heart,
		"S": Spade,
	}
)

const (
	Club SuitEnumType = iota
	Diamond
	Heart
	Spade
)
