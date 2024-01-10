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
	cards []*poker.Card
}

func (p *PairTestSuite) SetupSuite() {
	var cards []*poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, card)
	card = poker.NewCard(1, 2)
	cards = append(cards, card)
	p.cards = cards
}

func (p *PairTestSuite) TestPair() {
	pair := NewPatternPair(nil)
	validate := pair.Validate(p.cards)
	p.T().Log(validate)
}

func (p *PairTestSuite) TestPairMax() {
	pair := NewPatternPair(nil)
	pair.Validate(p.cards)
	p.Equal(*poker.NewCard(1, 2), pair.GetMax())
}

func (p *PairTestSuite) TestPairFail() {
	var cards []*poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, card)
	card = poker.NewCard(2, 2)
	cards = append(cards, card)
	pair := NewPatternPair(nil)
	p.Nil(pair.Validate(cards))
}
