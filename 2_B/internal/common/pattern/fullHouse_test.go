package pattern

import (
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestFullHouseTestSuite(t *testing.T) {
	suite.Run(t, new(FullHouseTestSuite))
}

type FullHouseTestSuite struct {
	suite.Suite
	cards []poker.Card
}

func (s *FullHouseTestSuite) SetupSuite() {
	var cards []poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, *card)
	card = poker.NewCard(1, 2)
	cards = append(cards, *card)
	card = poker.NewCard(2, 1)
	cards = append(cards, *card)
	card = poker.NewCard(2, 3)
	cards = append(cards, *card)
	card = poker.NewCard(2, 2)
	cards = append(cards, *card)
	s.cards = cards
}

func (s *FullHouseTestSuite) TestFullHouse() {
	fullHouse := NewPatternFullHouse(s.cards, nil)
	validate := fullHouse.Validate()
	s.True(validate)
}

func (s *FullHouseTestSuite) TestFullHouseMax() {
	fullHouse := NewPatternFullHouse(s.cards, nil)
	s.Equal(*poker.NewCard(2, 3), fullHouse.GetMax())
}

func (s *FullHouseTestSuite) TestFullHouseFail() {
	var cards []poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, *card)
	card = poker.NewCard(1, 2)
	cards = append(cards, *card)
	card = poker.NewCard(2, 3)
	cards = append(cards, *card)
	card = poker.NewCard(2, 3)
	cards = append(cards, *card)
	card = poker.NewCard(3, 4)
	cards = append(cards, *card)
	fullHouse := NewPatternFullHouse(cards, nil)
	s.Nil(fullHouse)
}
