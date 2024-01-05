package patternHandler

import (
	"bufio"
	"cosmos.big2/internal/common"
	"cosmos.big2/internal/common/domain"
	"cosmos.big2/internal/common/poker"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"strings"
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
	files := []string{
		"always-play-first-card.in",
		"fullhouse.in",
		"illegal-actions.in",
		"normal-no-error-play1.in",
		"normal-no-error-play2.in",
		"straight.in",
	}
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		deck := common.NewDeck()
		players := []domain.IPlayer{
			domain.NewHuman(),
			domain.NewHuman(),
			domain.NewHuman(),
			domain.NewHuman(),
		}
		big2 := domain.NewBig2(players, deck)
		i := -1
		scanner := bufio.NewScanner(file)

		var input []string
		for scanner.Scan() {
			i++
			text := scanner.Text()
			if i == 0 {
				deck.Reset()
				split := strings.Split(text, " ")
				for _, card := range split {
					deck.Push(poker.NewCard(poker.RankStringMap[card[2:len(card)-1]], poker.SuitStringMap[card[:1]]))
				}
				big2 = domain.NewBig2(players, deck)
				big2.DrawHand()
			} else if i > 0 && i <= 4 {
				big2.Players[i-1].NameSelf(text)
			} else {
				input = append(input, text)
			}
		}
		big2.PlayRound(input)

		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
