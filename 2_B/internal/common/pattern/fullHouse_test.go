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
	cards       []*poker.Card
	cardPattern ICardPattern
}

func (f *FullHouseTestSuite) SetupSuite() {
	var cards []*poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, card)
	card = poker.NewCard(1, 2)
	cards = append(cards, card)
	card = poker.NewCard(2, 1)
	cards = append(cards, card)
	card = poker.NewCard(2, 3)
	cards = append(cards, card)
	card = poker.NewCard(2, 2)
	cards = append(cards, card)
	f.cards = cards
	f.cardPattern = NewPatternFullHouse(nil)
}

func (f *FullHouseTestSuite) TestFullHouse() {
	f.cardPattern.Validate(f.cards)
}

func (f *FullHouseTestSuite) TestFullHouseMax() {
	validate := f.cardPattern.Validate(f.cards)
	f.Equal(*poker.NewCard(2, 3), validate.GetMax())
}

func (f *FullHouseTestSuite) TestFullHouseFail() {
	var cards []*poker.Card
	card := poker.NewCard(1, 1)
	cards = append(cards, card)
	card = poker.NewCard(1, 2)
	cards = append(cards, card)
	card = poker.NewCard(2, 3)
	cards = append(cards, card)
	card = poker.NewCard(2, 3)
	cards = append(cards, card)
	card = poker.NewCard(3, 4)
	cards = append(cards, card)
	f.Nil(f.cardPattern.Validate(cards))
}
