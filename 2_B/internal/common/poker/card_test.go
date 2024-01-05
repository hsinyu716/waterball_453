package poker

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCardTestSuite(t *testing.T) {
	suite.Run(t, new(CardTestSuite))
}

type CardTestSuite struct {
	suite.Suite
}

func (s *CardTestSuite) SetupSuite() {
}

func (s *CardTestSuite) TestCard() {
	s.Equal(RankMap[0], "3")
	s.Equal(RankMap[12], "2")
	s.Equal(SuitMap[0], "C")
	s.Equal(SuitMap[3], "S")
}
