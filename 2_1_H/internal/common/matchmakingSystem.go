package common

type MatchmakingSystem struct {
	matchBy MatchmakingStrategy
}

type IMatchmakingSystem interface {
	matchSort(i Individual, data []Individual) Individual
}

func NewMatchMakingSystem(matchBy MatchmakingStrategy) *MatchmakingSystem {
	return &MatchmakingSystem{
		matchBy,
	}
}

func (m *MatchmakingSystem) Match(data []Individual) {
	for _, datum := range data {
		individual := m.matchSort(datum, data)
		datum.match(individual)
	}
}

func (m *MatchmakingSystem) matchSort(i Individual, data []Individual) Individual {
	//removeIndex := 0
	//for index, datum := range data {
	//	if i.ID == datum.ID {
	//		removeIndex = index
	//		continue
	//	}
	//	data[index] = m.matchBy.sortTemplate(i, datum)
	//}
	//tmpData := deepcopy.MustAnything(data).([]Individual)
	//tmpData = append(tmpData[:removeIndex], tmpData[removeIndex+1:]...)
	//sort.Slice(tmpData, func(i, j int) bool {
	//	return m.matchBy.compareTo(tmpData, i, j)
	//})
	tmpData := m.matchBy.matchSort(i, data)
	return tmpData[0]
}
