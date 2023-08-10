package main

import (
	"cosmos.MatchmakingSystem/internal/common"
)

func main() {
	individual := new(common.Individual)
	data := individual.GenerateRandomIndividual()

	matchMakingSystem := common.NewMatchMakingSystem(common.NewDistanceBaseStrategy())
	matchMakingSystem.Match(data)

	matchMakingSystem = common.NewMatchMakingSystem(common.NewHabitsBaseStrategy())
	matchMakingSystem.Match(data)

	matchMakingSystem = common.NewMatchMakingSystem(common.NewReverseStrategy(common.NewDistanceBaseStrategy()))
	matchMakingSystem.Match(data)
}
