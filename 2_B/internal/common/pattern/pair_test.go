package pattern

import (
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestPairTestSuite(t *testing.T) {
	suite.Run(t, new(PairTestSuite))
}

type PairTestSuite struct {
	suite.Suite
	cards []poker.Card
}

func (s *PairTestSuite) SetupSuite() {
	var cards []poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, *card)
	card = poker.NewCard(1, 2)
	cards = append(cards, *card)
	s.cards = cards
}

func (s *PairTestSuite) TestPair() {
	pair := NewPatternPair(s.cards, nil)
	validate := pair.Validate()
	s.True(validate)
}

func (s *PairTestSuite) TestPairMax() {
	pair := NewPatternPair(s.cards, nil)
	s.Equal(*poker.NewCard(1, 2), pair.GetMax())
}

func (s *PairTestSuite) TestPairFail() {
	var cards []poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, *card)
	card = poker.NewCard(2, 2)
	cards = append(cards, *card)
	pair := NewPatternPair(cards, nil)
	s.Nil(pair)
}
