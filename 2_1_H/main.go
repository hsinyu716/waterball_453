package main

import (
	"cosmos.MatchmakingSystem/internal/common"
	"fmt"
)

func main() {
	individual := new(common.Individual)
	data := individual.GenerateRandomIndividual()

	fmt.Println("=====距離先決=========")
	matchMakingSystem := common.NewMatchMakingSystem(common.NewDistanceBaseStrategy())
	matchMakingSystem.Match(data)

	fmt.Println("=====興趣先決=========")
	matchMakingSystem = common.NewMatchMakingSystem(common.NewHabitsBaseStrategy())
	matchMakingSystem.Match(data)

	fmt.Println("=====反向距離=========")
	matchMakingSystem = common.NewMatchMakingSystem(common.NewReverseStrategy(common.NewDistanceBaseStrategy()))
	matchMakingSystem.Match(data)
}
