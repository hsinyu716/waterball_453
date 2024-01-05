package pattern

import (
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStraightTestSuite(t *testing.T) {
	suite.Run(t, new(StraightTestSuite))
}

type StraightTestSuite struct {
	suite.Suite
	cards []poker.Card
}

func (s *StraightTestSuite) SetupSuite() {
	var cards []poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, *card)
	card = poker.NewCard(2, 1)
	cards = append(cards, *card)
	card = poker.NewCard(3, 1)
	cards = append(cards, *card)
	card = poker.NewCard(5, 1)
	cards = append(cards, *card)
	card = poker.NewCard(4, 1)
	cards = append(cards, *card)
	s.cards = cards
}

func (s *StraightTestSuite) TestStraight() {
	straight := NewPatternStraight(s.cards, nil)

	var (
		cards []poker.Card
		card  *poker.Card
	)
	{
		card = poker.NewCard(11, 1)
		cards = append(cards, *card)
		card = poker.NewCard(12, 1)
		cards = append(cards, *card)
		card = poker.NewCard(0, 1)
		cards = append(cards, *card)
		card = poker.NewCard(1, 1)
		cards = append(cards, *card)
		card = poker.NewCard(2, 1)
		cards = append(cards, *card)
		straight.SetCards(cards)
		// [{12 1} {11 1} {2 1} {1 1} {0 1}]
		//   2      1      5     4     3
		s.True(straight.Validate())
	}

	{
		cards = []poker.Card{}
		card = poker.NewCard(12, 1)
		cards = append(cards, *card)
		card = poker.NewCard(0, 1)
		cards = append(cards, *card)
		card = poker.NewCard(1, 1)
		cards = append(cards, *card)
		card = poker.NewCard(2, 1)
		cards = append(cards, *card)
		card = poker.NewCard(3, 1)
		cards = append(cards, *card)
		straight.SetCards(cards)
		// [{12 1} {3 1} {2 1} {1 1} {0 1}]
		//    2     6     5     4     3
		s.True(straight.Validate())
	}
	{
		cards = []poker.Card{}
		card = poker.NewCard(11, 1)
		cards = append(cards, *card)
		card = poker.NewCard(10, 1)
		cards = append(cards, *card)
		card = poker.NewCard(9, 1)
		cards = append(cards, *card)
		card = poker.NewCard(8, 1)
		cards = append(cards, *card)
		card = poker.NewCard(7, 1)
		cards = append(cards, *card)
		straight.SetCards(cards)
		// [{11 1} {10 1} {9 1} {8 1} {7 1}]
		//    A     K     Q     J     10
		s.True(straight.Validate())
	}
	{
		cards = []poker.Card{}
		card = poker.NewCard(0, 1)
		cards = append(cards, *card)
		card = poker.NewCard(9, 1)
		cards = append(cards, *card)
		card = poker.NewCard(10, 1)
		cards = append(cards, *card)
		card = poker.NewCard(11, 1)
		cards = append(cards, *card)
		card = poker.NewCard(12, 1)
		cards = append(cards, *card)
		straight.SetCards(cards)
		// [{11 1} {10 1} {9 1} {0 1} {12 1}]
		//    A     K     Q     3     2
		s.True(straight.Validate())
	}
	{
		cards = []poker.Card{}
		card = poker.NewCard(0, 1)
		cards = append(cards, *card)
		card = poker.NewCard(1, 1)
		cards = append(cards, *card)
		card = poker.NewCard(10, 1)
		cards = append(cards, *card)
		card = poker.NewCard(11, 1)
		cards = append(cards, *card)
		card = poker.NewCard(12, 1)
		cards = append(cards, *card)
		straight.SetCards(cards)
		// [{12 1} {11 1} {10 1} {1 1} {0 1}]
		//    2     A     K     4     3
		s.True(straight.Validate())
	}
	//
	//cards = []common.Card{}
	//card = common.NewCard(12, 1)
	//cards = append(cards, *card)
	//card = common.NewCard(11, 1)
	//cards = append(cards, *card)
	//card = common.NewCard(10, 1)
	//cards = append(cards, *card)
	//card = common.NewCard(0, 1)
	//cards = append(cards, *card)
	//card = common.NewCard(1, 1)
	//cards = append(cards, *card)
	//straight.SetCards(cards)
	//// [{12 1} {11 1} {10 1} {1 1} {0 1}]
	////   2       A     K     4     3
	//s.True(straight.Validate())
}

func (s *StraightTestSuite) TestStraightMax() {
	straight := NewPatternStraight(s.cards, nil)
	s.Equal(*poker.NewCard(5, 1), straight.GetMax())
}

func (s *StraightTestSuite) TestStraightFail() {
	var cards []poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, *card)
	card = poker.NewCard(2, 1)
	cards = append(cards, *card)
	card = poker.NewCard(3, 1)
	cards = append(cards, *card)
	card = poker.NewCard(5, 1)
	cards = append(cards, *card)
	card = poker.NewCard(6, 1)
	cards = append(cards, *card)
	straight := NewPatternStraight(cards, nil)
	s.Nil(straight)
}
