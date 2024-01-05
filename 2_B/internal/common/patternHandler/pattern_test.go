package patternHandler

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPatternTestSuite(t *testing.T) {
	suite.Run(t, new(PatternTestSuite))
}

type PatternTestSuite struct {
	suite.Suite
	cards []poker.Card
}

func (s *PatternTestSuite) SetupSuite() {
	var cards []poker.Card
	card := poker.NewCard(7, 0)
	cards = append(cards, *card)
	card = poker.NewCard(7, 1)
	cards = append(cards, *card)
	s.cards = cards
}

func (s *PatternTestSuite) TestPattern() {
	top := pattern.NewPatternPair(s.cards, nil)
	validate := top.Validate()
	s.True(validate)

	var cards []poker.Card
	card := poker.NewCard(7, 2)
	cards = append(cards, *card)
	card = poker.NewCard(7, 3)
	cards = append(cards, *card)

	handler := NewSingleHandler(NewPairHandler(NewStraightHandler(nil)))
	compare := handler.Validate(pattern.NewPatternPair(cards, nil), top)
	s.True(compare)
}

func (s *PatternTestSuite) TestPatternStraight() {
	var cards []poker.Card
	card := poker.NewCard(2, 2)
	cards = append(cards, *card)
	card = poker.NewCard(3, 3)
	cards = append(cards, *card)
	card = poker.NewCard(4, 3)
	cards = append(cards, *card)
	card = poker.NewCard(5, 3)
	cards = append(cards, *card)
	card = poker.NewCard(6, 3)
	cards = append(cards, *card)
	topPlay := pattern.NewPatternSingle(cards, pattern.NewPatternStraight(cards, nil))

	cards = []poker.Card{}
	card = poker.NewCard(12, 2)
	cards = append(cards, *card)
	card = poker.NewCard(0, 3)
	cards = append(cards, *card)
	card = poker.NewCard(1, 3)
	cards = append(cards, *card)
	card = poker.NewCard(2, 3)
	cards = append(cards, *card)
	card = poker.NewCard(3, 3)
	cards = append(cards, *card)

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

	handler := NewSingleHandler(NewPairHandler(NewStraightHandler(nil)))

	c0 := pattern.NewPatternSingle(cards, pattern.NewPatternStraight(cards, nil))

	compare := handler.Validate(c0, topPlay)
	s.True(compare)
}
